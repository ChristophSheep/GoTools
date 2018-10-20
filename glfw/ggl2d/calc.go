package main // TODO

import (
	"fmt"
	"math"
)

const EPSILON = 0.00001

func equals(x, y float32, epsilon float64) bool {
	diff := float64(x - y)
	return math.Abs(diff) < epsilon
}

// calc section of two lines
//
func tryCalcSection(k1, d1, k2, d2 float32) (float32, float32, bool) {

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
func calcLine(x1, y1, x2, y2 float32) (float32, float32) {

	dx := x2 - x1
	dy := y2 - y1

	k := dx / dy

	d := y1 - k*x1

	return k, d
}

func calcMiddlePoint(x1, y1, x2, y2 float32) (float32, float32) {

	xm := 0.5 * (x1 + x2)
	ym := 0.5 * (y1 + y2)

	return xm, ym
}

func calcNormalLine(x1, y1, x2, y2 float32) (float32, float32) {

	dx := x2 - x1
	dy := y2 - y1

	kn := -dx / dy

	xm, ym := calcMiddlePoint(x1, y1, x2, y2)

	d := ym - kn*xm

	return kn, d
}

func tryCalcMiddlePoint(vertices []float32) (float32, float32, bool) {

	// p1
	// 	\      Mp
	//   \
	//    p2 ----- p3

	//const N = 2 // x y

	if len(vertices) < 6 {
		return 0.0, 0.0, false
	}
	x1 := vertices[0]
	y1 := vertices[1]

	x2 := vertices[2]
	y2 := vertices[3]

	x3 := vertices[4]
	y3 := vertices[5]

	k1, d1 := calcNormalLine(x1, y1, x2, y2)
	k2, d2 := calcNormalLine(x2, y2, x3, y3)

	return tryCalcSection(k1, d1, k2, d2)
}

func test1() {

	k1 := float32(1.0)
	d1 := float32(0.0)

	k2 := float32(-1.0)
	d2 := float32(2.0)

	x, y, ok := tryCalcSection(k1, d1, k2, d2)

	if ok {
		fmt.Printf("section x:%f y:%f\n", x, y)
	}
}

func test2() {

	k1 := float32(3.0 / 2.0)
	d1 := float32(0.0)

	k2 := float32(-1.0)
	d2 := float32(2.0)

	x, y, ok := tryCalcSection(k1, d1, k2, d2)

	if ok {
		fmt.Printf("section x:%f y:%f\n", x, y)
	}
}

func test3() {

	k1 := float32(2.0)
	d1 := float32(1.0)

	k2 := float32(2.0)
	d2 := float32(-1.0)

	x, y, ok := tryCalcSection(k1, d1, k2, d2)

	if ok {
		fmt.Printf("section x:%f y:%f\n", x, y)
	} else {
		fmt.Println("no section - lines parallel")
	}
}

func test4() {

	vertices := createCircle(0.5, 0.5, 0.5, 8)
	x, y, ok := tryCalcMiddlePoint(vertices)

	if ok {
		fmt.Printf("middlepoint x:%f y:%f\n", x, y)
	} else {
		fmt.Println("no middlepoint")
	}
}

func test5() {

	vertices := []float32{
		2.0, 7.0,
		3.0, 4.0,
		8.0, 3.0}

	x, y, ok := tryCalcMiddlePoint(vertices)

	if ok {
		fmt.Printf("middlepoint x:%f y:%f\n", x, y)
	} else {
		fmt.Println("no middlepoint")
	}
}

func test() {
	test1()
	test2()
	test3()
	test4()
	test5()
}
