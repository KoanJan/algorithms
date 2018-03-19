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

type Edge struct {
	From, To *Vertex
	W        int
}

func (edge Edge) Priority() int {
	return -edge.W
}

func NewEdge(from, to *Vertex, w int) *Edge {
	edge := &Edge{from, to, w}
	return edge
}

func CopyVertices(src map[*Vertex]bool) map[*Vertex]bool {
	dst := make(map[*Vertex]bool)
	for v := range src {
		dst[v] = true
	}
	return dst
}
