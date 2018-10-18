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

// Interfaces are named collections of method signatures.
//
type Object interface {
	//Type() string
	// Color() [3]int
	Draw()
}

type Vertex2 struct {
	X float32
	Y float32
}

type Lines struct {
	Vertices  []Vertex2
	LineWidth float32
	Color     [3]uint8
}

// NewSomething create new instance of Something
func NewLine(x1, y1, x2, y2 float32) Lines {
	obj := Lines{}
	obj.Vertices = append(obj.Vertices, Vertex2{x1, y1})
	obj.Vertices = append(obj.Vertices, Vertex2{x2, y2})
	obj.LineWidth = 1.0
	obj.Color = [3]uint8{255, 0, 0}
	return obj
}

func (l *Lines) Draw() {
	// Set line width
	gl.LineWidth(l.LineWidth)
	// Set Color
	gl.Color3ub(l.Color[0], l.Color[1], l.Color[2])
	// Set Vertices
	gl.Begin(gl.LINES)
	for _, v := range l.Vertices {
		gl.Vertex2f(v.X, v.Y)
	}
	gl.End()
}

// draw redraws the game board and the cells within.
func DrawScene(window *glfw.Window, objects []Object) {

	// Clear Screen
	//
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	//gl.UseProgram(prog) // TODO

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
	var objects []Object
	line := NewLine(-1, -1, 1, 1)
	objects = append(objects, &line)

	// Init GLFW
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	// Create Window
	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	// Init OpenGL
	//
	window.MakeContextCurrent()       // OK
	if err := gl.Init(); err != nil { // OK
		panic(err)
	}

	// Main loop
	//
	for !window.ShouldClose() {

		// Draw scene
		//
		DrawScene(window, objects)
	}
}
