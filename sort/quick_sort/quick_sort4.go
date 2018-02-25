package quick_sort

import (
	"fmt"

	"algorithms/sort/basic_sort"
)

func QuickSort4(a []int, simpleSortLevel int) []int {
	quickSort4(a, 0, len(a)-1, initializePivot3, simpleSortLevel)
	return a
}

func quickSort4(a []int, l, r int, initializePivot PivotInitializer, simpleSortLevel int) {
	size := r - l
	if size > 0 {
		fmt.Printf("l=%d r=%d\n", l, r)
		if size <= simpleSortLevel {
			basic_sort.InsertSort(a, l, r)
			fmt.Printf("size is %d and use simple sort algorithm\n", size)
			return
		}
		p := partition2(a, l, r, initializePivot)
		fmt.Printf("get p=%d\n", p)
		quickSort4(a, l, p-1, initializePivot, simpleSortLevel)
		quickSort4(a, p+1, r, initializePivot, simpleSortLevel)
	}
}
