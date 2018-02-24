package main

import (
	"algorithms/sort/basic_sort"
	"algorithms/sort/quick_sort"
	"fmt"
)

func main() {
	src := []int{12, 86, 81, 646, 47, 91, 75, 32}

	// quick sort
	fmt.Printf("BetterQuickSort: \n\t%v\n", quick_sort.QuickSort3(src))
	// insert sort
	fmt.Printf("InsertSort: \n\t%v\n", basic_sort.InsertSortAll(src))
}
