package main

import (
	"fmt"

	"algorithms/sort/basic_sort"
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

	src3 := []int{12, 86, 81, 646, 47, 91, 75, 32}
	// bubble sort
	basic_sort.BubbleSort(src3)
	fmt.Printf("BubbleSort: \n\t%v\n", src3)

	src4 := []int{12, 86, 81, 646, 47, 91, 75, 32}
	// select sort
	basic_sort.SelectSort(src4)
	fmt.Printf("SelectSort: \n\t%v\n", src4)

}
