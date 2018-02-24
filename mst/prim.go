package main

func Prim(g Graph) Graph {
	tree := NewGraph(g.Size())
	if g.Size() == 0 {
		return tree
	}

	// avoid taking same edge
	check := make([][]bool, g.Size())
	for i := 0; i < len(check); i++ {
		c := make([]bool, g.Size())
		check[i] = c
	}

	addedNodes := []int{}

	// shortest edge
	min := struct{ X, Y, W int }{Unconnected, Unconnected, Unconnected}
	for len(addedNodes) < g.Size() {

		// find the shortest edge
		for i := 0; i < g.Size(); i++ {
			for j := 0; j < g.Size(); j++ {
				if i == j || check[i][j] || check[j][i] {
					continue
				}
				if len(addedNodes) != 0 {
					if (InSlice(i, addedNodes) && InSlice(j, addedNodes)) || (!InSlice(i, addedNodes) && !InSlice(j, addedNodes)) {
						continue
					}
				}
				if g[i][j] < min.W {
					min.X = i
					min.Y = j
					min.W = g[i][j]
				}
			}
		}

		// break if cannot catch any available edges
		if min.X == Unconnected || min.Y == Unconnected {
			break
		}
		// record the edge that has been caught
		check[min.X][min.Y] = true
		check[min.Y][min.X] = true

		// add it to tree
		tree[min.X][min.Y] = min.W
		// keep no circle
		addedNodes = AppendUniq(addedNodes, min.X)
		addedNodes = AppendUniq(addedNodes, min.Y)

		// reset min
		min.X = Unconnected
		min.Y = Unconnected
		min.W = Unconnected
	}

	return tree
}
