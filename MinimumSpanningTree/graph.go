package main

import (
	"bytes"
	"fmt"
)

const (
	Unconnected int = 1 << 16
)

type Graph [][]int

func NewGraph(size int) Graph {
	graph := Graph{}
	for i := 0; i < size; i++ {
		col := []int{}
		for j := 0; j < size; j++ {
			col = append(col, Unconnected)
		}
		graph = append(graph, col)
	}
	return graph
}

func (g *Graph) Size() int {
	return len([][]int(*g))
}

func (g *Graph) String() string {
	buff := bytes.NewBufferString("")
	for i := 0; i < g.Size(); i++ {
		for j := 0; j < g.Size(); j++ {
			if (*g)[i][j] == Unconnected {
				buff.WriteString("  ∞  ")
			} else {
				v := (*g)[i][j]
				if v > 9 {
					buff.WriteString(fmt.Sprintf(" %d  ", v))
				} else {
					buff.WriteString(fmt.Sprintf("  %d  ", v))
				}
			}
		}
		buff.WriteString("\n")
	}
	return buff.String()
}

func NewRandGraph() Graph {
	g := NewGraph(RandIntn(10) + 3)
	for i := 0; i < g.Size(); i++ {
		for j := 0; j < g.Size(); j++ {
			if i == j {
				// circle
				g[i][j] = 0
			} else {
				g[i][j] = RandIntn(99) + 1
			}
		}
	}
	return g
}
