package main

import (
	"math"
)

type turtle struct {
	x     float64
	y     float64
	angle float64
}

func newTurtle() turtle {
	return turtle{
		x:     0.0,
		y:     0.0,
		angle: 0.0}
}

func rad(deg float64) float64 {
	return (deg / 180.0) * math.Pi
}

func deg(rad float64) float64 {
	return (rad / math.Pi) * 180.0
}

func moveTo(t turtle, x float64, y float64, vertices []float64) (turtle, []float64) {

	vertices = append(vertices, x)
	vertices = append(vertices, y)

	t.x = x
	t.y = y

	return t, vertices
}

func forward(t turtle, d float64, vertices []float64) (turtle, []float64) {
	return move(t, d, vertices)
}

func move(t turtle, d float64, vertices []float64) (turtle, []float64) {

	xn := d*math.Sin(rad(t.angle)) + t.x
	yn := d*math.Cos(rad(t.angle)) + t.y

	vertices = append(vertices, xn)
	vertices = append(vertices, yn)

	//fmt.Println(xn, yn)

	t.x = xn
	t.y = yn

	return t, vertices
}

func right(t turtle, delta float64) turtle {
	t.angle += delta
	return t
}

func left(t turtle, delta float64) turtle {
	t.angle -= delta
	return t
}

func turn(t turtle, delta float64) turtle {
	t.angle += delta
	return t
}
