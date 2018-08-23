package single

import (
	"algorithms/base"
	"algorithms/graph/core"
)

func Dijkstra(g *core.G, v core.V) map[core.V]int {

	// used edge set
	used := core.NewEs()
	// init d
	d := make(map[core.V]int)
	for eachV := range g.Vs {
		d[eachV] = core.Unreachable
	}
	d[v] = 0
	// init pq
	q := base.NewPriorityQueue()
	q.Enqueue(core.NewE(v, v, 0))
	for !q.IsEmpty() {
		e := q.Dequeue().(*core.E)
		for to, w := range g.Es[e.To] {
			if e.To == to || used.ExistE(e.To, to) {
				continue
			}
			// record the distance from the head vertex
			q.Enqueue(core.NewE(e.To, to, w))
			used.AddNonDirective(e.To, to, w)
		}
		// compare and save the smallest distance
		if d[e.From]+e.W < d[e.To] {
			d[e.To] = d[e.From] + e.W
		}
	}
	return d
}
