package multi

import (
	"testing"

	"algorithms/graph/core"
)

func testFloydWarshall() map[core.V]map[core.V]int {
	// init g
	g := core.NewG()
	for i := 1; i <= 10; i++ {
		g.Vs[core.V(i)] = true
	}
	var s, v1, v2, v3, v4 core.V = 1, 3, 5, 7, 9
	w1, w2, w3, w4 := 10, 20, 30, 20
	g.Es[s] = map[core.V]int{v1: w1, v2: w2, v3: w3, v4: w4}
	g.Es[2] = map[core.V]int{1: 2, 3: 1, 4: 7, 7: 3}
	g.Es[3] = map[core.V]int{2: 1, 4: 2, 6: 3, 8: 1}
	g.Es[4] = map[core.V]int{7: 1, 8: 2, 9: 3, 10: 1}
	g.Es[5] = map[core.V]int{1: 4, 6: 2, 3: 3, 4: 5}
	g.Es[6] = map[core.V]int{1: 3, 3: 4, 4: 3, 5: 3}
	g.Es[7] = map[core.V]int{2: 4, 4: 2, 5: 3, 6: 2}
	g.Es[8] = map[core.V]int{5: 2, 2: 5, 3: 3, 10: 4}
	g.Es[9] = map[core.V]int{1: 1, 5: 2, 7: 3, 9: 1}
	g.Es[10] = map[core.V]int{2: 1, 7: 3, 1: 3, 8: 2}
	// test
	return FloydWarshall(g)
}

func TestFloydWarshall(t *testing.T) {
	t.Log(testFloydWarshall())
}

func BenchmarkFloydWarshall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testFloydWarshall()
	}
}
