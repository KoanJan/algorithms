package mst

import (
	"algorithms/graph/core"
	"fmt"
)

func Prim(g *core.G) *core.G {
	// init mst
	g1 := core.NewG()
	for v := range g.Vs {
		g1.Vs[v] = true
		break
	}
	var (
		minV1 = core.NonexistV
		minV2 = core.NonexistV
		minW  = core.Unreachable
		todo  = len(g.Vs) - 1
	)
	// scan until all vertex added into g1
	for i := 0; i < todo; i++ {

		fmt.Println(g1)
		// reset minW
		minW = core.Unreachable
		// scan
		for v2 := range g.Vs {
			// existed vertex
			if g1.Vs[v2] {
				continue
			}
			// iterate
			oldMinW := minW
			for v1 := range g1.Vs {
				if g.Es[v1][v2] < minW {
					minW = g.Es[v1][v2]
					minV1 = v1
				}
				if g.Es[v2][v1] < minW {
					minW = g.Es[v2][v1]
					minV1 = v1
				}
			}
			if minW < oldMinW {
				minV2 = v2
			}
		}
		// add new vertex into g1
		g1.Vs[minV2] = true
		if len(g1.Es[minV1]) == 0 {
			g1.Es[minV1] = map[core.V]int{minV2: minW}
		} else {
			g1.Es[minV1][minV2] = minW
		}
		if len(g1.Es[minV2]) == 0 {
			g1.Es[minV2] = map[core.V]int{minV1: minW}
		} else {
			g1.Es[minV2][minV1] = minW
		}
	}

	return g1
}
