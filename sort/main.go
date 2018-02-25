package main

import (
	"fmt"

	"algorithms/sort/counting_sort"
	"algorithms/sort/quick_sort"
)

func main() {
	src := []int{12, 86, 81, 646, 47, 91, -75, 32}
	// quick sort
	fmt.Printf("BetterQuickSort: \n\t%v\n", quick_sort.QuickSort4(src, 4))

	src2 := []int{12, 86, 81, 646, 47, 91, -75, 32}
	// merge sort
	fmt.Printf("CountingSort: \n\t%v\n", counting_sort.CountingSort(src2))

}
