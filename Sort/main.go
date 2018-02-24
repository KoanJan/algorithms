package main

import (
	"algorithms/Sort/quicksort"
	"fmt"
)

func main() {
	src := []int{12, 86, 81, 646, 47, 91, 75, 32}

	// quick sort
	fmt.Printf("BetterQuickSort: \n\t%v\n", quicksort.QuickSort3(src))
}
