package main

import (
	"math"
)

func newVertices() []float64 {
	return []float64{}
}

func createCircle(xm, ym, r float64, sides uint16) []float64 {

	vertices := []float64{}

	pi2 := 2.0 * math.Pi
	delta := pi2 / float64(sides)

	for alpha := 0.0; alpha < pi2; alpha = alpha + delta {

		x := r*math.Sin(alpha) + xm
		vertices = append(vertices, x)

		y := r*math.Cos(alpha) + ym
		vertices = append(vertices, y)

	}

	return vertices
}

func createOvalTrack() []float64 {

	r := 25.0

	vertices := newVertices()
	t := newTurtle()

	t, vertices = moveTo(t, -r, -r/2.0, vertices)

	// TODO: Formula WRONG
	dAlpha := 2 * deg(1.0*math.Asin(1.0/(2.0*r))) // segment length = 1m

	//fmt.Printf("dAlpha: %f \n", dAlpha)

	// Start

	// 10m straight
	M := 25
	for i := 0; i < M; i++ {
		t, vertices = move(t, 1.0, vertices)
	}

	N := 5
	alpha := 0.0
	delta := dAlpha/float64(N) - 0.04 // Magic correction by trial
	S := int(360.0/dAlpha)/2 - 4      // Magic correction by trial

	// clothoide in
	for i := 0; i < N; i++ {
		alpha += delta
		t = turn(t, alpha)
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
		t = turn(t, alpha)
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
