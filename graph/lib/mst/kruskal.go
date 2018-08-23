package mst

import (
	"algorithms/base"
	"algorithms/graph/core"
)

func Kruskal(g *core.G) *core.G {

	// initial mst
	g1 := core.NewG()
	for v := range g.Vs {
		g1.AddV(v)
	}
	// prepare data structure
	eq := base.NewPriorityQueue()
	for v1, m := range g.Es {
		for v2, w := range m {
			eq.Enqueue(core.NewE(v1, v2, w))
		}
	}
	set := base.NewUFS()
	for v := range g.Vs {
		set.Add(int(v))
	}
	for !eq.IsEmpty() {
		// Kruskal
		e := eq.Dequeue().(*core.E)
		if !set.IsUnion(int(e.From), int(e.To)) {
			set.Union(int(e.From), int(e.To))
			g1.AddNonDirectiveE(e.From, e.To, e.W)
		}
	}
	return g1
}
