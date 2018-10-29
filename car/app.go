package main

import (
	"fmt"

	"github.com/go-gl/glfw/v3.2/glfw"
)

func createScene() Scene {

	// Examples data
	//
	// http://graphics.stanford.edu/data/3Dscanrep/

	// Layers
	layer0 := Layer{}
	layer1 := Layer{}
	layer2 := Layer{}

	// Create Model

	trackVerts, velocityData := createAnyTrack()

	scaleFactorCentrifugal := 5.0 / 1000.0
	scaleFactorVelocity := 2.0 / 100.0

	centrifugalVectors := calcCentrifugalVectors(trackVerts, velocityData, scaleFactorCentrifugal)
	veloctiyVectors := calcVelocityVectors(trackVerts, velocityData, scaleFactorVelocity)

	// Create ViewModels

	ovalTrack := createLineStripe(trackVerts)
	layer1.objects = append(layer1.objects, &ovalTrack)

	centrifugalVectorLines := createLines(centrifugalVectors)
	layer1.objects = append(layer1.objects, &centrifugalVectorLines)

	velocityVectorLines := createLines(veloctiyVectors)
	velocityVectorLines.Color = [3]uint8{255, 0, 0}
	layer1.objects = append(layer1.objects, &velocityVectorLines)

	trackPoints := createPoints(trackVerts)
	layer1.objects = append(layer1.objects, &trackPoints)

	// Track Area

	trackWidth := 8.0 // m
	trackAreaVerts := createTrackArea(trackVerts, trackWidth)
	trackAreaQuads := createQuadStrip(trackAreaVerts)
	layer0.objects = append(layer0.objects, &trackAreaQuads)

	trackAreaPoints := createPoints(trackAreaVerts)
	layer0.objects = append(layer0.objects, &trackAreaPoints)

	// Example ideal line

	idealLineVerts := createAnyTrackIdealLine()
	idealLine := createLineStripe(idealLineVerts)
	idealLine.Color = [3]uint8{64, 255, 64}
	layer2.objects = append(layer2.objects, &idealLine)
	/*
		centrifugalVectorsIL := calcCentrifugalVectors(idealLineVerts, velocityData, scaleFactorCentrifugal)
		centrifugalVectorLinesIL := createLines(centrifugalVectorsIL)
		centrifugalVectorLinesIL.Color = [3]uint8{128, 255, 128}
		layer2.objects = append(layer2.objects, &centrifugalVectorLinesIL)
	*/
	// Camera
	camera := Camera{}
	camera.scale = 0.02
	camera.position[0] = 0.0
	camera.position[1] = 0.0

	// Scene
	scene := Scene{}

	scene.layers = append(scene.layers, layer0)
	scene.layers = append(scene.layers, layer1)
	scene.layers = append(scene.layers, layer2)

	scene.camera = camera

	return scene
}

// Global TODO: ???
//
var (
	scene   Scene
	windowX int = 800
	windowY int = 800
)

// go run turtle.go calc.go calcTests.go model.go controller.go view.go viewModel.go app.go

func main() {

	// TODO: if tests OK start else ERROR do'nt start
	tests()

	// Init GL, GLFW
	//
	window := InitGlAndCreateWindow(windowX, windowY, "gl2d demo")
	if window == nil {
		glfw.Terminate()
		return
	}

	window.SetKeyCallback(keyCallback) // Controller

	// Create Scene
	//
	startTime := glfw.GetTime()

	scene = createScene()

	endTime := glfw.GetTime()
	deltaTime := endTime - startTime
	fmt.Printf("create scene took %f seconds \n", deltaTime)

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
