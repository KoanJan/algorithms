package multi

import "algorithms/graph/core"

func FloydWarshall(graph *core.Graph) map[*core.Vertex]map[*core.Vertex]int {
	// initial
	disMap := make(map[*core.Vertex]map[*core.Vertex]int)
	for v := range graph.Vertices {
		disMap[v] = make(map[*core.Vertex]int)
		disMap[v][v] = 0
	}
	// iterate
	for v1 := range graph.Vertices {
		for v2 := range graph.Vertices {
			for k := range graph.Vertices {
				d1, exist1 := disMap[v1][k]
				if exist1 {
					d2, exist2 := disMap[k][v2]
					if exist2 {
						d, exist := disMap[v1][v2]
						if !exist || d1+d2 < d {
							disMap[v1][v2] = d1 + d2
						}
					}
				}
			}
		}
	}
	return disMap
}
