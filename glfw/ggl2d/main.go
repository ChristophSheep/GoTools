package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func ClearScene() {
	gl.Color3ub(255, 255, 255)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func SetCamera() {
	gl.LoadIdentity()
	gl.Scalef(0.02, 0.02, 1.0)
}

// draw redraws the game board and the cells within.
func DrawScene(drawables []Drawables) {
	for _, d := range drawables {
		d.Draw()
	}
}

func initWindow(width, height int, title string) *glfw.Window {

	// Init GLFW
	//
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

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

func createLines(vertices []float64) Lines {
	return Lines{
		Color:    [3]uint8{255, 128, 64},
		Width:    2.0,
		Vertices: vertices}
}

func createLineStripes(vertices []float64) LineStripes {
	return LineStripes{
		Color:    [3]uint8{128, 64, 255},
		Width:    1.0,
		Vertices: vertices}
}

func createLineLoops(vertices []float64) LineLoops {
	return LineLoops{
		Color:    [3]uint8{128, 64, 255},
		Width:    1.0,
		Vertices: vertices}
}

func createPoints(vertices []float64) Points {
	return Points{
		Color:    [3]uint8{255, 255, 255},
		Size:     2.0,
		Vertices: vertices}
}

func createObjects() []Drawables {
	var objects []Drawables

	/*
		linesVerts := []float32{
			+0.0, -0.5, -0.0, +0.5,
			-1.0, +1.0, +1.0, -1.0,
			-0.5, +0.0, +0.5, -0.0,
			-1.0, -1.0, +1.0, +1.0}

		stripeVerts := []float32{
			+0.0, +0.5,
			-0.5, +0.0,
			-0.0, -0.5,
			+0.5, -0.0}

		loopVerts := []float32{
			+0.0, +0.5,
			+0.0, +0.0,
			+0.5, +0.0}

			lines := createLines(linesVerts)
			objects = append(objects, &lines)

			lineStripe := createLineStripes(stripeVerts)
			objects = append(objects, &lineStripe)

			lineLoop := createLineLoops(loopVerts)
			objects = append(objects, &lineLoop)

			circle := createLineLoops(createCircle(0.5, 0.5, 0.5, 8))
			circle.Color = [3]uint8{63, 31, 255}
			objects = append(objects, &circle)

			points := createPoints(linesVerts)
			objects = append(objects, &points)

	*/

	ovalTrackVerts := createOvalTrack()

	normals := calcNormalVectors(ovalTrackVerts)
	normalLines := createLines(normals)
	objects = append(objects, &normalLines)

	ovalTrack := createLineLoops(ovalTrackVerts)
	objects = append(objects, &ovalTrack)

	trackPoints := createPoints(ovalTrackVerts)
	objects = append(objects, &trackPoints)

	return objects
}

// go run calc.go geom.go prim.go main.go
// go run calc.go geom.go prim.go turtle.go main.go

func main() {

	tests()

	// Examples data
	//
	// http://graphics.stanford.edu/data/3Dscanrep/

	// Create Objects
	//
	objects := createObjects()

	// Init GL, GLFW
	//
	window := initWindow(700, 700, "gl2d demo")

	// Main loop
	//
	for !window.ShouldClose() {
		ClearScene()
		SetCamera()
		DrawScene(objects)
		glfw.PollEvents()
		window.SwapBuffers()
	}

	// Clean up
	//
	glfw.Terminate()
}

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
