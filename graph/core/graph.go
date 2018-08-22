package core

// Graph
type G struct {
	Vs map[V]bool
	Es map[V]map[V]int
}

func NewG() *G {
	g := G{make(map[V]bool), make(map[V]map[V]int)}
	return &g
}
