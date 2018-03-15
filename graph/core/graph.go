package core

type Graph struct {
	Vertices map[*Vertex]bool
	Edges    map[*Vertex]map[*Vertex]int
}

type Color int

type Vertex struct {
	Value int
	Color Color
}
