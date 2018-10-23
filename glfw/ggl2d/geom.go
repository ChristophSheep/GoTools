package main

import (
	"fmt"
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

func calcVec(x1, y1, x2, y2 float64) (float64, float64) {

	vx := x2 - x1
	vy := y2 - y1

	return vx, vy
}

func normalize(vx, vy float64) (float64, float64) {

	length := math.Sqrt(vx*vx + vy*vy)

	vx = vx / length
	vy = vy / length

	return vx, vy
}

func calcNormal(vx, vy float64) (float64, float64) {

	nx := -vy
	ny := vx

	return normalize(nx, ny)
}

func calcNormalVectors(vertices []float64) []float64 {

	count := len(vertices) / N

	normals := []float64{}

	for n := 0; n < count; n++ {

		x1, y1 := getXY(vertices, n-1)
		x2, y2 := getXY(vertices, n+0)

		//            nx,ny
		//  x2,y2  ^ ------>
		//         |
		//         |  vx,vy
		//  x1,y1  |

		vx, vy := calcVec(x1, y1, x2, y2)
		nx, ny := calcNormal(vx, vy)

		normals = append(normals, x2)
		normals = append(normals, y2)

		normals = append(normals, x2+nx)
		normals = append(normals, y2+ny)
	}

	return normals
}

func calcAlpha(s float64, r float64) float64 {

	GK := s / 2.0
	HP := r

	// sin(half_a) = GK / HP

	halfAlphaRad := math.Asin(GK / HP)
	alpha := 2.0 * deg(halfAlphaRad)

	return alpha
}

func createOvalTrack2() []float64 {

	x := -40.0
	y := 0.0
	s := 1.0

	vertices := newVertices()
	t := newTurtle()

	t, vertices = moveTo(t, x, y, vertices)

	S := 20
	N := 41
	M := 5
	alpha := 2.0
	delta := alpha / float64(15) // 15 by design

	//
	// 0° > -80° --> Half circle
	//
	for i := 0; i < N; i++ {
		t, vertices = forward(t, s, vertices)
		t = right(t, alpha)
	}

	fmt.Printf("Alpha after circle: %f \n", t.angle)

	//
	// -80° > -90° --> Clothoide
	//
	sum := 0.0
	for i := 0; i < M; i++ {
		t, vertices = forward(t, s, vertices)
		alpha = alpha - delta
		sum += alpha
		t = turn(t, alpha)
	}

	fmt.Printf("Sum Alpha after clothoide: %f \n", sum)

	//
	// -90°, 20m --> Straight
	//
	for i := 0; i < S; i++ {
		t, vertices = forward(t, s, vertices)
	}

	//
	// -90° > -80° --> Clothoide
	//
	for i := 0; i < M; i++ {
		t, vertices = forward(t, s, vertices)
		alpha = alpha + delta
		t = turn(t, alpha)
	}

	fmt.Printf("Alpha after clothoide: %f \n", alpha)

	for i := 0; i < N; i++ {
		t, vertices = forward(t, s, vertices)
		t = right(t, alpha)
	}

	fmt.Printf("Angle after circle: %f \n", t.angle)

	//
	// -80° > -180° -> circle 40m
	//

	return vertices
}

func createOvalTrack() []float64 {

	r := 25.0 // radius of circle in meter
	s := 1.0  // single segment length in meter

	vertices := newVertices()
	t := newTurtle()

	t, vertices = moveTo(t, -r, -r/2.0, vertices)

	dAlpha := calcAlpha(s, r)

	N := 5                            // Clothoide Segments
	alpha := 0.0                      // Start Angle
	delta := dAlpha/float64(N) - 0.04 // Magic Correction by trial
	S := int(360.0/dAlpha)/2 - 4      // Magic Correction by trial
	M := 25                           // Straight Segments a 1m

	fmt.Printf("clothoide delta: %f \n", delta)
	fmt.Printf("circle dAlpha: %f \n", dAlpha)

	// straight
	for i := 0; i < M; i++ {
		t, vertices = move(t, s, vertices)
	}

	// clothoide in
	for i := 0; i < N; i++ {
		alpha += delta
		t = turn(t, alpha)
		t, vertices = move(t, s, vertices)
	}

	// circle bow
	for i := 0; i < S; i++ {
		t = turn(t, dAlpha)
		t, vertices = move(t, s, vertices)
	}

	// clothoide out
	for i := 0; i < N; i++ {
		alpha -= delta
		t = turn(t, alpha)
		t, vertices = move(t, s, vertices)
	}

	// straight
	for i := 0; i < M; i++ {
		t, vertices = move(t, s, vertices)
	}

	// clothoide in
	for i := 0; i < N; i++ {
		alpha += delta
		t = turn(t, alpha)
		t, vertices = move(t, s, vertices)
	}

	// circle bow
	for i := 0; i < S; i++ {
		t = turn(t, dAlpha)
		t, vertices = move(t, s, vertices)
	}

	N-- // MAGIC CORRECTION !!TODO!!

	// clothoide out
	for i := 0; i < N; i++ {
		alpha -= delta
		t = turn(t, alpha)
		t, vertices = move(t, s, vertices)
	}

	return vertices
}
