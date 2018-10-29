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

func createTrackArea(vertices []float64, width float64) []float64 {

	areaVertices := []float64{}

	//  lx,ly                rx,ry
	//  -----                -----
	//              |
	//   p4 --------+-------- p5
	// 	  		    |
	//   p2 --------+-------- p3
	// 			    |
	//   p0 --------+-------- p1
	//     		    |

	halfWidth := width / 2.0
	nvectors := calcNormalVectors(vertices)
	count := len(nvectors) / N

	for n := 0; n < count; n++ {

		vx, vy := getXY(nvectors, n)

		//fmt.Printf("n: %d, vx: %f,  vy: %f \n", n, vx, vy)

		xn, yn := getXY(vertices, n)

		rx := vx * halfWidth
		ry := vy * halfWidth

		lx := -1.0 * vx * halfWidth
		ly := -1.0 * vy * halfWidth

		areaVertices = append(areaVertices, xn+lx)
		areaVertices = append(areaVertices, yn+ly)

		areaVertices = append(areaVertices, xn+rx)
		areaVertices = append(areaVertices, yn+ry)

	}
	return areaVertices
}

func calcVelocityVectors(vertices []float64, velocityData []float64, scaleFactor float64) []float64 {
	result := []float64{}

	count := len(vertices) / N

	for n := 0; n < count; n++ {
		x1, y1 := getXY(vertices, n-1)
		xn, yn := getXY(vertices, n+0)

		vx, vy := calcVec(x1, y1, xn, yn)
		nx, ny := calcNormalVector(vx, vy)

		velocity := 0.0 // unit factor
		if n < len(velocityData) {
			velocity = velocityData[n]
			fmt.Printf("n: %d, velocity: %f \n", n, velocity)

		}

		scaleX := velocity * scaleFactor
		scaleY := velocity * scaleFactor

		vvx, vvy := scale(nx, ny, -scaleX, -scaleY)

		result = append(result, xn)
		result = append(result, yn)

		result = append(result, xn+vvx)
		result = append(result, yn+vvy)

	}

	return result
}

