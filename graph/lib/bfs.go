package lib

import (
	"algorithms/base"
	"algorithms/graph/core"
)

const (
	ColorWhite core.Color = iota
	ColorGray
	ColorBlack
)

func NaiveBFS(graph *core.Graph, value int) bool {
	var v *core.Vertex
	for v = range graph.Vertices {
		break
	}
	queue := base.NewFastQueue()
	enqueue(queue, v)
	for !queue.IsEmpty() {
		node := queue.Dequeue().(*core.Vertex)
		if node.Value == value {
			return true
		}
		for to := range graph.Edges[node] {
			if to.Color == ColorWhite {
				enqueue(queue, to)
			}
		}
		node.Color = ColorBlack
	}
	return false
}

func enqueue(queue base.Queue, v *core.Vertex) {
	v.Color = ColorGray
	queue.Enqueue(v)
}
