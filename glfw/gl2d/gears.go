package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.0/glfw"
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

type Vertex2 struct {
	X float32
	Y float32
}

type Object interface {
	Draw()
}

type Lines struct {
	vertices []Vertex2
	Draw()
}

func (l *Lines) Draw() {
	gl.Begin(gl.LINES)
	for vertex := range l.vertices {
		gl.Vertex2f(vertex.X, vertex.Y)	
	}
	gl.End(
}

// draw redraws the game board and the cells within.
func Draw(prog uint32, window *glfw.Window, objects []*object) {

	// Clear Screen
	//
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(prog)


	// Draw Objects
	//
	for obj := range objects {
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
	var objects []object
	
	lines := &Lines{}

	objects = append(objects, obj)


	// Main loop
	//
	for !window.ShouldClose() {

		// Draw scene
		//
		draw(prog, window, objects)
	}
}
