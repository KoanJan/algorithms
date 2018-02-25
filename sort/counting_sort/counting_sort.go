package counting_sort

// CountingSort
func CountingSort(a []int) []int {
	b := make([]int, len(a)+1)
	countingSort(a, b, max(a))
	return b[1:]
}

func countingSort(a, b []int, k int) {
	c := make([]int, k+1)
	for j := 0; j < len(a); j++ {
		c[a[j]] += 1
	}
	for i := 1; i <= k; i++ {
		c[i] += c[i-1]
	}
	for j := len(a) - 1; j >= 0; j-- {
		b[c[a[j]]] = a[j]
		c[a[j]] -= 1
	}
}

func max(a []int) int {
	var r int
	for _, v := range a {
		if v > r {
			r = v
		}
	}
	return r
}
