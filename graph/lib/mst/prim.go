package mst

import (
	"algorithms/graph/core"
)

func SlowPrim(graph *core.Graph) *core.Graph {
	// initial
	g := new(*core.Graph)
	g.Vertices = make(map[*core.Vertex]bool)
	g.Edges = make(map[*core.Vertex]map[*core.Vertex]int)
	for v := range graph.Vertices {
		g.Vertices[v] = true
		break
	}
	// compare
	for !equalVertices(g.Vertices, graph.Vertices) {
		v1, v2, w := scan(g, graph)
		if v1 == nil {
			panic("unknown error")
		}
		g.Vertices[v2] = true
		g.Edges[v1][v2] = w
		g.Edges[v2][v1] = w
	}
	return g
}

func scan(g1, g2 *core.Graph) (*core.Vertex, *core.Vertex, int) {
	for v1 := range g1.Vertices {
		rankLimit := len(g2.Edges[v1])
		for rank := 1; rank <= rankLimit; rank++ {
			v2, w := minW(g2.Edges[v1], rank)
			if !g1.Vertices[v2] {
				return v1, v2, w
			}
		}
	}
	return nil, nil, -1
}

// rank must be larger than 0
func minW(w map[*core.Vertex]int, rank int) (*core.Vertex, int) {
	abort := make(map[*core.Vertex]bool)
	var minV *core.Vertex
	minW := -1
	for ; rank > 0; rank-- {
		// initial
		for v, we := range w {
			if !abort[v] {
				minV = v
				minW = we
				break
			}
		}
		// compare
		for v, we := range w {
			if !abort[v] && we < minW {
				minV = v
				minW = we
			}
		}
		abort[minV] = true
	}
	return minV, minW
}

func equalVertices(vc1, vc2 map[*core.Vertex]bool) bool {
	if len(vc1) != len(vc2) {
		return false
	}
	for v1 := range vc1 {
		if _, existed := vc2[v1]; !existed {
			return false
		}
	}
	return true
}
