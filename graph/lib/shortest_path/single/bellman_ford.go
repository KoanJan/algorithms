package single

import "algorithms/graph/core"

func BellmanFord(g *core.G, s core.V) (map[core.V]int, bool) {
	// init distance set
	d := make(map[core.V]int)
	for v := range g.Vs {
		d[v] = core.Unreachable
	}
	d[s] = 0
	// init iteration times
	// a route may contains at most lne(g.Vs) node
	n := len(g.Vs)
	for i := 0; i < n; i++ {
		// iterate with 'd(v) ?= d(u) + w(u, v)'
		for u, e := range g.Es {
			for v, w := range e {
				if d[v] > d[u]+w {
					d[v] = d[u] + w
				}
			}
		}
	}
	// iterate to check minus w exists if 'd(v) > d(u) + w(u, v)'
	for u, e := range g.Es {
		for v, w := range e {
			if d[v] > d[u]+w {
				return nil, false
			}
		}
	}
	return d, true
}
