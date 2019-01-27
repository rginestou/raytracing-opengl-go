package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	mgl "github.com/go-gl/mathgl/mgl32"
)

// Run ...
func Run(window *glfw.Window) {
	// Shader program
	prog := createShaderProgram("engine/assets/basic.vs", "engine/assets/basic.fs")

	// Canvas
	vertices := []float32{
		-1, -1, 1, -1, -1, 1,
		1, -1, 1, 1, -1, 1,
	}
	uvs := []float32{
		-1, 1, 1, 1, -1, -1,
		1, 1, 1, -1, -1, -1,
	}
	vao := makeVAO(vertices, uvs)

	// Scene geometry
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
	width, height := window.GetSize()

	// Uniforms
	gl.UseProgram(prog)
	texUniform := gl.GetUniformLocation(prog, gl.Str("tex\x00"))
	nTrianglesUniform := gl.GetUniformLocation(prog, gl.Str("n_triangles\x00"))
	dirUniform := gl.GetUniformLocation(prog, gl.Str("dir\x00"))
	eyeUniform := gl.GetUniformLocation(prog, gl.Str("eye\x00"))
	rightUniform := gl.GetUniformLocation(prog, gl.Str("right\x00"))
	upUniform := gl.GetUniformLocation(prog, gl.Str("up\x00"))
	widthUniform := gl.GetUniformLocation(prog, gl.Str("width\x00"))
	heightUniform := gl.GetUniformLocation(prog, gl.Str("height\x00"))

	materialAmbiantUniform := gl.GetUniformLocation(prog, gl.Str("material.ambiant\x00"))
	materialDiffuseUniform := gl.GetUniformLocation(prog, gl.Str("material.diffuse\x00"))
	materialSpecularUniform := gl.GetUniformLocation(prog, gl.Str("material.specular\x00"))
	materialShininessUniform := gl.GetUniformLocation(prog, gl.Str("material.shininess\x00"))

	lightDirUniform := gl.GetUniformLocation(prog, gl.Str("light.direction\x00"))
	lightAmbiantUniform := gl.GetUniformLocation(prog, gl.Str("light.ambiant\x00"))
	lightIntensityUniform := gl.GetUniformLocation(prog, gl.Str("light.intensity\x00"))

	gl.Uniform1i(texUniform, 0)
	gl.Uniform1i(nTrianglesUniform, int32(len(triangles)/9))
	gl.Uniform1i(widthUniform, int32(width))
	gl.Uniform1i(heightUniform, int32(height))

	gl.Uniform4f(materialAmbiantUniform, 0.25, 0.25, 0.25, 1)
	gl.Uniform4f(materialDiffuseUniform, 0.4, 0.4, 0.4, 1)
	gl.Uniform4f(materialSpecularUniform, 0.774597, 0.774597, 0.774597, 1)
	gl.Uniform1f(materialShininessUniform, 76.8)

	gl.Uniform3f(lightDirUniform, -1.54, -0.2, 2.3)
	gl.Uniform4f(lightAmbiantUniform, 1, 1, 1, 1)
	gl.Uniform1f(lightIntensityUniform, 0.95)

	// Engine state
	isRunning := true

	window.SetKeyCallback(func(w *glfw.Window, k glfw.Key, st int, a glfw.Action, mk glfw.ModifierKey) {
		if a == glfw.Press && k == glfw.KeyEscape {
			isRunning = false
		}
	})

	mx, _ := window.GetCursorPos()
	theta := 0.0
	window.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		dx := xpos - mx
		theta += dx * 0.01
		mx = xpos
	})

	glfw.SwapInterval(1)
	for !window.ShouldClose() && isRunning {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// View
		eye := mgl.Vec3{7, 1.5, 7}
		rot := mgl.Rotate3DY(float32(theta))
		eye = rot.Mul3x1(eye)

		center := mgl.Vec3{0, 0, 0}
		view := mgl.LookAtV(eye, center, mgl.Vec3{0, 1, 0}).Inv()

		dir := view.Mul4x1(mgl.Vec4{0, 0, 1, 0})
		right := view.Mul4x1(mgl.Vec4{1, 0, 0, 0})
		up := view.Mul4x1(mgl.Vec4{0, 1, 0, 0})

		gl.Uniform3f(dirUniform, dir.X(), dir.Y(), dir.Z())
		gl.Uniform3f(eyeUniform, eye.X(), eye.Y(), eye.Z())
		gl.Uniform3f(rightUniform, right.X(), right.Y(), right.Z())
		gl.Uniform3f(upUniform, up.X(), up.Y(), up.Z())

		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, tex)

		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, 6)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
