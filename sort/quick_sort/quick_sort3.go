package quick_sort

func QuickSort3(a []int) []int {
	quickSort2(a, 0, len(a)-1, initializePivot3)
	return a
}

func initializePivot3(a []int, l, r int) int {
	if l-r < 2 {
		return l
	}
	m := (l + r) / 2
	if a[l] > a[m] {
		a[l], a[m] = a[m], a[l]
	}
	if a[r] > a[m] {
		a[r], a[m] = a[m], a[r]
	}
	if a[l] < a[r] {
		a[l], a[r] = a[r], a[l]
	}
	return l
}
