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

func NaiveBFS(g *core.G, value int) bool {
	var v *core.V
	for v = range g.Vs {
		break
	}
	queue := base.NewFastQueue()
	enqueue(queue, v)
	for !queue.IsEmpty() {
		node := queue.Dequeue().(*core.V)
		if node.Value == value {
			return true
		}
		for to := range g.Es[node] {
			if to.Color == ColorWhite {
				enqueue(queue, to)
			}
		}
		node.Color = ColorBlack
	}
	return false
}

func enqueue(q base.Queue, v *core.V) {
	v.Color = ColorGray
	q.Enqueue(v)
}
