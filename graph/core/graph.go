package core

import (
	"bytes"
	"fmt"
)

// Graph
type G struct {
	Vs map[V]bool
	Es map[V]map[V]int
}

func (g *G) AddV(v V) {
	// init v
	g.Vs[v] = true
	// init e
	if len(g.Es[v]) == 0 {
		g.Es[v] = make(map[V]int, len(g.Vs)-1)
		for other := range g.Vs {
			if other == v {
				g.Es[v][v] = 0
				continue
			}
			g.Es[v][other] = Unreachable
			g.Es[other][v] = Unreachable
		}
	}
}

func (g *G) AddE(v1, v2 V, w int) {
	g.Es[v1][v2] = w
}

func (g *G) AddNonDirectiveE(v1, v2 V, w int) {
	g.Es[v1][v2] = w
	g.Es[v2][v1] = w
}

func (g *G) String() string {
	buf := bytes.NewBufferString("\nv: ")
	for v := range g.Vs {
		buf.WriteString(fmt.Sprintf("%d ", v))
	}
	buf.WriteString("\ne: \n")
	for v1, e := range g.Es {
		for v2, w := range e {
			buf.WriteString(fmt.Sprintf("%d - %d: %d\n", v1, v2, w))
		}
	}
	return buf.String()
}

func NewG() *G {
	g := G{make(map[V]bool), make(map[V]map[V]int)}
	return &g
}
