package lib

import "algorithms/graph/core"

func DFS(graph *core.Graph, value int) bool {
	return dfs(graph.Vertex, value)
}

func dfs(vertex *core.Vertex, value int) bool {
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
