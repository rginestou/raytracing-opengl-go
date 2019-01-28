package engine

import (
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func textureFromData(data unsafe.Pointer, length int, id uint32, iformat int32, format uint32, t uint32) uint32 {
	var textureID uint32
	gl.GenTextures(1, &textureID)

	gl.ActiveTexture(id)
	gl.BindTexture(gl.TEXTURE_1D, textureID)

	gl.TexParameteri(gl.TEXTURE_1D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_1D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)

	gl.TexImage1D(gl.TEXTURE_1D, 0, iformat, int32(length), 0, format, t, data)

	return textureID
}
