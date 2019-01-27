#version 330

uniform float triangles[];
uniform int n_triangles;

in vec3 color;
in vec2 UV;

out vec4 frag_colour;

void main() {
	float r = triangles[3];
	float g = triangles[4];
	float b = triangles[5];
	frag_colour = vec4(r, g, b, 1.0);
}
