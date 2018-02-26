package basic_sort

// BubbleSort
func BubbleSort(a []int) {
	ln := len(a)
	if ln < 2 {
		return
	}
	for i := 0; i < ln-1; i++ {
		for j := ln - 1; j > i; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}
