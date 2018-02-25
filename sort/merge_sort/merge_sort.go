package merge_sort

func MergeSort(a []int) []int {
	m := len(a) / 2
	return merge(a[0:m], a[m:])
}

func merge(a, b []int) []int {
	lna, lnb := len(a), len(b)
	if lna > 1 {
		a = merge(a[0:lna/2], a[lna/2:])
	}
	if lnb > 1 {
		b = merge(b[0:lnb/2], b[lnb/2:])
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
