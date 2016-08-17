package main

import (
	"fmt"
)

// test
func main() {
	a := NewPoint("a")
	b := NewPoint("b")
	c := NewPoint("c")
	d := NewPoint("d")
	e := NewPoint("e")
	f := NewPoint("f")
	g := NewPoint("g")
	h := NewPoint("h")

	points := []*Point{a, b, c, d, e, f, g, h}

	edges := map[string]map[string]int{}
	edges[a.Name] = map[string]int{c.Name: 1, d.Name: 2, f.Name: 3}
	edges[b.Name] = map[string]int{c.Name: 2, e.Name: 3, g.Name: 6}
	edges[c.Name] = map[string]int{b.Name: 2, e.Name: 4, h.Name: 1}
	edges[d.Name] = map[string]int{a.Name: 1, g.Name: 1}
	edges[e.Name] = map[string]int{b.Name: 5, c.Name: 2, d.Name: 1, h.Name: 4}
	edges[f.Name] = map[string]int{a.Name: 1, c.Name: 2, g.Name: 1}
	edges[g.Name] = map[string]int{b.Name: 1, c.Name: 5, h.Name: 1}
	edges[h.Name] = map[string]int{a.Name: 1, d.Name: 5, e.Name: 9, f.Name: 1}

	graph := NewGraph(points, edges)
	fmt.Printf("graph points: %v\n", graph.Points)
	fmt.Printf("graph edges: %v\n", graph.Edges)
	fmt.Printf("Dijkstra result: Dijkstra(a, h, graph)=%d\n", Dijkstra(a, h, graph))
}
