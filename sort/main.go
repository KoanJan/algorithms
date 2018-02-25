package main

import (
	"fmt"

	"algorithms/sort/quick_sort"
	"algorithms/sort/radix_sort"
)

func main() {
	src := []int{12, 86, 81, 646, 47, 91, 75, 32}
	// quick sort
	fmt.Printf("BetterQuickSort: \n\t%v\n", quick_sort.QuickSort4(src, 4))

	src2 := []int{12, 86, 81, 646, 47, 91, 75, 32}
	// merge sort
	fmt.Printf("CountingSort: \n\t%v\n", radix_sort.RadixSort(src2))

}
