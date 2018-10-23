package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func ClearBackground() {
	gl.Color3ub(255, 255, 255)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

type Layer struct {
	objects []Drawables
}

type Camera struct {
	scale    float64
	position [2]float64
}

type Scene struct {
	layers []Layer
	camera Camera
}

func SetCamera(camera Camera) {
	gl.LoadIdentity()
	gl.Scaled(camera.scale, camera.scale, 1.0)
	gl.Translated(camera.position[0], camera.position[1], 0.0)
}

// draw redraws the game board and the cells within.
func DrawScene(scene Scene) {
	for _, layer := range scene.layers {
		for _, obj := range layer.objects {
			obj.Draw()
		}
	}
}

func initWindow(width, height int, title string) *glfw.Window {

	// Init GLFW
	//
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	// Anti Aliasing
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

func createLines(vertices []float64) Lines {
	return Lines{
		Color:    [3]uint8{200, 200, 200},
		Width:    2.0,
		Vertices: vertices}
}

func createLineStripes(vertices []float64) LineStripes {
	return LineStripes{
		Color:    [3]uint8{128, 64, 255},
		Width:    3.0,
		Vertices: vertices}
}

func createLineLoops(vertices []float64) LineLoops {
	return LineLoops{
		Color:    [3]uint8{255, 255, 128},
		Width:    3.0,
		Vertices: vertices}
}

func createPoints(vertices []float64) Points {
	return Points{
		Color:    [3]uint8{64, 64, 255},
		Size:     7.0,
		Vertices: vertices}
}

func createScene() Scene {
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

	ovalTrackVerts := createOvalTrack2()

	normals := calcNormalVectors(ovalTrackVerts)
	normalLines := createLines(normals)
	objects = append(objects, &normalLines)

	ovalTrack := createLineLoops(ovalTrackVerts)
	objects = append(objects, &ovalTrack)

	trackPoints := createPoints(ovalTrackVerts)
	objects = append(objects, &trackPoints)

	layer := Layer{}
	layer.objects = objects

	camera := Camera{}
	camera.scale = 0.02
	camera.position[0] = 0.0
	camera.position[1] = 0.0

	scene := Scene{}
	scene.layers = append(scene.layers, layer)
	scene.camera = camera

	return scene
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

	keyCallback := func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

		// ZOOM - in, out
		//
		if mods == glfw.ModShift {
			if key == glfw.KeyUp {
				scene.camera.scale *= 0.5
			}
			if key == glfw.KeyDown {
				scene.camera.scale /= 0.5
			}
		}

		// MOVE CAMERA - left, right, up, down
		//
		if key == glfw.KeyUp {
			scene.camera.position[1] += 1.0
		}
		if key == glfw.KeyDown {
			scene.camera.position[1] -= 1.0
		}
		if key == glfw.KeyLeft {
			scene.camera.position[0] -= 1.0
		}
		if key == glfw.KeyRight {
			scene.camera.position[0] += 1.0
		}

	}

	window.SetKeyCallback(keyCallback)

	// Main loop
	//
	for !window.ShouldClose() {
		ClearBackground()
		SetCamera(scene.camera)
		DrawScene(scene)
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
