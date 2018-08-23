package core

// in order to avoid from overflowing
const Unreachable = 1<<32 - 1

// E means edge
type E struct {
	From, To V
	W        int
}

func (edge *E) Priority() int {
	return edge.W
}

func NewE(from, to V, w int) *E {
	edge := &E{from, to, w}
	return edge
}

// Es is a edge set
type Es map[V]map[V]int

func (es Es) ExistE(v1, v2 V) bool {
	if _, exist := es[v1]; !exist {
		return false
	}
	_, exist := es[v1][v2]
	return exist
}

func (es Es) Get(v1, v2 V) int {
	if e, exist1 := es[v1]; exist1 {
		if w, exist2 := e[v2]; exist2 {
			return w
		}
	}
	return Unreachable
}

func (es Es) Add(v1, v2 V, w int) {
	if _, exist1 := es[v1]; !exist1 {
		es[v1] = make(map[V]int)
	}
	es[v1][v2] = w
}

func (es Es) AddNonDirective(v1, v2 V, w int) {
	es.Add(v1, v2, w)
	es.Add(v2, v1, w)
}

func NewEs() Es {
	return make(Es)
}
