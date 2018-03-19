package single

import (
	"algorithms/base"
	"algorithms/graph/core"
)

type distance struct {
	from, to *core.Vertex
	weight   int
}

func (d distance) Priority() int {
	return -d.weight
}

func Dijkstra(graph *core.Graph, vertex *core.Vertex) map[*core.Vertex]int {

	disMap := make(map[*core.Vertex]int)
	disMap[vertex] = 0
	disQueue := base.NewPriorityQueue()
	disQueue.Enqueue(distance{vertex, vertex, 0})
	for !disQueue.IsEmpty() {
		e := disQueue.Dequeue().(distance)
		for to, weight := range graph.Edges[e.to] {
			// record the distance from the head vertex
			disQueue.Enqueue(distance{vertex, to, weight + e.weight})
		}
		// compare and save the smallest distance
		d, exited := disMap[e.to]
		if (!exited) || e.weight < d {
			disMap[e.to] = e.weight
		}
	}
}
