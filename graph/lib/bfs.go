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
	v := graph.Vertex
	queue := base.NewFastQueue()
	enqueue(queue, v)
	for !queue.IsEmpty() {
		node := queue.Dequeue().(*core.Vertex)
		if node.Value == value {
			return true
		}
		for _, edge := range node.Edges {
			next := edge.To
			if next.Color == ColorWhite {
				enqueue(queue, next)
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
