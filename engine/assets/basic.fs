#version 330

uniform vec3 dir;
uniform vec3 eye;
uniform vec3 up;
uniform vec3 right;

uniform int width;
uniform int height;

uniform int n_objs;
uniform sampler1D scene;
uniform sampler1D n_triangles;

in vec3 color;
in vec2 UV;

out vec4 frag_color;

float EPSILON = 1E-5;

struct intersect {
	bool ok;
	float dist;
	vec3 inter;
	vec3 normal;
};

struct face_intersect {
	int face_id;
	int obj;
	vec3 inter;
	vec3 normal;
};

struct material_t {
	vec4 ambiant;
	vec4 diffuse;
	vec4 specular;
	float shininess;
};

uniform material_t material[2];

struct light_t {
	vec3 direction;
	vec4 ambiant;
	float intensity;
};

uniform light_t light;

intersect is_ray_through(vec3 ray, vec3 origin, vec3 p1, vec3 p2, vec3 p3) {
	// Compute edges
	vec3 U = p2 - p1;
	vec3 V = p3 - p1;
	vec3 W = p3 - p2;

	// Compute the normal (check if not null)
	vec3 direction = cross(U, V);
	if (length(direction) <= 0.0) {
		discard;
	}

	vec3 n = normalize(direction);
	float d = dot(n, p1);

	// Get the intersection with the triangle plan
	float t = (d - dot(n, origin)) / dot(n, ray);
	vec3 intersection = origin + t * ray;

	// Get the result
	vec3 n1 = cross( U, intersection - p1);
	vec3 n2 = cross( W, intersection - p2);
	vec3 n3 = cross(-V, intersection - p3);

	bool b1 = dot(n1, n2) >= 0;
	bool b2 = dot(n2, n3) >= 0;

	intersect inter;
	inter.ok = b1 && b2;
	inter.dist = -t;
	inter.inter = intersection;
	inter.normal = n;
	return inter;
}

face_intersect face_through(vec3 ray, vec3 origin) {
	face_intersect f_inter;
	f_inter.face_id = -1;
	float min_dist = 1E5;

	int start = 0;
	for (int obj = 0; obj < n_objs; obj++) {
		int n_tri = int(texelFetch(n_triangles, obj, 0).r);
		for (int i = start; i < start + 3*n_tri; i+=3) {
			vec3 p1 = texelFetch(scene, i+0, 0).rgb;
			vec3 p2 = texelFetch(scene, i+1, 0).rgb;
			vec3 p3 = texelFetch(scene, i+2, 0).rgb;

			intersect inter = is_ray_through(ray, eye, p1, p2, p3);
			if (inter.ok && inter.dist > EPSILON && inter.dist < min_dist) {
				min_dist = inter.dist;
				f_inter.obj = obj;
				f_inter.face_id = i;
				f_inter.inter = inter.inter;
				f_inter.normal = inter.normal;
			}
		}

		start += 3 * n_tri;
	}

	return f_inter;
}

bool is_in_shadow(int face_id, vec3 origin) {
	int start = 0;
	for (int obj = 0; obj < n_objs; obj++) {
		int n_tri = int(texelFetch(n_triangles, obj, 0).r);
		for (int i = start; i < start + 3*n_tri; i+=3) {
			if (i == face_id) continue;

			vec3 p1 = texelFetch(scene, i+0, 0).rgb;
			vec3 p2 = texelFetch(scene, i+1, 0).rgb;
			vec3 p3 = texelFetch(scene, i+2, 0).rgb;
			intersect inter = is_ray_through(-light.direction, origin, p1, p2, p3);

			if (inter.ok && -inter.dist > EPSILON)
				return true;
		}

		start += 3 * n_tri;
	}

	return false;
}

void main() {
	// Antialiasing (supersampling)
	float offsets[15] = float[](
		0, 0, 0.4,
		0.3, 0.3, 0.15,
		0.3, -0.3, 0.15,
		-0.3, -0.3, 0.15,
		-0.3, 0.3, 0.15);

	vec4 color = vec4(0, 0, 0, 1);
	for (int o = 0; o < 5; o++) {
		float offset_x = offsets[3*o+0];
		float offset_y = offsets[3*o+1];
		float weight = offsets[3*o+2];

		// Ray vector
		vec3 ray = normalize(normalize(dir)*3 +
			right * (UV.x*width/height+offset_x/width) +
			up * (UV.y+offset_y/height));

		face_intersect f_inter = face_through(ray, eye);

		vec4 color_tmp = light.ambiant;
		if (f_inter.face_id >= 0) {
			int obj = f_inter.obj;

			// Face exposition
			float exposition = dot(normalize(f_inter.normal), normalize(-light.direction));

			// Ambiant color
			vec4 ambiant_color = light.ambiant * material[obj].ambiant;
			vec4 diffuse_color = light.intensity * material[obj].diffuse * exposition;
			vec4 specular_color = light.intensity * material[obj].specular * 0;

			color_tmp = ambiant_color;
			if (!is_in_shadow(f_inter.face_id, f_inter.inter)) {
				color_tmp += diffuse_color + specular_color;
			}
		}

		color += weight * color_tmp;
	}

	frag_color = color;
}
