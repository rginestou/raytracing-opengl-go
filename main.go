package main

import (
	"raytracing/engine"
	"raytracing/window"
)

func main() {
	window.Create(600, 600)

	engine.Run(window.GetPtr())

	window.Stop()
}
