#version 330

uniform vec3 dir;
uniform vec3 eye;
uniform vec3 up;
uniform vec3 right;

uniform int n_triangles;
uniform sampler1D tex;

in vec3 color;
in vec2 UV;

out vec4 frag_color;

struct intersect {
	bool ok;
	float dist;
	vec3 inter;
};

struct face_intersect {
	int face_id;
	vec3 inter;
};

vec3 sun_dir = vec3(2, 1, 0.5);

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
	inter.dist = t;
	inter.inter = intersection;
	return inter;
}

face_intersect face_through(vec3 ray, vec3 origin) {
	face_intersect f_inter;
	f_inter.face_id = -1;
	float min_dist = 10000;

	for (int i = 0; i < 3*n_triangles; i+=3) {
		vec3 p1 = texelFetch(tex, i+0, 0).rgb;
		vec3 p2 = texelFetch(tex, i+1, 0).rgb;
		vec3 p3 = texelFetch(tex, i+2, 0).rgb;

		intersect inter = is_ray_through(ray, eye, p1, p2, p3);
		if (inter.ok && -inter.dist > 0.01 && -inter.dist < min_dist) {
			min_dist = inter.dist;
			f_inter.face_id = i;
			f_inter.inter = inter.inter;
		}
	}

	return f_inter;
}

bool is_in_shadow(int face_id, vec3 origin) {
	for (int i = 0; i < 3*n_triangles; i+=3) {
		if (i == face_id) continue;

		vec3 p1 = texelFetch(tex, i+0, 0).rgb;
		vec3 p2 = texelFetch(tex, i+1, 0).rgb;
		vec3 p3 = texelFetch(tex, i+2, 0).rgb;
		intersect inter = is_ray_through(-sun_dir, origin, p1, p2, p3);

		if (inter.ok && inter.dist > 0.01)
			return true;
	}

	return false;
}

void main() {
	vec3 ray = normalize(normalize(dir)*3 + right * UV.x + up * UV.y);

	vec3 color = vec3(1,1,1);
	face_intersect f_inter = face_through(ray, eye);
	if (f_inter.face_id == -1) {
	} else if (is_in_shadow(f_inter.face_id, f_inter.inter)) {
		color = vec3(0.1, 0.1, 0.1);
	} else {
		color = vec3(0.5, 0.5, 0.5);
	}

	frag_color = vec4(color, 1.0);
}
