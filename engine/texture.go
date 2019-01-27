package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

func textureFromData(data []float32) uint32 {
	// Create one OpenGL texture
	var textureID uint32
	gl.GenTextures(1, &textureID)

	// "Bind" the newly created texture : all future texture functions will modify this texture
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_1D, textureID)

	gl.TexParameteri(gl.TEXTURE_1D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_1D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)

	gl.TexImage1D(gl.TEXTURE_1D, 0, gl.RGB32F, int32(len(data)), 0, gl.RGB, gl.FLOAT, gl.Ptr(data))

	return textureID
}
