package main

import (
	"github.com/go-gl/gl/v2.1/gl"
)

//  TODOS
//  =====
//
// 		- Drawing order
// 		- Id
// 		- Name

// ----------------------------------------------------------------------------
// Drawables
//

type Drawables interface { // Interfaces are named collections of method signatures.
	Draw()
}

// ----------------------------------------------------------------------------
// Points
//

type Points struct {
	Color    [3]uint8
	Size     float64
	Vertices []float64 // x y x y
}

func (ps *Points) Draw() {
	gl.PointSize(float32(ps.Size)) // DIFFERENT
	gl.Color3ub(ps.Color[0], ps.Color[1], ps.Color[2])
	gl.Begin(gl.POINTS) // DIFFERENT
	for i := 0; i < len(ps.Vertices); i = i + 2 {
		x := ps.Vertices[i+0]
		y := ps.Vertices[i+1]
		gl.Vertex2d(x, y)
	}
	gl.End()
}

// ----------------------------------------------------------------------------
// Lines
//

type Lines struct {
	Color    [3]uint8
	Width    float64
	Vertices []float64 // x y x y   x y x y   x y x y
}

func (ls *Lines) Draw() {
	gl.LineWidth(float32(ls.Width)) // DIFFERENT
	gl.Color3ub(ls.Color[0], ls.Color[1], ls.Color[2])
	gl.Begin(gl.LINES) // DIFFERENT
	for i := 0; i < len(ls.Vertices); i = i + 2 {
		x := ls.Vertices[i+0]
		y := ls.Vertices[i+1]
		gl.Vertex2d(x, y)
	}
	gl.End()
}

// ----------------------------------------------------------------------------
// LineStripes
//

type LineStripes struct {
	Color    [3]uint8
	Width    float64
	Vertices []float64 // x y -- x y -- x y -- x y -- x y -- x y
}

func (ls *LineStripes) Draw() {

	gl.LineWidth(float32(ls.Width)) // DIFFERENT

	gl.Color3ub(ls.Color[0], ls.Color[1], ls.Color[2])

	gl.Begin(gl.LINE_STRIP) // DIFFERENT
	for i := 0; i < len(ls.Vertices); i = i + 2 {
		x := ls.Vertices[i+0]
		y := ls.Vertices[i+1]
		gl.Vertex2d(x, y)
	}
	gl.End()
}

// ----------------------------------------------------------------------------
// LineLoops
//

type LineLoops struct {
	Color    [3]uint8
	Width    float64
	Vertices []float64 // x y -- x y -- x y -- x y -- x y -- x y --> START

}

func (ls *LineLoops) Draw() {

	gl.LineWidth(float32(ls.Width))

	gl.Color3ub(ls.Color[0], ls.Color[1], ls.Color[2])

	gl.Begin(gl.LINE_LOOP)
	for i := 0; i < len(ls.Vertices); i = i + 2 {
		x := ls.Vertices[i+0]
		y := ls.Vertices[i+1]
		gl.Vertex2d(x, y)
	}
	gl.End()
}

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
