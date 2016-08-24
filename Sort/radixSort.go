package main

// RadixSort
func RadixSort(a []int, radix ...int) []int {
	r := 10
	if len(radix) > 0 {
		if radix[0] > 1 {
			r = radix[0]
		}
	}
	radixSort(a, r)
	return a
}

func radixSort(a []int, radix int) {
	r := radix
	radix = 1
	n := len(a)
	m := maxBit(a, r)
	tmp := make([]int, n)
	count := make([]int, r)
	var i, j, k int
	//
	for i = 1; i <= m; i++ {

		// clean
		for j = 0; j < r; j++ {
			count[j] = 0
		}
		tmp = make([]int, n)
		// records of each bucket
		for j = 0; j < n; j++ {
			k = (a[j] / radix) % r
			count[k] += 1
		}
		// make sort
		for j = 1; j < r; j++ {
			count[j] += count[j-1]
		}
		// collection
		for j = n - 1; j >= 0; j-- {
			k = (a[j] / radix) % r
			tmp[count[k]-1] = a[j]
			count[k] -= 1
		}
		// copy
		for j = 0; j < n; j++ {
			a[j] = tmp[j]
		}
		radix *= r
	}
}

// count digits of number
func digit(n, radix int) int {
	d := 1
	for n >= radix {
		n /= radix
		d += 1
	}
	return d
}

func maxBit(a []int, radix int) int {
	m := 1
	for _, v := range a {
		_d := digit(v, radix)
		if _d > m {
			m = _d
		}
	}
	return m
}
