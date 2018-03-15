package lib

import "algorithms/graph/core"

func BellmanFord(graph *core.Graph, vertex *core.Vertex) (map[*core.Vertex]int, bool) {
	disMap := make(map[*core.Vertex]int)
	n := len(graph.Vertices)
	// n-1 times
	for i := 1; i < n; i++ {
		// iterate all the edges
		for s, edge := range graph.Edges {
			for t, w := range edge {
				if _, existed := disMap[s]; existed {
					_, newVertex := disMap[t]
					if newVertex || disMap[t] > disMap[s]+w {
						disMap[t] = disMap[s] + w
					}
				}
			}
		}
	}
	// check if there's a ring
	for s, edge := range graph.Edges {
		for t, w := range edge {
			if disMap[t] > disMap[s]+w {
				return nil, false
			}
		}
	}
	return disMap, true
}
