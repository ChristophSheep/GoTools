package main

import (
	"fmt"
	"math"
)

type turtle struct {
	x     float32
	y     float32
	angle float32
}

func newTurtle() turtle {
	return turtle{
		x:     0.0,
		y:     0.0,
		angle: 0.0}
}

func rad(deg float32) float64 {
	return float64(deg/180.0) * math.Pi
}

func deg(rad float64) float32 {
	return float32((rad / math.Pi) * 180.0)
}

func moveTo(t turtle, x float32, y float32, vertices []float32) (turtle, []float32) {

	vertices = append(vertices, x)
	vertices = append(vertices, y)

	t.x = x
	t.y = y

	return t, vertices
}

func move(t turtle, d float32, vertices []float32) (turtle, []float32) {

	xn := d*float32(math.Sin(rad(t.angle))) + t.x
	yn := d*float32(math.Cos(rad(t.angle))) + t.y

	vertices = append(vertices, xn)
	vertices = append(vertices, yn)

	fmt.Println(xn, yn)

	t.x = xn
	t.y = yn

	return t, vertices
}

func turn(t turtle, delta float32) turtle {
	t.angle += delta
	return t
}

func createHalfOvalTrack() []float32 {

	r := 25.0

	vertices := []float32{}
	t := newTurtle()

	t, vertices = moveTo(t, float32(-r), float32(-r/2.0), vertices)

	// TODO: Formula WRONG
	dAlpha := 2 * deg(1.0*math.Asin(1.0/(2.0*r))) // segment length = 1m

	fmt.Printf("dAlpha: %f \n", dAlpha)

	// Start

	// 10m straight
	M := 25
	for i := 0; i < M; i++ {
		t, vertices = move(t, 1.0, vertices)
	}

	N := 5
	alpha := float32(0.0)
	delta := dAlpha/float32(N) - 0.07 // Magic correction
	S := int(360.0/dAlpha)/2 - 2      // Magic correction

	// clothoide in
	for i := 0; i < N; i++ {
		alpha += delta
		t = turn(t, delta)
		t, vertices = move(t, 1.0, vertices)
	}

	// circle bow

	for i := 0; i < S; i++ {
		t = turn(t, dAlpha)
		t, vertices = move(t, 1.0, vertices)
	}

	// clothoide out
	for i := 0; i < N; i++ {
		alpha -= delta
		t = turn(t, alpha)
		t, vertices = move(t, 1.0, vertices)
	}

	// straight
	for i := 0; i < M; i++ {
		t, vertices = move(t, 1.0, vertices)
	}

	// clothoide in
	for i := 0; i < N; i++ {
		alpha += delta
		t = turn(t, delta)
		t, vertices = move(t, 1.0, vertices)
	}

	// circle bow
	for i := 0; i < S; i++ {
		t = turn(t, dAlpha)
		t, vertices = move(t, 1.0, vertices)
	}

	// clothoide out
	for i := 0; i < N; i++ {
		alpha -= delta
		t = turn(t, alpha)
		t, vertices = move(t, 1.0, vertices)
	}

	return vertices
}
