package window

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

var window *glfw.Window

// GetPtr ...
func GetPtr() *glfw.Window {
	return window
}

// Create ...
func Create(width, height int) {
	if err := glfw.Init(); err != nil {
		log.Panicln(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// Create Window
	var err error
	window, err = glfw.CreateWindow(width, height, "RayTracing", nil, nil)
	if err != nil {
		log.Panicln(err)
	}

	window.MakeContextCurrent()

	// Init GL context
	if err := gl.Init(); err != nil {
		log.Panicln(err)
	}

	// Display version
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)
}

// Stop the window context
func Stop() {
	glfw.Terminate()
}