func calcCentrifugalVectors(vertices []float64, velocityData []float64, scaleFactor float64) []float64 {

	result := []float64{}

	_, radi, vectors := calcMiddlePointAndRadiAndVectors(vertices)

	//_, _, vectors := calcMiddlePointAndRadiAndVectors(vertices)

	count := len(vectors) / N

	for n := 0; n < count; n++ {

		velocity := 0.0 // unit factor
		if n < len(velocityData) {
			velocity = velocityData[n]
			fmt.Printf("n: %d, velocity: %f \n", n, velocity)

		}

		vx, vy := getXY(vectors, n)

		// Centrifugal vectors
		// are opposite to vector from point to (arc) middle point
		vx *= -1.0
		vy *= -1.0

		//
		// TODO: Fg = m * v^2 * k
		//

		radius := radi[n] // radius = inf(1) -> k = 0
		k := 1.0 / radius

		m := 1.0

		// TODO: ONLY SHOW SPEED
		Fg := m * velocity * velocity * k

		scaleX := Fg * scaleFactor
		scaleY := Fg * scaleFactor

		vx, vy = normalize(vx, vy)
		vx, vy = scale(vx, vy, scaleX, scaleY)

		xn, yn := getXY(vertices, n)

		result = append(result, xn)
		result = append(result, yn)

		result = append(result, xn+vx)
		result = append(result, yn+vy)
	}

	return result
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
		nx, ny := calcNormalVector(vx, vy)

		normals = append(normals, nx)
		normals = append(normals, ny)

		//normals = append(normals, x2)
		//normals = append(normals, y2)

		//normals = append(normals, x2+nx)
		//normals = append(normals, y2+ny)
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

// ----------------------------------------------------------------------------
// Segment types:
//
//     - line
//     - arc
//     - clothoide
//

const (
	IN    = true
	OUT   = false
	RIGHT = true
	LEFT  = false
)

func arc(steps int, s float64, isRight bool, k float64, vs []float64, t Turtle) (Turtle, []float64) {
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

func clothoide(steps int, s float64, isRight bool, in bool, k float64, vs []float64, t Turtle) (Turtle, []float64) {

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

func _speedUp_(steps int, s float64, a float64, vv []float64, tv TurtleV) (TurtleV, []float64) {
	tv = setAcceleration(tv, a)
	for step := 0; step < steps; step++ {
		tv, vv = drive(tv, s, vv)
	}
	return tv, vv
}

func _breakDown_(steps int, s float64, a float64, vv []float64, tv TurtleV) (TurtleV, []float64) {
	tv = setAcceleration(tv, -a)
	for step := 0; step < steps; step++ {
		tv, vv = drive(tv, s, vv)
	}
	return tv, vv
}

func createTurtleVerts() (Turtle, []float64) {
	vs := createVertices()
	t := createTurtle()
	return t, vs
}

func createTurtleVData() (TurtleV, []float64) {
	vs := []float64{}
	tv := createTurtleV()
	return tv, vs
}

func createAnyTrackSmall() []float64 {

	t, vs := createTurtleVerts()

	x := -0.0 // Start point
	y := -0.0 // Start point
	t, vs = moveTo(t, x, y, vs)

	s := 4.0 // s .. step length 1m
	k := 4.0 // k .. krümmung alpha 2 Grad after s 1m

	//t, vs = clothoide(5, s, RIGHT, IN, k, vs, t)
	t, vs = arc(2, s, RIGHT, k, vs, t)

	return vs
}

func createAnyTrack() ([]float64, []float64) {

	t, vs := createTurtleVerts()
	tv, vv := createTurtleVData()

	s0 := 0.0   // meter
	v0 := 250.0 // km/h
	a := 5.0    // acceleration km/h per meter
	b := 10.0   // break km/h per meter
	tv = setAbsolute(tv, s0, a, v0)

	x := -20.0 // Start point
	y := -10.0 // Start point
	t, vs = moveTo(t, x, y, vs)

	s := 1.0 // s .. step length 1m
	k := 4.0 // k .. krümmung alpha 2 Grad after s 1m

	tv, vv = _speedUp_(6, s, a, vv, tv)

	t, vs = line(5, s, vs, t)

	tv, vv = _breakDown_(10, s, b, vv, tv)

	t, vs = clothoide(15, s, RIGHT, IN, k, vs, t)
	t, vs = arc(25, s, RIGHT, k, vs, t)
	t, vs = clothoide(5, s, RIGHT, OUT, k, vs, t)
	t, vs = clothoide(5, s, LEFT, IN, k, vs, t)
	t, vs = arc(55, s, LEFT, k, vs, t)
	t, vs = clothoide(5, s, LEFT, OUT, k, vs, t)
	t, vs = line(41, s, vs, t)
	t, vs = clothoide(5, s, LEFT, IN, k, vs, t)
	t, vs = arc(15, s, LEFT, k, vs, t)
	t, vs = clothoide(5, s, LEFT, OUT, k, vs, t)
	t, vs = line(33, s, vs, t)
	t, vs = clothoide(5, s, LEFT, IN, k+2, vs, t)
	t, vs = arc(10, s, LEFT, k+2, vs, t)
	t, vs = clothoide(5, s, LEFT, OUT, k+2, vs, t)
	t, vs = line(30, s, vs, t)
	t, vs = clothoide(5, s, LEFT, IN, k+4, vs, t)
	t, vs = arc(18, s, LEFT, k+4, vs, t)
	t, vs = clothoide(5, s, LEFT, OUT, k+4, vs, t)
	t, vs = line(13, s, vs, t)
	t, vs = clothoide(5, s, RIGHT, IN, k+6, vs, t)
	t, vs = arc(4, s, RIGHT, k+6, vs, t)
	t, vs = clothoide(5, s, RIGHT, OUT, k+6, vs, t)
	t, vs = line(5, s, vs, t)

	return vs, vv
}

func createAnyTrackIdealLine() []float64 {

	t, vs := createTurtleVerts()

	x := -24.0 // Start point
	y := -10.0 // Start point
	t, vs = moveTo(t, x, y, vs)

	s := 1.0 // s .. step length 1m
	k := 4.0 // k .. krümmung alpha 2 Grad after s 1m

	t, vs = line(5, s, vs, t)
	t, vs = clothoide(15, s, RIGHT, IN, k, vs, t)
	t, vs = arc(23, s, RIGHT, k, vs, t)
	t, vs = clothoide(5, s, RIGHT, OUT, k, vs, t)
	t, vs = line(8, s, vs, t)
	t, vs = clothoide(5, s, LEFT, IN, k, vs, t)
	t, vs = arc(62, s, LEFT, k-0.6, vs, t)
	t, vs = clothoide(5, s, LEFT, OUT, k, vs, t)
	t, vs = line(29, s, vs, t)
	t, vs = clothoide(5, s, LEFT, IN, k-2.0, vs, t)
	t, vs = arc(32, s, LEFT, k-2.0, vs, t)
	t, vs = clothoide(5, s, LEFT, OUT, k-2.0, vs, t)
	t, vs = line(21, s, vs, t)
	t, vs = clothoide(5, s, LEFT, IN, k-0.5, vs, t)
	t, vs = arc(22, s, LEFT, k-0.5, vs, t)
	t, vs = clothoide(5, s, LEFT, OUT, k-0.5, vs, t)
	t, vs = line(26, s, vs, t)
	t, vs = clothoide(5, s, LEFT, IN, k+1.7, vs, t)
	t, vs = arc(32, s, LEFT, k+1.7, vs, t)
	t, vs = clothoide(5, s, LEFT, OUT, k+1.7, vs, t)
	t, vs = line(6, s, vs, t)
	t, vs = clothoide(5, s, RIGHT, IN, k+2.5, vs, t)
	t, vs = arc(12, s, RIGHT, k+2.5, vs, t)
	t, vs = clothoide(5, s, RIGHT, OUT, k+2.5, vs, t)
	t, vs = line(5, s, vs, t)

	return vs
}

func createEightTrack() []float64 {

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

	//t, vs = arc(isRight, as, s, k, vs, t)
	//t, vs = clothoide(isRight, out, cs, s, k, vs, t)
	t, vs = line(ls+1, s, vs, t)
	t, vs = clothoide(cs, s, RIGHT, IN, k, vs, t)
	t, vs = arc(as, s, RIGHT, k, vs, t)
	t, vs = clothoide(cs, s, RIGHT, OUT, k, vs, t)
	t, vs = line(ls-8, s, vs, t)
	t, vs = clothoide(cs, s, LEFT, IN, k, vs, t)
	t, vs = arc(as, s, LEFT, k, vs, t)

	return vs
}
