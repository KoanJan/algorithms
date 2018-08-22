package multi

import "algorithms/graph/core"

func FloydWarshall(g *core.G) map[*core.V]map[*core.V]int {
	// initial
	disMap := make(map[*core.V]map[*core.V]int)
	for v := range g.Vs {
		disMap[v] = make(map[*core.V]int)
		disMap[v][v] = 0
	}
	// iterate
	for v1 := range g.Vs {
		for v2 := range g.Vs {
			for k := range g.Vs {
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
