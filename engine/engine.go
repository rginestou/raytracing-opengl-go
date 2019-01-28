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
		-1.0, 0.0, -1.0, -1.0, 0.0, 1.0, -1.0, 2.0, 1.0,
		1.0, 2.0, -1.0, -1.0, 0.0, -1.0, -1.0, 2.0, -1.0,
		1.0, 0.0, 1.0, -1.0, 0.0, -1.0, 1.0, 0.0, -1.0,
		1.0, 2.0, -1.0, 1.0, 0.0, -1.0, -1.0, 0.0, -1.0,
		-1.0, 0.0, -1.0, -1.0, 2.0, 1.0, -1.0, 2.0, -1.0,
		1.0, 0.0, 1.0, -1.0, 0.0, 1.0, -1.0, 0.0, -1.0,
		-1.0, 2.0, 1.0, -1.0, 0.0, 1.0, 1.0, 0.0, 1.0,
		1.0, 2.0, 1.0, 1.0, 0.0, -1.0, 1.0, 2.0, -1.0,
		1.0, 0.0, -1.0, 1.0, 2.0, 1.0, 1.0, 0.0, 1.0,
		1.0, 2.0, 1.0, 1.0, 2.0, -1.0, -1.0, 2.0, -1.0,
		1.0, 2.0, 1.0, -1.0, 2.0, -1.0, -1.0, 2.0, 1.0,
		1.0, 2.0, 1.0, -1.0, 2.0, 1.0, 1.0, 0.0, 1.0,

		10.0, 0.0, 10.0, 10.0, 0.0, -10.0, -10.0, 0.0, -10.0,
		10.0, 0.0, 10.0, -10.0, 0.0, -10.0, -10.0, 0.0, 10.0,
	}
	counts := []float32{
		12, 2,
	}

	scene := textureFromData(gl.Ptr(triangles), len(triangles), gl.TEXTURE0, gl.RGB32F, gl.RGB, gl.FLOAT)
	nTriangles := textureFromData(gl.Ptr(counts), len(counts), gl.TEXTURE1, gl.R32F, gl.RED, gl.FLOAT)
	width, height := window.GetSize()

	// Uniforms
	gl.UseProgram(prog)
	nObjsUniform := gl.GetUniformLocation(prog, gl.Str("n_objs\x00"))
	sceneUniform := gl.GetUniformLocation(prog, gl.Str("scene\x00"))
	nTrianglesUniform := gl.GetUniformLocation(prog, gl.Str("n_triangles\x00"))
	dirUniform := gl.GetUniformLocation(prog, gl.Str("dir\x00"))
	eyeUniform := gl.GetUniformLocation(prog, gl.Str("eye\x00"))
	rightUniform := gl.GetUniformLocation(prog, gl.Str("right\x00"))
	upUniform := gl.GetUniformLocation(prog, gl.Str("up\x00"))
	widthUniform := gl.GetUniformLocation(prog, gl.Str("width\x00"))
	heightUniform := gl.GetUniformLocation(prog, gl.Str("height\x00"))

	material0AmbiantUniform := gl.GetUniformLocation(prog, gl.Str("material[0].ambiant\x00"))
	material0DiffuseUniform := gl.GetUniformLocation(prog, gl.Str("material[0].diffuse\x00"))
	material0SpecularUniform := gl.GetUniformLocation(prog, gl.Str("material[0].specular\x00"))
	material0ShininessUniform := gl.GetUniformLocation(prog, gl.Str("material[0].shininess\x00"))

	material1AmbiantUniform := gl.GetUniformLocation(prog, gl.Str("material[1].ambiant\x00"))
	material1DiffuseUniform := gl.GetUniformLocation(prog, gl.Str("material[1].diffuse\x00"))
	material1SpecularUniform := gl.GetUniformLocation(prog, gl.Str("material[1].specular\x00"))
	material1ShininessUniform := gl.GetUniformLocation(prog, gl.Str("material[1].shininess\x00"))

	lightDirUniform := gl.GetUniformLocation(prog, gl.Str("light.direction\x00"))
	lightAmbiantUniform := gl.GetUniformLocation(prog, gl.Str("light.ambiant\x00"))
	lightIntensityUniform := gl.GetUniformLocation(prog, gl.Str("light.intensity\x00"))

	gl.Uniform1i(nObjsUniform, int32(len(counts)))
	gl.Uniform1i(sceneUniform, 0)
	gl.Uniform1i(nTrianglesUniform, 1)
	gl.Uniform1i(widthUniform, int32(width))
	gl.Uniform1i(heightUniform, int32(height))

	gl.Uniform4f(material0AmbiantUniform, 0.25, 0.25, 0.25, 1)
	gl.Uniform4f(material0DiffuseUniform, 0.4, 0.4, 0.4, 1)
	gl.Uniform4f(material0SpecularUniform, 0.774597, 0.774597, 0.774597, 1)
	gl.Uniform1f(material0ShininessUniform, 76.8)

	gl.Uniform4f(material1AmbiantUniform, 0.3, 0.3, 0.3, 1)
	gl.Uniform4f(material1DiffuseUniform, 1, 1, 1, 0)
	gl.Uniform4f(material1SpecularUniform, 1, 1, 1, 0)
	gl.Uniform1f(material1ShininessUniform, 0)

	gl.Uniform3f(lightDirUniform, -1.54, -4, 2.3)
	gl.Uniform4f(lightAmbiantUniform, 1, 1, 1, 1)
	gl.Uniform1f(lightIntensityUniform, 0.95)

	// Engine state
	isRunning := true

	window.SetKeyCallback(func(w *glfw.Window, k glfw.Key, st int, a glfw.Action, mk glfw.ModifierKey) {
		if a == glfw.Press && k == glfw.KeyEscape {
			isRunning = false
		}
	})

	buttonClicked := false
	window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
		if button == glfw.MouseButton1 {
			if action == glfw.Press {
				buttonClicked = true
			} else if action == glfw.Release {
				buttonClicked = false
			}
		}
	})

	mx, _ := window.GetCursorPos()
	theta := 0.0
	window.SetCursorPosCallback(func(w *glfw.Window, xpos float64, ypos float64) {
		if buttonClicked {
			dx := xpos - mx
			theta += dx * 0.01
		}
		mx = xpos
	})

	window.SetSizeCallback(func(w *glfw.Window, width, height int) {
		gl.Uniform1i(widthUniform, int32(width))
		gl.Uniform1i(heightUniform, int32(height))
	})

	gl.Enable(gl.BLEND)
	glfw.SwapInterval(1)
	for !window.ShouldClose() && isRunning {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		width, height := window.GetSize()
		gl.Viewport(0, 0, int32(width), int32(height))

		// View
		eye := mgl.Vec3{7, 2.5, 7}
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
		gl.BindTexture(gl.TEXTURE_1D, scene)
		gl.ActiveTexture(gl.TEXTURE1)
		gl.BindTexture(gl.TEXTURE_1D, nTriangles)

		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, 6)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
