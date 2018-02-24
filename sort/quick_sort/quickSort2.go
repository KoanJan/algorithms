package quick_sort

import "fmt"

func QuickSort2(a []int) []int {
	quickSort2(a, 0, len(a)-1, initializePivot2)
	return a
}

func initializePivot2(a []int, l, r int) int {
	return l
}

func quickSort2(a []int, l, r int, initializePivot PivotInitializer) {
	if l < r {
		fmt.Printf("l=%d r=%d\n", l, r)
		p := partition2(a, l, r, initializePivot)
		fmt.Printf("get p=%d\n", p)
		quickSort2(a, l, p-1, initializePivot)
		quickSort2(a, p+1, r, initializePivot)
	}
}

func partition2(a []int, start, end int, initializePivot PivotInitializer) int {
	fmt.Printf("before partition, a: %v\n", a)
	pivot := a[initializePivot(a, start, end)]
	fmt.Printf("pivot=%d\n", pivot)
	l, r := start+1, end
	for l < r {
		for r > l {
			if a[r] > pivot {
				r--
			} else {
				break
			}
		}
		for l < r {
			if a[l] <= pivot {
				l++
			} else {
				break
			}
		}
		if l != r {
			a[l], a[r] = a[r], a[l]
			l++
			r--
		}
	}
	if a[r] <= pivot {
		a[start], a[r] = a[r], a[start]
		fmt.Printf("after partition, a: %v\n", a)
		return r
	} else {
		fmt.Printf("after partition, a: %v\n", a)
		return l - 1
	}
}
