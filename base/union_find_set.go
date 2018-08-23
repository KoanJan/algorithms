package base

// UFS means union find set
type UFS map[int]int

func (u UFS) Add(v int, p ...int) {
	if len(p) > 0 {
		root := u.Find(p[0])
		u[root] -= 1
		u[v] = p[0]
	} else {
		u[v] = -1
	}
}

func (u UFS) Find(v int) int {
	if _, exist := u[v]; !exist {
		u[v] = -1
		return v
	}
	for u[v] > 0 {
		v = u[v]
	}
	return v
}

func (u UFS) Union(v ...int) {
	n := len(v)
	if n < 2 {
		return
	}
	p := u.Find(v[0])
	for i := 1; i < n; i++ {
		pi := u.Find(v[i])
		if p != pi {
			u[p] += u[pi]
			u[pi] = p
		}
	}
}

func (u UFS) IsUnion(v ...int) bool {
	n := len(v)
	if n == 0 {
		return false
	}
	p := u.Find(v[0])
	for i := 1; i < n; i++ {
		if p != u.Find(v[i]) {
			return false
		}
	}
	return true
}

func NewUFS() UFS {
	return make(UFS)
}
