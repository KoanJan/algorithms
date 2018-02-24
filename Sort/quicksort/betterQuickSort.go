package quicksort

import "fmt"

func BetterQuickSort(a []int) []int {
	betterQuickSort(a, 0, len(a)-1)
	return a
}

func betterQuickSort(a []int, l, r int) {
	if l < r {
		fmt.Printf("l=%d r=%d\n", l, r)
		p := partition2(a, l, r)
		fmt.Printf("get p=%d\n", p)
		betterQuickSort(a, l, p-1)
		betterQuickSort(a, p+1, r)
	}
}

func partition2(a []int, start, end int) int {
	fmt.Printf("before partition, a: %v\n", a)
	pivot := a[start]
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
