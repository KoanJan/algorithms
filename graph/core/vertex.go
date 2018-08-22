package core

type Color int

// V means vertex
type V struct {
	Value int
	Color Color
}

func CopyVertices(src map[*V]bool) map[*V]bool {
	dst := make(map[*V]bool)
	for v := range src {
		dst[v] = true
	}
	return dst
}
