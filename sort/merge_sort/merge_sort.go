package merge_sort

import "algorithms/sort/basic_sort"

func MergeSort(a []int) []int {
	m := len(a) / 2
	return merge(a[0:m], a[m:], 1, nerverSort)
}

func nerverSort(a []int, l, r int) {
}

func merge(a, b []int, threshold int, sorter basic_sort.Sorter) []int {
	lna, lnb := len(a), len(b)
	if lna > threshold {
		a = merge(a[0:lna/2], a[lna/2:], threshold, sorter)
	} else {
		sorter(a, 0, lna-1)
	}
	if lnb > threshold {
		b = merge(b[0:lnb/2], b[lnb/2:], threshold, sorter)
	} else {
		sorter(b, 0, lnb-1)
	}
	result := make([]int, 0, lna+lnb)
	i, j := 0, 0
	for i < lna && j < lnb {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	for i < lna {
		result = append(result, a[i])
		i++
	}
	for j < lnb {
		result = append(result, b[j])
		j++
	}
	return result
}
