package main // TODO

import (
	"math"
)

// ----------------------------------------------------------------------------

const N = 2 // Dimension 2 = xy, Dimension 3 = xyz

const EPSILON = 0.00001 // Epsilon for comparing

// ----------------------------------------------------------------------------

func equalsWithEpsilon(x, y float64, epsilon float64) bool {
	diff := x - y
	return math.Abs(diff) < epsilon
}

/*
 Calc section of two lines
 given k, d of line
*/
func tryCalcSection(k1, d1, k2, d2 float64) (float64, float64, bool) {

	// Check if lines parallel

	if equalsWithEpsilon(k2, k1, EPSILON) {
		return math.Inf(1), math.Inf(1), false
	}

	// TODO: one line is vertical
	//
	//
	//    | k = inf
	// ---+---- k = 0
	//    |
	//

	divisor := k2 - k1

	y := (k2*d1 - k1*d2) / divisor
	x := (y - d2) / k2

	if k2 == 0.0 {
		x = (y - d1) / k1
	}

	return x, y, true
}

/*
 Calc k and d of a line from two points
*/
func calcLine(x1, y1, x2, y2 float64) (float64, float64) {

	dx := x2 - x1
	dy := y2 - y1

	k := dx / dy

	d := y1 - k*x1

	return k, d
}

/*
  Calc middle of a line given by 2 points
*/
func calcMiddlePoint(x1, y1, x2, y2 float64) (float64, float64) {

	xm := 0.5 * (x1 + x2)
	ym := 0.5 * (y1 + y2)

	return xm, ym
}

/*
  Calc phi of vector vx,vy
*/
func calcPhi(vx, vy float64) float64 {
	return math.Atan(vx / vy)
}

/*
  Calc distance of 2 points
*/
func calcDistance(x1, y1, x2, y2 float64) float64 {

	xx := (x2 - x1) * (x2 - x1)
	yy := (y2 - y1) * (y2 - y1)

	return math.Sqrt(xx + yy)
}

/*
  Calc normal line (streckensymetrale) of 2 points
*/
func calcLineSymetrale(x1, y1, x2, y2 float64) (float64, float64) {

	dx := x2 - x1
	dy := y2 - y1

	kn := -dx / dy

	xm, ym := calcMiddlePoint(x1, y1, x2, y2)

	d := ym - kn*xm

	return kn, d
}

/*
 Calc normal vector of given vector vx vy
*/
func calcNormalVector(vx, vy float64) (float64, float64) {

	nx := -vy
	ny := vx

	return normalize(nx, ny)
}

/*
 Calc vector if given two points x1 y1 -> x2 y2
*/
func calcVec(x1, y1, x2, y2 float64) (float64, float64) {

	vx := x2 - x1
	vy := y2 - y1

	return vx, vy
}

/*
 Normaize given vector to length 1.0
*/
func normalize(vx, vy float64) (float64, float64) {

	length := math.Sqrt(vx*vx + vy*vy)

	vx = vx / length
	vy = vy / length

	return vx, vy
}

/*
 Scale vector with factor x and y
*/
func scale(vx, vy, scaleX, scaleY float64) (float64, float64) {
	return vx * scaleX, vy * scaleY
}

/*
 Get x y of list of vertices xy xy xy ..
*/
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

/*
 Calc middle point of 2 lines p1 -> p2, p2 -> p3
*/
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

	k1, d1 := calcLineSymetrale(xl, yl, xn, yn)
	k2, d2 := calcLineSymetrale(xn, yn, xm, ym)

	return tryCalcSection(k1, d1, k2, d2)
}

/*
 Calc radius of 2 lines
*/
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

/*
 Calc middle point of arc or clothoide or line
	- arc has a middle point and radius
	- clothoide each segments has middle point
	- line has a middle point -> infinity
*/
func calcMiddlePointAndRadiAndVectors(vertices []float64) ([]float64, []float64, []float64) {

	middlePoints := []float64{}
	radi := []float64{}
	centrifugalVectors := []float64{}

	count := len(vertices) / N

	for n := 0; n < count; n++ {

		xl, yl := getXY(vertices, n-1)
		xn, yn := getXY(vertices, n+0)
		xu, yu := getXY(vertices, n+1)

		// Strecken Symmetrale
		// Straight line symetric line
		//
		k1, d1 := calcLineSymetrale(xl, yl, xn, yn) // TODO: Rename "Streckensymmetrale"
		k2, d2 := calcLineSymetrale(xn, yn, xu, yu)

		xm, ym, ok := tryCalcSection(k1, d1, k2, d2)

		//fmt.Printf("n: %d - xm: %f, ym: %f \n", n, xm, ym)

		cvx := 0.0
		cvy := 0.0
		r := 0.0

		// if section lines parallel -> r inf, cv are inf
		//
		if ok {
			r = calcDistance(xn, yn, xm, ym)
			//cvx, cvy = normalize(calcVec(xm, ym, xn, yn))
			cvx, cvy = calcVec(xn, yn, xm, ym)

			//fmt.Printf("n: %d - cvx: %f, cvy: %f \n", n, cvx, cvy)

		} else {
			xm = math.Inf(1)
			ym = math.Inf(1)
			r = math.Inf(1)
		}

		// TODO: Performance append - faster method
		// create a array with fixed size and work with foo[i]
		//
		middlePoints = append(middlePoints, xm)
		middlePoints = append(middlePoints, ym)

		radi = append(radi, r)

		centrifugalVectors = append(centrifugalVectors, cvx)
		centrifugalVectors = append(centrifugalVectors, cvy)
	}

	return middlePoints, radi, centrifugalVectors
}

/*
  Calc all radi of a polyline
*/
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
