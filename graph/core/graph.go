package core

type Graph struct {
	Vertices map[*Vertex]bool
}

type Color int

type Vertex struct {
	Value int
	Color Color
	Edges []*Edge
}

type Edge struct {
	From, To *Vertex
	Weight   int
}
