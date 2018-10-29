package main

/*
View has to do with openGL.
Model has nothing to do with the view.

*/
import (
	"github.com/go-gl/gl/v2.1/gl"
)

// TODO:

// - Dimension const N = 2
// - Id, Name
// - Style (ColorARGB, Width)

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type Drawables interface {
	Draw()
}

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

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

type Points struct {
	Color    [3]uint8
	Width    float64
	Vertices []float64
}

type Lines struct {
	Color    [3]uint8
	Width    float64
	Vertices []float64
}

type LineStripe struct {
	Color    [3]uint8
	Width    float64
	Vertices []float64
}

type LineLoops struct {
	Color    [3]uint8
	Width    float64
	Vertices []float64
}

type QuadStrip struct {
	Color    [3]uint8
	Width    float64
	Vertices []float64 // xl,yl ---- xr,yr
}

// ----------------------------------------------------------------------------
// Creator functions
// ----------------------------------------------------------------------------

func createLines(vertices []float64) Lines {
	return Lines{
		Color:    [3]uint8{200, 200, 200},
		Width:    1.0,
		Vertices: vertices}
}

func createLineStripe(vertices []float64) LineStripe {
	return LineStripe{
		Color:    [3]uint8{64, 64, 255},
		Width:    3.0,
		Vertices: vertices}
}

func createQuadStrip(vertices []float64) QuadStrip {
	return QuadStrip{
		Color:    [3]uint8{32, 32, 32},
		Width:    1.0,
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
		Color:    [3]uint8{128, 128, 210},
		Width:    2.0,
		Vertices: vertices}
}

// ----------------------------------------------------------------------------
// Draw functions
// ----------------------------------------------------------------------------

/*
	Points
*/
func (ps *Points) Draw() {

	gl.PointSize(float32(ps.Width)) // DIFFERENT

	gl.Color3ub(ps.Color[0], ps.Color[1], ps.Color[2])

	gl.Begin(gl.POINTS) // DIFFERENT
	for i := 0; i < len(ps.Vertices); i = i + 2 {
		x := ps.Vertices[i+0]
		y := ps.Vertices[i+1]
		gl.Vertex2d(x, y)
	}
	gl.End()
}

/*
	Quadstrip
*/
func (ps *QuadStrip) Draw() {

	//gl.PointLineWidth(float32(ps.Width)) // DIFFERENT

	gl.Color3ub(ps.Color[0], ps.Color[1], ps.Color[2])

	gl.Begin(gl.QUAD_STRIP) // DIFFERENT

	for i := 0; i < len(ps.Vertices); i = i + (2 * N) {
		xl := ps.Vertices[i+0]
		yl := ps.Vertices[i+1]
		gl.Vertex2d(xl, yl)

		xr := ps.Vertices[i+2]
		yr := ps.Vertices[i+3]
		gl.Vertex2d(xr, yr)
	}

	gl.End()
}

/*
	Lines
*/
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

/*
	LineStripe
*/
func (ls *LineStripe) Draw() {

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

/*
	LineLoops
*/
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

// TODO
// ----

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
