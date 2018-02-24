package main

import (
	"fmt"
)

func main() {
	// test
	g := NewRandGraph()
	fmt.Println("Raw graph:")
	fmt.Println(g.String())

	tree := Prim(g)
	fmt.Println("Spanning tree:")
	fmt.Println(tree.String())
}
