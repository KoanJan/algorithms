package main

import (
	"fmt"
)

// test
func main() {

	// data
	str1 := "FOFF"
	str2 := "FOCF"

	// lcs
	c, b := LCS(str1, str2)

	fmt.Printf("c: \n\n")
	for _, _c := range c {
		fmt.Printf("%v\n", _c)
	}

	fmt.Printf("\nb: \n\n")
	for _, _b := range b {
		fmt.Printf("%v\n", _b)
	}

}
