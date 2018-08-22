package lib

import (
	"algorithms/base"
	"algorithms/graph/core"
)

func NaiveBFS(g *core.G, dst core.V) bool {
	// map vertex and color
	vc := make(map[core.V]Color, len(g.Vs))
	for v := range g.Vs {
		vc[v] = White
	}
	// init queue
	q := base.NewFastQueue()
	for v := range g.Vs {
		q.Enqueue(v)
		vc[v] = Gray
		break
	}
	// iterate
	for !q.IsEmpty() {
		from := q.Dequeue().(core.V)
		if from == dst {
			return true
		}
		for to := range g.Es[from] {
			if vc[to] == White {
				q.Enqueue(to)
				vc[to] = Gray
			}
		}
		vc[from] = Black
	}
	return false
}
