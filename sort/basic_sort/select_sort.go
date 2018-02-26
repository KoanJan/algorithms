package basic_sort

// SelectSort
func SelectSort(a []int) {
	ln := len(a)
	if ln < 2 {
		return
	}
	for i := 0; i < ln-1; i++ {
		minIdx := i
		for j := ln - 1; j > i; j-- {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		a[i], a[minIdx] = a[minIdx], a[i]
	}
}
