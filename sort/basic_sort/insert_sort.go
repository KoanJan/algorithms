package basic_sort

func InsertSortAll(a []int) []int {
	end := len(a) - 1
	if end > 0 {
		InsertSort(a, 0, len(a)-1)
	}
	return a
}

// InsertSort will sort a list of numbers from index 'start' to 'end' where the smaller number on left size
func InsertSort(a []int, start, end int) {
	for i := start; i <= end; i++ {
		for j := i; j > start; j-- {
			x, y := a[i], a[j]
			if x < y {
				a[i], a[j] = y, x
			}
		}
	}
}
