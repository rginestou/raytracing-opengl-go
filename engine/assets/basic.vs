#version 330

in vec2 vert;
in vec2 uv;

out vec2 UV;

void main() {
	gl_Position = vec4(vert, 0, 1);
	UV = uv;
}
