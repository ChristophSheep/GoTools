package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

func createScene() Scene {

	// Examples data
	//
	// http://graphics.stanford.edu/data/3Dscanrep/

	// Create Model
	ovalTrackVerts := createOvalTrack2()
	normalVectors := calcNormalVectors(ovalTrackVerts)

	// Create ViewModels
	var objects []Drawables

	ovalTrack := createLineLoops(ovalTrackVerts)
	objects = append(objects, &ovalTrack)

	normalVectorLines := createLines(normalVectors)
	objects = append(objects, &normalVectorLines)

	trackPoints := createPoints(ovalTrackVerts)
	objects = append(objects, &trackPoints)

	// Layer
	layer := Layer{}
	layer.objects = objects

	// Camera
	camera := Camera{}
	camera.scale = 0.02
	camera.position[0] = 0.0
	camera.position[1] = 0.0

	// Scene
	scene := Scene{}
	scene.layers = append(scene.layers, layer)
	scene.camera = camera

	return scene
}

// go run calc.go geom.go prim.go main.go
// go run calc.go geom.go prim.go turtle.go main.go

func main() {

	// tests()

	// model := createModel()
	// viewModel = createViewModel(model)

	// Create Objects
	//
	scene := createScene()

	//	glfw.WindowHint(glfw.DEPTH_BITS, 16);
	//  glfw.WindowHint(glfw.TRANSPARENT_FRAMEBUFFER, GLFW_TRUE);

	// Init GL, GLFW
	//
	window := initWindow(700, 700, "gl2d demo")
	if window == nil {
		glfw.Terminate()
		return
	}

	window.SetKeyCallback(keyCallback) // Controller

	// Main loop
	//
	for !window.ShouldClose() {

		ClearBackground()
		SetCamera(scene.camera)
		DrawScene(scene)

		glfw.PollEvents()
		window.SwapBuffers()

		// TODO: wait 60fps
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
