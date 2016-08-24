package main

// QuickSort
func QuickSort(a []int) []int {
	quickSort(a, 1, len(a)-1)
	return a
}

func quickSort(a []int, p, r int) {
	if p < r {
		q := partition(a, p, r)
		quickSort(a, p, q-1)
		quickSort(a, q+1, r)
	}
}

func partition(a []int, p, r int) int {
	x := a[r]
	i := p - 1
	for j := p; j < r; j++ {
		if a[j] <= x {
			i += 1
			swap(a, i, j)
		}
	}
	swap(a, i+1, r)
	return i + 1
}

func swap(a []int, i, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}
