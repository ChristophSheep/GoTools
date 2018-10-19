package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Why Go’s structs are superior to class-based inheritance
// ========================================================
// see https://medium.com/@simplyianm/why-gos-structs-are-superior-to-class-based-inheritance-b661ba897c67
//

//
// see https://de.wikipedia.org/wiki/Grafisches_Primitiv
// see https://www.khronos.org/opengl/wiki/Primitive

//
// see https://tour.golang.org/moretypes/2

//
//      ** S T O P **   ->  WRONG WAY
//

// Computer is fast.
// Do not use object
// Used lists of point (streams)
// Hash table for fast access to Id or Name
//
// For example:
//

// https://www.khronos.org/opengl/wiki/Primitive#Point_primitives
// https://programming.guide/go/read-file-line-by-line.html
// https://stackoverflow.com/questions/14426366/what-is-an-idiomatic-way-of-representing-enums-in-go

func ClearScene() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// draw redraws the game board and the cells within.
func DrawScene(drawables []Drawables) {
	for _, d := range drawables {
		d.Draw()
	}
}

func initWindow(width, height int, title string) *glfw.Window {
	// Init GLFW
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	// Create Window
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

func createLines(vertices []float32) Lines {
	return Lines{
		Color:    [3]uint8{255, 128, 64},
		Width:    1.0,
		Vertices: vertices}
}

func createLineStripes(vertices []float32) LineStripes {
	return LineStripes{
		Color:    [3]uint8{128, 64, 255},
		Width:    1.0,
		Vertices: vertices}
}

func createPoints(vertices []float32) Points {
	return Points{
		Color:    [3]uint8{255, 255, 0},
		Size:     8.0,
		Vertices: vertices}
}

func main() {

	// Create Objects
	//
	var objects []Drawables

	vertices := []float32{
		+0.0, -0.5, -0.0, +0.5,
		-0.5, +0.0, +0.5, -0.0,
		-1.0, +1.0, +1.0, -1.0,
		-1.0, -1.0, +1.0, +1.0}

	lines := createLines(vertices)
	objects = append(objects, &lines)

	lineStripes := createLineStripes(vertices)
	objects = append(objects, &lineStripes)

	points := createPoints(vertices)
	objects = append(objects, &points)

	// Init GL, GLFW
	//
	window := initWindow(400, 400, "gl2d demo")

	// Main loop
	//
	for !window.ShouldClose() {
		ClearScene()
		DrawScene(objects)
		glfw.PollEvents()
		window.SwapBuffers()
	}

	// Clean up
	//
	glfw.Terminate()
}
