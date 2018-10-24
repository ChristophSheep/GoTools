package model

import (
	"math"
)

type Turtle struct {
	x     float64
	y     float64
	angle float64
}

func newTurtle() Turtle {
	return Turtle{
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

func moveTo(t Turtle, x float64, y float64, vertices []float64) (Turtle, []float64) {

	vertices = append(vertices, x)
	vertices = append(vertices, y)

	t.x = x
	t.y = y

	return t, vertices
}

func forward(t Turtle, d float64, vertices []float64) (Turtle, []float64) {
	return move(t, d, vertices)
}

func move(t Turtle, d float64, vertices []float64) (Turtle, []float64) {

	xn := d*math.Sin(rad(t.angle)) + t.x
	yn := d*math.Cos(rad(t.angle)) + t.y

	vertices = append(vertices, xn)
	vertices = append(vertices, yn)

	t.x = xn
	t.y = yn

	return t, vertices
}

func right(t Turtle, delta float64) Turtle {
	t.angle += delta
	return t
}

func left(t Turtle, delta float64) Turtle {
	t.angle -= delta
	return t
}

func turn(t Turtle, delta float64) Turtle {
	t.angle += delta
	return t
}
