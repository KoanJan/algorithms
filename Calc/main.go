package main

import (
	"fmt"
	"math"
)

func main() {
	// test
	var (
		a1 int = -591681
		b1 int = 591681
	)
	fmt.Printf("%d + %d = %d\t", a1, b1, Add(a1, b1))
	fmt.Printf("check answer: %d\n", a1+b1)

	var (
		a2 int = 59666
		b2 int = 59060
	)
	fmt.Printf("%d - %d = %d\t", a2, b2, Minus(a2, b2))
	fmt.Printf("check answer: %d\n", a2-b2)

	var (
		a3 int = 15
		b3 int = 15
	)
	fmt.Printf("%d * %d = %d\t", a3, b3, Multi(a3, b3))
	fmt.Printf("check answer: %d\n", a3*b3)

	var (
		a4 int = 1862312
		b4 int = 64845
	)
	fmt.Printf("%d / %d = %d\t", a4, b4, Div(a4, b4))
	fmt.Printf("check answer: %d\n", a4/b4)

	var (
		a5 int = 8
		b5 int = 3
	)
	fmt.Printf("pow(%d, %d) = %d\t", a5, b5, Pow(a5, b5))
	fmt.Printf("check answer: %d\n", int(math.Pow(float64(a5), float64(b5))))
}
