package main

/*

Turtle is used to create vertices along a path
for instance a cirlce or clothoide or a straight line.

*/

import (
	"math"
)

type Turtle struct {
	x     float64
	y     float64
	angle float64
}

func createTurtle() Turtle {
	return Turtle{
		x:     0.0,
		y:     0.0,
		angle: 0.0}
}

// TODO: This helpers must exists somewhere

func rad(deg float64) float64 {
	return (deg / 180.0) * math.Pi
}

func deg(rad float64) float64 {
	return (rad / math.Pi) * 180.0
}

/*
	Move to position the turtle at the given coordinates
	and create the position in vertices
*/
func moveTo(t Turtle, x float64, y float64, vs []float64) (Turtle, []float64) {

	vs = append(vs, x)
	vs = append(vs, y)

	t.x = x
	t.y = y

	return t, vs
}

/*
	Move the turtle the distance d forward
	and add this new position to the vertices list
*/
func forward(t Turtle, d float64, vs []float64) (Turtle, []float64) {

	xn := d*math.Sin(rad(t.angle)) + t.x
	yn := d*math.Cos(rad(t.angle)) + t.y

	vs = append(vs, xn)
	vs = append(vs, yn)

	t.x = xn
	t.y = yn

	return t, vs
}

/*
	Turn the turtle right with given delta angle
*/
func right(t Turtle, delta float64) Turtle {
	return turn(t, +delta)
}

/*
	Turn the turtke left with given delta angle
*/
func left(t Turtle, delta float64) Turtle {
	return turn(t, -delta)
}

/*
	Turn right with position delta and
	turn left with negative delta
*/
func turn(t Turtle, delta float64) Turtle {
	t.angle += delta
	return t
}
