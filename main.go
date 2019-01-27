package main

import (
	"raytracing/engine"
	"raytracing/window"
)

func main() {
	window.Create(512, 512)

	engine.Run(window.GetPtr())

	window.Stop()
}
