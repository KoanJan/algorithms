package main

import (
	"fmt"
)

func EuclidGCD(a, b int64) int64 {
	if b == 0 || b == a {
		return a
	}
	if a > b {
		return EuclidGCD(b, a%b)
	} else {
		return EuclidGCD(a, b%a)
	}
}

func main() {
	// test
	var (
		a int64 = 10248888413810
		b int64 = 20488888035480
	)
	fmt.Printf("EuclidGCD(%d, %d) is %d", a, b, EuclidGCD(a, b))
}
