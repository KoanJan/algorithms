package single

import (
	"algorithms/base"
	"algorithms/graph/core"
)

func Dijkstra(g *core.G, v core.V) map[core.V]int {

	disMap := make(map[core.V]int)
	disMap[v] = 0
	disQueue := base.NewPriorityQueue()
	disQueue.Enqueue(core.NewEdge(v, v, 0))
	for !disQueue.IsEmpty() {
		e := disQueue.Dequeue().(*core.E)
		for to, w := range g.Es[e.To] {
			// record the distance from the head vertex
			disQueue.Enqueue(core.NewEdge(v, to, w+e.W))
		}
		// compare and save the smallest distance
		d, exited := disMap[e.To]
		if (!exited) || e.W < d {
			disMap[e.To] = e.W
		}
	}
	return disMap
}
