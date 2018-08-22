package multi

import "algorithms/graph/core"

func FloydWarshall(g *core.G) map[core.V]map[core.V]int {
	// initial
	ds := make(map[core.V]map[core.V]int)
	for v := range g.Vs {
		ds[v] = make(map[core.V]int)
		for v2 := range g.Vs {
			ds[v][v2] = core.Unreachable
		}
		for v2, w := range g.Es[v] {
			ds[v][v2] = w
		}
		ds[v][v] = 0
	}
	// iterate
	for v1 := range g.Vs {
		for v2 := range g.Vs {
			for k := range g.Vs {
				if ds[v1][v2] > ds[v1][k]+ds[k][v2] {
					ds[v1][v2] = ds[v1][k] + ds[k][v2]
				}
			}
		}
	}
	return ds
}
