package main

// Q queue
type Q struct {
	data []int
}

// Push push
func (q *Q) Push(p int) {
	q.data = append(q.data, p)
}

// Pop pop
func (q *Q) Pop() int {
	if q.IsEmpty() {
		return -1
	}
	p := q.data[0]
	q.data = q.data[1:]
	return p
}

// IsEmpty return if queue is empty
func (q *Q) IsEmpty() bool {
	return len(q.data) == 0
}

// BFS
func BFS(graph [][]bool, s int) ([]int, []int) {

	// out of index
	if s >= len(graph) {
		panic("out of index")
	}

	// init
	var (
		q       = Q{}
		d       = []int{}
		colors  = []int{}
		parents = []int{}
	)
	for i := 0; i < len(graph); i++ {
		if graph[s][i] {
			d = append(d, 1)
		} else {
			d = append(d, Gigantic)
		}
		colors = append(colors, PointColorWhite)
		parents = append(parents, NilParent)
	}

	// search
	colors[s] = PointColorGray
	d[s] = 0
	parents[s] = NilParent

	q.Push(s)

	for !q.IsEmpty() {
		u := q.Pop()
		for v := range graph[u] {
			if colors[v] == PointColorWhite && graph[u][v] {
				colors[v] = PointColorGray
				d[v] = d[u] + 1
				parents[v] = u
				q.Push(v)
			}
		}
		colors[u] = PointColorBlack
	}

	return d, parents

}
