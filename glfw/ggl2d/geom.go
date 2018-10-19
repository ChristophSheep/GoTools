package main

import (
	"math"
)

func createCircle(x, y, r float32, count uint16) []float32 {

	vertices := []float32{}
	delta := 2.0 * math.Pi / float64(count)

	for alpha := 0.0; alpha < math.Pi*2; alpha = alpha + delta {
		x := float64(r) * math.Sin(alpha)
		y := float64(r) * math.Cos(alpha)
		vertices = append(vertices, float32(x))
		vertices = append(vertices, float32(y))
	}

	return vertices
}
