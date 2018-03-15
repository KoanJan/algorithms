package lib

import "algorithms/graph/core"

func NaiveDFS(graph *core.Graph, value int) bool {
	var v *core.Vertex
	for v = range graph.Vertices {
		if len(graph.Edges[v]) > 0 {
			break
		}
		return false
	}
	return dfs(graph, v, value)
}

func dfs(graph *core.Graph, vertex *core.Vertex, value int) bool {
	if vertex.Color == ColorBlack {
		return false
	}
	vertex.Color = ColorBlack
	if vertex.Value == value {
		return true
	}
	for to := range graph.Edges[vertex] {
		if dfs(graph, to, value) {
			return true
		}
	}
	return false
}
