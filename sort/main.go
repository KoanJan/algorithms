package main

import (
	"fmt"

	"algorithms/sort/merge_sort"
	"algorithms/sort/quick_sort"
)

func main() {
	src := []int{12, 86, 81, 646, 47, 91, 75, 32}
	// quick sort
	fmt.Printf("BetterQuickSort: \n\t%v\n", quick_sort.QuickSort4(src, 4))

	src2 := []int{12, 86, 81, 646, 47, 91, 75, 32}
	// merge sort
	fmt.Printf("MergeSort: \n\t%v\n", merge_sort.MergeSort(src2))

}
