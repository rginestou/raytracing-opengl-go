package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	mgl "github.com/go-gl/mathgl/mgl32"
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

	triangles := []float32{
		-1.0, -1.0, -1.0, -1.0, -1.0, 1.0, -1.0, 1.0, 1.0,
		1.0, 1.0, -1.0, -1.0, -1.0, -1.0, -1.0, 1.0, -1.0,
		1.0, -1.0, 1.0, -1.0, -1.0, -1.0, 1.0, -1.0, -1.0,
		1.0, 1.0, -1.0, 1.0, -1.0, -1.0, -1.0, -1.0, -1.0,
		-1.0, -1.0, -1.0, -1.0, 1.0, 1.0, -1.0, 1.0, -1.0,
		1.0, -1.0, 1.0, -1.0, -1.0, 1.0, -1.0, -1.0, -1.0,
		-1.0, 1.0, 1.0, -1.0, -1.0, 1.0, 1.0, -1.0, 1.0,
		1.0, 1.0, 1.0, 1.0, -1.0, -1.0, 1.0, 1.0, -1.0,
		1.0, -1.0, -1.0, 1.0, 1.0, 1.0, 1.0, -1.0, 1.0,
		1.0, 1.0, 1.0, 1.0, 1.0, -1.0, -1.0, 1.0, -1.0,
		1.0, 1.0, 1.0, -1.0, 1.0, -1.0, -1.0, 1.0, 1.0,
		1.0, 1.0, 1.0, -1.0, 1.0, 1.0, 1.0, -1.0, 1.0,
	}
	tex := textureFromData(triangles)

	gl.UseProgram(prog)

	texUniform := gl.GetUniformLocation(prog, gl.Str("tex\x00"))
	gl.Uniform1i(texUniform, 0)

	nTrianglesUniform := gl.GetUniformLocation(prog, gl.Str("n_triangles\x00"))
	gl.Uniform1i(nTrianglesUniform, int32(len(triangles)/9))

	// View
	dirUniform := gl.GetUniformLocation(prog, gl.Str("dir\x00"))
	eyeUniform := gl.GetUniformLocation(prog, gl.Str("eye\x00"))
	rightUniform := gl.GetUniformLocation(prog, gl.Str("right\x00"))
	upUniform := gl.GetUniformLocation(prog, gl.Str("up\x00"))

	eye := mgl.Vec3{3, -0.3, -2}
	center := mgl.Vec3{1, 0.4, 0}
	view := mgl.LookAtV(eye, center, mgl.Vec3{0, 1, 0}).Inv()

	dir := view.Mul4x1(mgl.Vec4{0, 0, 1, 0})
	right := view.Mul4x1(mgl.Vec4{1, 0, 0, 0})
	up := view.Mul4x1(mgl.Vec4{0, 1, 0, 0})
	gl.Uniform3f(dirUniform, dir.X(), dir.Y(), dir.Z())
	gl.Uniform3f(eyeUniform, eye.X(), eye.Y(), eye.Z())
	gl.Uniform3f(rightUniform, right.X(), right.Y(), right.Z())
	gl.Uniform3f(upUniform, up.X(), up.Y(), up.Z())

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
