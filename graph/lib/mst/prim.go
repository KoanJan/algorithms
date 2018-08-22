package mst

import (
	"algorithms/graph/core"
)

func SlowPrim(graph *core.G) *core.G {
	// init
	g := new(core.G)
	g.Vs = make(map[core.V]bool)
	g.Es = make(map[core.V]map[core.V]int)
	for v := range graph.Vs {
		g.Vs[v] = true
		break
	}
	// compare
	for !equalVertices(g.Vs, graph.Vs) {
		v1, v2, w := scan(g, graph)
		if v1 == core.NonexistV {
			panic("unknown error")
		}
		g.Vs[v2] = true
		g.Es[v1][v2] = w
		g.Es[v2][v1] = w
	}
	return g
}

func scan(g1, g2 *core.G) (core.V, core.V, int) {
	for v1 := range g1.Vs {
		rankLimit := len(g2.Es[v1])
		for rank := 1; rank <= rankLimit; rank++ {
			v2, w := minW(g2.Es[v1], rank)
			if !g1.Vs[v2] {
				return v1, v2, w
			}
		}
	}
	return core.NonexistV, core.NonexistV, -1
}

// rank must be larger than 0
func minW(w map[core.V]int, rank int) (core.V, int) {
	abort := make(map[core.V]bool)
	var minV core.V
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

func equalVertices(vc1, vc2 map[core.V]bool) bool {
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
