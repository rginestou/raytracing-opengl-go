package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Run ...
func Run(window *glfw.Window) {
	prog := createShaderProgram("engine/assets/basic.vs", "engine/assets/basic.fs")

	vertices := []float32{
		-1, -1, 1, -1, -1, 1,
		1, -1, 1, 1, -1, 1,
	}
	uvs := []float32{
		0, 0, 1, 0, 0, 1,
		1, 0, 1, 1, 0, 1,
	}

	vao := makeVAO(vertices, uvs)

	gl.UseProgram(prog)
	triangles := []float32{0, 0, 0, 1, 1, 0, 0, 1, 0}
	trianglesUniform := gl.GetUniformLocation(prog, gl.Str("triangles\x00"))
	gl.Uniform1fv(trianglesUniform, int32(len(triangles)), &triangles[0])
	nTriangles := len(triangles) / 9
	nTrianglesUniform := gl.GetUniformLocation(prog, gl.Str("n_triangles\x00"))
	gl.Uniform1i(nTrianglesUniform, int32(nTriangles))

	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	glfw.SwapInterval(1)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, 6)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
