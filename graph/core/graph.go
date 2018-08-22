package core

// Graph
type G struct {
	Vs map[*V]bool
	Es map[*V]map[*V]int
}
