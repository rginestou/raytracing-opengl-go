#version 330

uniform vec3 dir;
uniform vec3 origin;
uniform vec3 up;
uniform vec3 right;

uniform int n_triangles;
uniform sampler1D tex;

in vec3 color;
in vec2 UV;

out vec4 frag_color;

bool isRayThrough(vec3 ray, vec3 p1, vec3 p2, vec3 p3) {
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
	vec3 inter = origin + t * ray;

	// Get the result
	vec3 n1 = cross( U, inter - p1);
	vec3 n2 = cross( W, inter - p2);
	vec3 n3 = cross(-V, inter - p3);

	bool b1 = dot(n1, n2) >= 0;
	bool b2 = dot(n2, n3) >= 0;

	return b1 && b2;
}

void main() {
	vec3 ray = normalize(dir + right * UV.x + up * UV.y);

	bool visible = false;
	for (int i = 0; i < 3*n_triangles; i+=3) {
		vec3 p1 = texelFetch(tex, i+0, 0).rgb;
		vec3 p2 = texelFetch(tex, i+1, 0).rgb;
		vec3 p3 = texelFetch(tex, i+2, 0).rgb;

		visible = visible || isRayThrough(ray, p1, p2, p3);
	}

	float r = 0.0;
	if (visible)
		r = 1.0;

	frag_color = vec4(r, 0, 0, 1.0);
}
