package merge_sort

import "algorithms/sort/basic_sort"

func MergeSort2(a []int, threshold int) []int {
	m := len(a) / 2
	return merge(a[0:m], a[m:], threshold, basic_sort.InsertSort)
}
