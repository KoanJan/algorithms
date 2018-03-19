package single

import (
	"algorithms/base"
	"algorithms/graph/core"
)

func Dijkstra(graph *core.Graph, vertex *core.Vertex) map[*core.Vertex]int {

	disMap := make(map[*core.Vertex]int)
	disMap[vertex] = 0
	disQueue := base.NewPriorityQueue()
	disQueue.Enqueue(core.NewEdge(vertex, vertex, 0))
	for !disQueue.IsEmpty() {
		e := disQueue.Dequeue().(*core.Edge)
		for to, weight := range graph.Edges[e.To] {
			// record the distance from the head vertex
			disQueue.Enqueue(core.NewEdge(vertex, to, weight+e.W))
		}
		// compare and save the smallest distance
		d, exited := disMap[e.To]
		if (!exited) || e.W < d {
			disMap[e.To] = e.W
		}
	}
}
