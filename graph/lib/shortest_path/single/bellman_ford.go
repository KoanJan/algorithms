package single

import "algorithms/graph/core"

func BellmanFord(graph *core.Graph, vertex *core.Vertex) (map[*core.Vertex]int, bool) {
	disMap := make(map[*core.Vertex]int)
	n := len(graph.Vertices)
	// n-1 times
	relaxed := true
	for i := 1; relaxed && i < n; i++ {
		// iterate all the edges
		relaxed = false
		for s, edge := range graph.Edges {
			for t, w := range edge {
				if _, existed := disMap[s]; existed {
					_, newVertex := disMap[t]
					if newVertex || disMap[t] > disMap[s]+w {
						disMap[t] = disMap[s] + w
						relaxed = true
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
