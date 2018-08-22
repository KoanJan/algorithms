package lib

import "algorithms/graph/core"

func NaiveDFS(g *core.G, dst core.V) bool {
	// random source node
	for from := range g.Vs {
		return findFrom(g, from, dst)
	}
}

func findFrom(g *core.G, from, dst core.V) bool {
	toList := g.Es[from]
	for to := range toList {
		if to == dst && findFrom(g, to, dst) {
			return true
		}
	}
	return false
}
