package view

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func ClearBackground() {
	gl.Color3ub(255, 255, 255)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func SetCamera(camera Camera) {
	gl.LoadIdentity()
	gl.Scaled(camera.scale, camera.scale, 1.0)
	gl.Translated(camera.position[0], camera.position[1], 0.0)
}

func DrawScene(scene Scene) {
	for _, layer := range scene.layers {
		for _, obj := range layer.objects {
			obj.Draw()
		}
	}
}

func InitGlAndCreateWindow(width, height int, title string) *glfw.Window {

	// Init GLFW
	//
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	// Anti-Aliasing
	//
	glfw.WindowHint(glfw.Samples, 4)

	// Create Window
	//
	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	// Init OpenGL
	//
	if err := gl.Init(); err != nil {
		panic(err)
	}

	return window
}
