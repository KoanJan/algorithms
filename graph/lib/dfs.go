package lib

import "algorithms/graph/core"

func NaiveDFS(graph *core.Graph, value int) bool {
	var v *core.Vertex
	for v = range graph.Vertices {
		break
	}
	return dfs(v, value)
}

func dfs(vertex *core.Vertex, value int) bool {
	if vertex.Color == ColorBlack {
		return false
	}
	vertex.Color = ColorBlack
	if vertex.Value == value {
		return true
	}
	for _, edge := range vertex.Edges {
		if dfs(edge.To, value) {
			return true
		}
	}
	return false
}
