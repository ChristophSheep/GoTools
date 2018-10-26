package main // TODO

import (
	"fmt"
	"math"
)

const EPSILON = 0.00001

func equals(x, y float64, epsilon float64) bool {
	diff := x - y
	return math.Abs(diff) < epsilon
}

// calc section of two lines
//
func tryCalcSection(k1, d1, k2, d2 float64) (float64, float64, bool) {

	if equals(k2, k1, EPSILON) {
		return 0.0, 0.0, false
	}

	divisor := k2 - k1
	y := (k2*d1 - k1*d2) / divisor
	x := (y - d2) / k2

	return x, y, true
}

// calc k and d of line from points
//
func calcLine(x1, y1, x2, y2 float64) (float64, float64) {

	dx := x2 - x1
	dy := y2 - y1

	k := dx / dy

	d := y1 - k*x1

	return k, d
}

func calcMiddlePoint(x1, y1, x2, y2 float64) (float64, float64) {

	xm := 0.5 * (x1 + x2)
	ym := 0.5 * (y1 + y2)

	return xm, ym
}

func calcDistance(x1, y1, x2, y2 float64) float64 {

	xx := (x2 - x1) * (x2 - x1)
	yy := (y2 - y1) * (y2 - y1)

	return math.Sqrt(xx + yy)
}

func calcNormalLine(x1, y1, x2, y2 float64) (float64, float64) {

	dx := x2 - x1
	dy := y2 - y1

	kn := -dx / dy

	xm, ym := calcMiddlePoint(x1, y1, x2, y2)

	d := ym - kn*xm

	return kn, d
}

const N = 2 // Dimension 2 = xy, Dimension 3 = xyz

func getXY(vertices []float64, n int) (float64, float64) {

	//  For example:
	//
	//   60 Points by 120 xy coords
	//
	//   ... -2 -1 0 1 2 3 ... -1  0  1  2 ...
	//   ... 58 59 0 1 2 3 ... 59 60 61 62 ...

	count := len(vertices) / N

	if n < 0 {
		n = n + count
	}

	if n >= count {
		n = n - count
	}

	x := vertices[2*n+0]
	y := vertices[2*n+1]

	return x, y
}

func tryCalcMiddlePoint(vertices []float64, n int) (float64, float64, bool) {

	// p1
	// 	\ <--r--Mp
	//   \
	//    p2 ----- p3

	//const N = 2 // x y

	if len(vertices) < 6 {
		return 0.0, 0.0, false
	}

	xl, yl := getXY(vertices, n-1)
	xn, yn := getXY(vertices, n+0)
	xm, ym := getXY(vertices, n+1)

	k1, d1 := calcNormalLine(xl, yl, xn, yn)
	k2, d2 := calcNormalLine(xn, yn, xm, ym)

	return tryCalcSection(k1, d1, k2, d2)
}

func tryCalcRadius(vertices []float64, n int) (float64, bool) {

	xm, ym, ok := tryCalcMiddlePoint(vertices, n)

	// lines parallel -> radius +inf
	if ok == false {
		return math.Inf(1), true
	}

	xn, yn := getXY(vertices, n)
	r := calcDistance(xn, yn, xm, ym)

	return r, true
}

func calcRadi(vertices []float64) []float64 {

	radi := []float64{}
	count := len(vertices) / 2

	for n := 0; n < count; n++ {
		r, ok := tryCalcRadius(vertices, n)
		if ok {
			radi = append(radi, r)
		} else {
			radi = append(radi, math.Inf(+1))
		}
	}

	return radi
}

func test1() {

	k1 := 1.0
	d1 := 0.0

	k2 := -1.0
	d2 := +2.0

	x, y, ok := tryCalcSection(k1, d1, k2, d2)

	if ok {
		fmt.Printf("section x:%f y:%f\n", x, y)
	}
}

func test2() {

	k1 := 3.0 / 2.0
	d1 := 0.0

	k2 := -1.0
	d2 := +2.0

	x, y, ok := tryCalcSection(k1, d1, k2, d2)

	if ok {
		fmt.Printf("section x:%f y:%f\n", x, y)
	}
}

func test3() {

	k1 := 2.0
	d1 := 1.0

	k2 := +2.0
	d2 := -1.0

	x, y, ok := tryCalcSection(k1, d1, k2, d2)

	if ok {
		fmt.Printf("section x:%f y:%f\n", x, y)
	} else {
		fmt.Println("no section - lines parallel")
	}
}

func test4() {

	vertices := createCircle(0.5, 0.5, 0.5, 8)
	x, y, ok := tryCalcMiddlePoint(vertices, 1)

	if ok {
		fmt.Printf("middlepoint x:%f y:%f\n", x, y)
	} else {
		fmt.Println("no middlepoint")
	}
}

func test5() {

	vertices := []float64{
		2.0, 7.0,
		3.0, 4.0,
		8.0, 3.0}

	x, y, ok := tryCalcMiddlePoint(vertices, 1)

	if ok {
		fmt.Printf("middlepoint x: %f y: %f", x, y)
		fmt.Println()
		r, ok := tryCalcRadius(vertices, 1)
		if ok {
			fmt.Printf("radius: %f", r)
			fmt.Println()
			fmt.Printf("kruemmung: %f at x: 2.0, y: 7.0", 1/r)
			fmt.Println()
			v := 50.0  // m/s = 180 km/h
			m := 700.0 // kg
			fg := m * v * v * (1.0 / float64(r))
			fmt.Printf("Fg: %f kN", fg/1000.0)
			fmt.Println()
		}

	} else {
		fmt.Println("no middlepoint")
	}
}

func test6() {

	ovalTrackVerts := createOvalTrack()
	radi := calcRadi(ovalTrackVerts)

	for n := 0; n < len(radi); n++ {
		fmt.Printf("n: %d r:%5.3f \n", n, radi[n])
	}
}

func tests() {

	/*

		test1()
		test2()
		test3()
		test4()
		test5()

	*/
	//test6()
}
