package main

// MergeSort
func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	return merge(MergeSort(a[0:len(a)/2]), MergeSort(a[len(a)/2:]))
}

func merge(a []int, b []int) []int {
	var (
		res []int = []int{}
		ca  int   = 0
		cb  int   = 0
	)
	for {
		// index out
		if ca == len(a) || cb == len(b) {
			break
		}
		if a[ca] < b[cb] {
			res = append(res, a[ca])
			ca++
		} else {
			res = append(res, b[cb])
			cb++
		}
	}
	for ca < len(a) {
		res = append(res, a[ca])
		ca++
	}
	for cb < len(b) {
		res = append(res, b[cb])
		cb++
	}
	return res
}
