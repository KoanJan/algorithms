package mst

import (
	"algorithms/base"
	"algorithms/graph/core"
)

func Kruskal(graph *core.Graph) *core.Graph {

	// initial mst
	mst := new(core.Graph)
	mst.Vertices = core.CopyVertices(graph.Vertices)
	mst.Edges = make(map[*core.Vertex]map[*core.Vertex]int)
	for v := range mst.Vertices {
		mst.Edges[v] = make(map[*core.Vertex]int)
	}
	// prepare data structure
	edgeQueue := base.NewPriorityQueue()
	for v1, m := range graph.Edges {
		for v2, w := range m {
			edgeQueue.Enqueue(core.NewEdge(v1, v2, w))
		}
	}
	vertices := make([]interface{}, 0, len(graph.Vertices))
	for v := range graph.Vertices {
		vertices = append(vertices, v)
	}
	set := base.NewUnionFindSet(vertices...)
	for !edgeQueue.IsEmpty() {
		// Kruskal
		edge := edgeQueue.Dequeue().(*core.Edge)
		if set.Find(edge.From) != set.Find(edge.To) {
			set.SetParent(edge.To, edge.From)
			mst.Edges[edge.From][edge.To] = edge.W
			mst.Edges[edge.To][edge.From] = edge.W
		}
	}
	return mst

}
