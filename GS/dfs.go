package main

var time = 0

// DFS
func DFS(graph [][]bool, s int) [][]int {

	// out of index
	if s >= len(graph) {
		panic("out of index")
	}

	var (
		d = make([]int, len(graph))
		f = make([]int, len(graph))

		colors  = make([]int, len(graph))
		parents = make([][]int, len(graph))
	)

	// init
	for u := range graph[s] {
		if graph[s][u] {
			colors[u] = PointColorWhite
		}
		parent := []int{}
		for i := 0; i < len(graph); i++ {
			parent = append(parent, NilParent)
		}
		parents[u] = parent
	}

	for u := range graph[s] {
		if colors[u] == PointColorWhite {

			// DFS-visit
			dfsVisit(graph, parents, colors, d, f, u, u)
		}
	}

	return parents
}

// dfsVisit
func dfsVisit(graph [][]bool, parents [][]int, colors, d, f []int, u, sp int) {
	colors[u] = PointColorGray
	time += 1
	d[u] = time
	for v := range graph[u] {
		if colors[v] == PointColorWhite {
			parents[sp][v] = u
			dfsVisit(graph, parents, colors, d, f, v, sp)
		}
	}
	colors[u] = PointColorBlack
	time += 1
	f[u] = time
}
