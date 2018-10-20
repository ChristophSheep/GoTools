package main

import (
	"math"
)

func createCircle(xm, ym, r float32, sides uint16) []float32 {

	vertices := []float32{}

	pi2 := 2.0 * math.Pi
	delta := pi2 / float64(sides)

	for alpha := 0.0; alpha < pi2; alpha = alpha + delta {

		x := r*float32(math.Sin(alpha)) + xm
		vertices = append(vertices, x)

		y := r*float32(math.Cos(alpha)) + ym
		vertices = append(vertices, y)
		
	}

	return vertices
}
