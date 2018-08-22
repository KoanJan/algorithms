package lib

import "algorithms/graph/core"

func NaiveDFS(g *core.G, value int) bool {
	var v *core.V
	for v = range g.Vs {
		if len(g.Es[v]) > 0 {
			break
		}
		return false
	}
	return dfs(g, v, value)
}

func dfs(g *core.G, v *core.V, value int) bool {
	if v.Color == ColorBlack {
		return false
	}
	v.Color = ColorBlack
	if v.Value == value {
		return true
	}
	for to := range g.Es[v] {
		if dfs(g, to, value) {
			return true
		}
	}
	return false
}
