package main

import (
	"fmt"
)

func main() {
	src := []int{12, 86, 81, 646, 47, 91, 75, 32}

	// merge sort
	fmt.Printf("MergeSort: \n\t%v\n", MergeSort(src))
	// quick sort
	fmt.Printf("QuickSort: \n\t%v\n", QuickSort(src))
	// counting sort
	fmt.Printf("CountingSort: \n\t%v\n", CountingSort(src))
	// radix sort
	fmt.Printf("RadixSort: \n\t%v\n", RadixSort(src))
}
