package single

import (
	"testing"

	"algorithms/graph/core"
)

func testDijkstra() map[core.V]int {
	// init g
	g := core.NewG()
	for i := 1; i <= 10; i++ {
		g.AddV(core.V(i))
	}
	// 1
	g.AddNonDirectiveE(1, 3, 10)
	g.AddNonDirectiveE(1, 5, 20)
	g.AddNonDirectiveE(1, 7, 30)
	g.AddNonDirectiveE(1, 9, 20)
	// 2
	g.AddNonDirectiveE(2, 1, 2)
	g.AddNonDirectiveE(2, 3, 1)
	g.AddNonDirectiveE(2, 4, 7)
	g.AddNonDirectiveE(2, 7, 3)
	// 3
	g.AddNonDirectiveE(3, 2, 1)
	g.AddNonDirectiveE(3, 4, 2)
	g.AddNonDirectiveE(3, 6, 3)
	g.AddNonDirectiveE(3, 8, 1)
	// 4
	g.AddNonDirectiveE(4, 7, 1)
	g.AddNonDirectiveE(4, 8, 2)
	g.AddNonDirectiveE(4, 9, 3)
	g.AddNonDirectiveE(4, 10, 1)
	// 5
	g.AddNonDirectiveE(5, 1, 4)
	g.AddNonDirectiveE(5, 6, 2)
	g.AddNonDirectiveE(5, 3, 3)
	g.AddNonDirectiveE(5, 4, 5)
	// 6
	g.AddNonDirectiveE(6, 1, 3)
	g.AddNonDirectiveE(6, 3, 4)
	g.AddNonDirectiveE(6, 4, 3)
	g.AddNonDirectiveE(6, 5, 3)
	// 7
	g.AddNonDirectiveE(7, 2, 4)
	g.AddNonDirectiveE(7, 4, 2)
	g.AddNonDirectiveE(7, 5, 3)
	g.AddNonDirectiveE(7, 6, 2)
	// 8
	g.AddNonDirectiveE(8, 5, 2)
	g.AddNonDirectiveE(8, 2, 5)
	g.AddNonDirectiveE(8, 3, 3)
	g.AddNonDirectiveE(8, 10, 4)
	// 9
	g.AddNonDirectiveE(9, 1, 1)
	g.AddNonDirectiveE(9, 5, 2)
	g.AddNonDirectiveE(9, 7, 3)
	g.AddNonDirectiveE(9, 9, 1)
	// 10
	g.AddNonDirectiveE(10, 2, 1)
	g.AddNonDirectiveE(10, 7, 3)
	g.AddNonDirectiveE(10, 1, 3)
	g.AddNonDirectiveE(10, 8, 2)
	// test
	return Dijkstra(g, 1)
}

func TestDijkstra(t *testing.T) {
	t.Log(testDijkstra())
}
