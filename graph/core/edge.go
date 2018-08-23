package core

// in order to avoid from overflowing
const Unreachable = 1<<32 - 1

// E means edge
type E struct {
	From, To V
	W        int
}

func (edge *E) Priority() int {
	return edge.W
}

func NewEdge(from, to V, w int) *E {
	edge := &E{from, to, w}
	return edge
}
