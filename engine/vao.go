package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

const sizeOfFloat32 = 4

// Returns a vertex array from the vertices provided
func makeVAO(vertices []float32, uvs []float32) uint32 {
	// Create VAO buffer
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	// Create Vertices buffer
	if vertices != nil {
		var vbo uint32
		gl.GenBuffers(1, &vbo)
		gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
		gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*sizeOfFloat32, gl.Ptr(vertices), gl.STATIC_DRAW)

		gl.EnableVertexAttribArray(0)
		gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)
	}

	// Create UVs buffer
	if uvs != nil {
		var ubo uint32
		gl.GenBuffers(1, &ubo)
		gl.BindBuffer(gl.ARRAY_BUFFER, ubo)
		gl.BufferData(gl.ARRAY_BUFFER, len(uvs)*sizeOfFloat32, gl.Ptr(uvs), gl.STATIC_DRAW)

		gl.EnableVertexAttribArray(1)
		gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, nil)
	}

	// Disable VAO
	gl.BindVertexArray(0)

	return vao
}
