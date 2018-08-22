package core

const NonexistV V = 1<<63 - 1

// V means vertex, and its value is used as a ID
type V int

func CopyVertices(src map[V]bool) map[V]bool {
	dst := make(map[V]bool)
	for v := range src {
		dst[v] = true
	}
	return dst
}
