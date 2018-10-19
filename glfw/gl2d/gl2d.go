package main

import (
	"fmt"

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

//  id 		 	uint64		= 0
//  type		uint32		= 1 // lines
//  name     	string		= "lines_1"
//  lineWidth 	float32		= 1.0
//  color 	 	[3]uint8	= 255 0 0 0  // rgba or argb
//  vertices 	[]float32	= x y x y x y .... x y
//

//  id 		 	uint64		= 1
//  type		uint32		= 2 // polyline -> line_strip -> line_loop (closed strip)
//  name     	string		= "polyline2D_1"
//  lineWidth 	float32		= 1.0
//  color 	 	[3]uint8	= 255 0 0 0  // rgba or argb
//  vertices 	[]float32	= x y x y x y .... x y
//

// See how AutoCAD Works
// A long list of objects

// Interfaces are named collections of method signatures.
//
type Object interface {
	Draw()
}

// Implement interface
//
func (pl *PolyLine) Draw() {
	// TODO
}

/*
// NewSomething create new instance of Something
func NewLine(x1, y1, x2, y2 float32) Lines {
	obj := Lines{}
	obj.Vertices = append(obj.Vertices, Vertex2{x1, y1})
	obj.Vertices = append(obj.Vertices, Vertex2{x2, y2})
	obj.LineWidth = 2.0
	obj.Color = [3]uint8{255, 128, 0}
	return obj
}

func (l *Lines) Draw() {
	// Set line width
	//
	gl.LineWidth(l.LineWidth)
	// Set Color
	//
	gl.Color3ub(l.Color[0], l.Color[1], l.Color[2])
	// Set Vertices
	//
	gl.Begin(gl.LINES)
	for _, v := range l.Vertices {
		gl.Vertex2f(v.X, v.Y)
	}
	gl.End()
}
*/

// draw redraws the game board and the cells within.
func DrawScene(window *glfw.Window, objects []Object, angle *float32) {
	// Clear Screen
	//
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	// Draw Objects
	//
	for _, obj := range objects {
		obj.Draw()
	}

	// Poll Events
	glfw.PollEvents()
	// Swap Buffer
	window.SwapBuffers()
}

func main() {

	// Create Objects
	//

	const N = 4

	points := [4 * N]float32{
		+0.0, -0.5, -0.0, +0.5,
		-0.5, +0.0, +0.5, -0.0,
		-1.0, +1.0, +1.0, -1.0,
		-1.0, -1.0, +1.0, +1.0}

	var objects []Object

	for i := 0; i < len(points); i = i + N {
		fmt.Println(i)
		line := NewLine(points[i+0], points[i+1], points[i+2], points[i+3])
		objects = append(objects, &line)
	}

	// Init GLFW
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	// Create Window
	window, err := glfw.CreateWindow(400, 400, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	// Init OpenGL
	//
	window.MakeContextCurrent()       // OK
	if err := gl.Init(); err != nil { // OK
		panic(err)
	}

	angle := float32(0.0)

	// Main loop
	//
	for !window.ShouldClose() {
		DrawScene(window, objects, &angle)
	}
}
