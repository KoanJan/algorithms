package lib

import (
	"algorithms/base"
	"algorithms/graph/core"
)

type distance core.Edge

func (d distance) Priority() int {
	return -d.Weight
}

func Dijkstra(vertex *core.Vertex) map[*core.Vertex]int {

	disMap := make(map[*core.Vertex]int)
	disMap[vertex] = 0
	disQueue := base.NewPriorityQueue()
	disQueue.Enqueue(distance{vertex, vertex, 0})
	for !disQueue.IsEmpty() {
		e := disQueue.Dequeue().(*core.Edge)
		for _, edge := range e.To.Edges {
			// record the distance from the head vertex
			disQueue.Enqueue(distance{vertex, edge.To, e.Weight + edge.Weight})
		}
		// compare and save the smallest distance
		d, exited := disMap[e.To]
		if (!exited) || e.Weight < d {
			disMap[e.To] = e.Weight
		}
	}
}
