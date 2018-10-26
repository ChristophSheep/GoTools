package main

/*
Model create objects of the world
It is a model of the virtual world.
*/

import (
	"fmt"
	"math"
)

func createVertices() []float64 {
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

func arc(isRight bool, steps int, s float64, k float64, vs []float64, t Turtle) (Turtle, []float64) {
	for step := 0; step < steps; step++ {
		t, vs = forward(t, s, vs)
		if isRight {
			t = right(t, k)
		} else {
			t = left(t, k)
		}
	}
	return t, vs
}

func clothoide(isRight bool, in bool, steps int, s float64, k float64, vs []float64, t Turtle) (Turtle, []float64) {

	dk := k / float64(steps)
	ck := 0.0

	if in {
		ck = 0.0
	} else {
		ck = k
	}

	for step := 0; step < steps; step++ {
		t, vs = forward(t, s, vs)

		if in {
			ck = ck + dk
		} else {
			ck = ck - dk
		}

		//fmt.Printf("clothoide ck:%f", ck)

		if isRight {
			t = right(t, ck)
		} else {
			t = left(t, ck)
		}
	}
	return t, vs
}

func line(steps int, s float64, vs []float64, t Turtle) (Turtle, []float64) {
	for step := 0; step < steps; step++ {
		t, vs = forward(t, s, vs)
	}
	return t, vs
}

func createAnyTrack() []float64 {

	vs := createVertices()
	t := createTurtle()

	x := -20.0 // Start point
	y := -30.0 // Start point
	t, vs = moveTo(t, x, y, vs)

	s := 1.0 // Step length 1m
	k := 2.0 // Krümmung alpha 2 Grad

	as := 125 // Arc segements
	ls := 55  // Line segments
	cs := 10  // Clothoide segments

	in := true
	out := false
	isRight := true
	isLeft := false

	//t, vs = arc(isRight, as, s, k, vs, t)
	//t, vs = clothoide(isRight, out, cs, s, k, vs, t)
	t, vs = line(ls+1, s, vs, t)
	t, vs = clothoide(isRight, in, cs, s, k, vs, t)
	t, vs = arc(isRight, as, s, k, vs, t)
	t, vs = clothoide(isRight, out, cs, s, k, vs, t)
	t, vs = line(ls-8, s, vs, t)
	t, vs = clothoide(isLeft, in, cs, s, k, vs, t)
	t, vs = arc(isLeft, as, s, k, vs, t)

	return vs
}

func createOvalTrack2() []float64 {

	x := -40.0
	y := 0.0
	s := 1.0

	vertices := createVertices()
	t := createTurtle()

	t, vertices = moveTo(t, x, y, vertices)

	S := 20
	N := 41
	M := 5
	alpha := 2.0
	delta := alpha / float64(15) // 15 by design

	//
	// 0° > -80° --> Half circle
	//

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

	vertices := createVertices()
	t := createTurtle()

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
		t, vertices = forward(t, s, vertices)
	}

	// clothoide in
	for i := 0; i < N; i++ {
		alpha += delta
		t = turn(t, alpha)
		t, vertices = forward(t, s, vertices)
	}

	// circle bow
	for i := 0; i < S; i++ {
		t = turn(t, dAlpha)
		t, vertices = forward(t, s, vertices)
	}

	// clothoide out
	for i := 0; i < N; i++ {
		alpha -= delta
		t = turn(t, alpha)
		t, vertices = forward(t, s, vertices)
	}

	// straight
	for i := 0; i < M; i++ {
		t, vertices = forward(t, s, vertices)
	}

	// clothoide in
	for i := 0; i < N; i++ {
		alpha += delta
		t = turn(t, alpha)
		t, vertices = forward(t, s, vertices)
	}

	// circle bow
	for i := 0; i < S; i++ {
		t = turn(t, dAlpha)
		t, vertices = forward(t, s, vertices)
	}

	// clothoide out
	for i := 0; i < N; i++ {
		alpha -= delta
		t = turn(t, alpha)
		t, vertices = forward(t, s, vertices)
	}

	return vertices
}
