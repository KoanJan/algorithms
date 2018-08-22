package mst

import (
	"algorithms/base"
	"algorithms/graph/core"
)

func Kruskal(g *core.G) *core.G {

	// initial mst
	mst := new(core.G)
	mst.Vs = core.CopyVertices(g.Vs)
	mst.Es = make(map[*core.V]map[*core.V]int)
	for v := range mst.Vs {
		mst.Es[v] = make(map[*core.V]int)
	}
	// prepare data structure
	eq := base.NewPriorityQueue()
	for v1, m := range g.Es {
		for v2, w := range m {
			eq.Enqueue(core.NewEdge(v1, v2, w))
		}
	}
	vs := make([]interface{}, 0, len(g.Vs))
	for v := range g.Vs {
		vs = append(vs, v)
	}
	set := base.NewUnionFindSet(vs...)
	for !eq.IsEmpty() {
		// Kruskal
		e := eq.Dequeue().(*core.E)
		if set.Find(e.From) != set.Find(e.To) {
			set.SetParent(e.To, e.From)
			mst.Es[e.From][e.To] = e.W
			mst.Es[e.To][e.From] = e.W
		}
	}
	return mst

}
