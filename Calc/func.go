package main

// Add
func Add(a, b int) int {
	if b == 0 {
		return a
	}
	return Add(a^b, (a&b)<<1)
}

// Minus
func Minus(a, b int) int {
	var i int = 1
	for i > 0 && ((b & i) == 0) {
		i <<= 1
	}
	for i > 0 {
		i <<= 1
		b ^= i
	}
	return Add(a, b)
}

// Multi
func Multi(a, b int) int {
	var (
		ans int = 0
		i   int = 1
	)
	for i > 0 {
		if b&i > 0 {
			ans = Add(ans, a)
		}
		i <<= 1
		a <<= 1
	}
	return ans
}

// Div
func Div(a, b int) int {
	if a < b {
		return 0
	}
	// simplify
	for b&1 == 0 {
		a >>= 1
		b >>= 1
	}
	var (
		ans int = 0
		i   int = 1
		db  int = 1
	)
	for a >= b {
		b <<= 1
		db = Add(db, 1)
	}
	b >>= 1
	db = Minus(db, 1)
	for a >= i {
		i <<= 1
	}
	i >>= 1
	for db > 0 {
		if a >= b && a&i > 0 {
			a = Minus(a, b)
			ans = Add(ans, 1)
		}
		for i > a {
			i >>= 1
		}
		b >>= 1
		db = Minus(db, 1)
		ans <<= 1
	}
	ans >>= 1
	return ans
}
