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
		-1, -1, 1, -1, -1, 1,
		1, -1, 1, 1, -1, 1,
	}

	vao := makeVAO(vertices, uvs)

	triangles := []float32{0, 0, 0, 0.3, 0.3, 0, 0, 0.3, 0}
	tex := textureFromData(triangles)

	gl.UseProgram(prog)

	texUniform := gl.GetUniformLocation(prog, gl.Str("tex\x00"))
	gl.Uniform1i(texUniform, 0)

	nTrianglesUniform := gl.GetUniformLocation(prog, gl.Str("n_triangles\x00"))
	gl.Uniform1i(nTrianglesUniform, int32(len(triangles)/9))

	// View
	dirUniform := gl.GetUniformLocation(prog, gl.Str("dir\x00"))
	originUniform := gl.GetUniformLocation(prog, gl.Str("origin\x00"))
	rightUniform := gl.GetUniformLocation(prog, gl.Str("right\x00"))
	upUniform := gl.GetUniformLocation(prog, gl.Str("up\x00"))

	dir := []float32{0, 0, 1}
	origin := []float32{0, 0, -1}
	right := []float32{1, 0, 0}
	up := []float32{0, 1, 0}
	gl.Uniform3f(dirUniform, dir[0], dir[1], dir[2])
	gl.Uniform3f(originUniform, origin[0], origin[1], origin[2])
	gl.Uniform3f(rightUniform, right[0], right[1], right[2])
	gl.Uniform3f(upUniform, up[0], up[1], up[2])

	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	glfw.SwapInterval(1)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, tex)

		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, 6)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
