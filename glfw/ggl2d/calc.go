package main // TODO

import (
	"fmt"
)

//  calc section of two lines
//
func tryCalcSection(k1, d1, k2, d2 float32) (float32, float32, bool) {

	if k2-k1 == 0 {
		return 0.0, 0.0, false // TODO undefined
	}

	y := (k2*d1 - k1*d2) / (k2 - k1)
	x := (y - d2) / k2

	return x, y, true
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
		fmt.Println("no section")
	}
}

func test() {
	test1()
	test2()
	test3()
}
