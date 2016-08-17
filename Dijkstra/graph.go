package main

// Point
type Point struct {
	Name string
}

func NewPoint(name string) *Point {
	return &Point{Name: name}
}

// Graph
type Graph struct {
	Points []*Point
	Edges  map[string]map[string]int // key is Point name and value is weights of edge
}

func (g *Graph) CountPoints() int {
	return len(g.Points)
}

func NewGraph(points []*Point, edges map[string]map[string]int) *Graph {
	g := Graph{Points: points, Edges: edges}
	return &g
}
