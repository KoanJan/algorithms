package heap_sort

func HeapSort(a []int) {
	n := len(a)
	if n < 2 {
		return
	}
	// make a heap
	for i := 1; i < n; i++ {
		for j := i; j > 0; j >>= 1 {
			if a[j] > a[j>>1] {
				a[j], a[j>>1] = a[j>>1], a[j]
			} else {
				break
			}
		}
	}
	// pop
	for i := n - 1; i > 0; i-- {
		a[i], a[0] = a[0], a[i]
		for j := 0; ; {
			// max sub node
			x := j<<1 + 1
			y := x + 1
			if y < i && a[x] < a[y] && a[j] < a[y] {
				a[j], a[y] = a[y], a[j]
				j = y
			} else if x < i && a[j] < a[x] {
				a[j], a[x] = a[x], a[j]
				j = x
			} else {
				break
			}
		}
	}
}
