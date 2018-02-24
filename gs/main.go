package main

import (
	"fmt"
)

func main() {
	// test

	// BFS
	graph1 := [][]bool{
		[]bool{true, true, false, true, false},
		[]bool{true, true, true, false, true},
		[]bool{false, true, true, false, false},
		[]bool{true, false, false, true, true},
		[]bool{false, true, false, true, true},
	}
	distances1, parents1 := BFS(graph1, 0)
	fmt.Printf("BFS test: \n\tdistances1: %v\n\tparents1: %v\n", distances1, parents1)

	// DFS
	graph2 := [][]bool{
		[]bool{true, true, true, true, true},
		[]bool{true, true, true, true, true},
		[]bool{true, true, true, true, true},
		[]bool{true, true, true, true, true},
		[]bool{true, true, true, true, true},
	}
	parents2 := DFS(graph2, 0)
	fmt.Printf("DFS test: \n\tparents2: %v\n", parents2)
}
