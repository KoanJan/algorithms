package main

import (
	"fmt"
)

// BM str1 in str2, return [] if not existed
func BM(str1, str2 string) []int {
	var (
		bs1   = []byte(str1)
		bs2   = []byte(str2)
		right = map[byte]int{}
		n     = len(bs1)
		m     = len(bs2)
		res   = []int{}
	)

	// set right
	for i2, b2 := range bs2 {
		right[b2] = i2
	}

	fmt.Printf("params:\nbs1: %s\nbs2: %s\nright: %v\nn: %d\nm: %d\n\n", str1, str2, right, n, m)

	for i := 0; i < n-m; {

		for j := m - 1; j >= 0; j-- {

			// DEBUG
			fmt.Printf("%s[%d] = %c while %s[%d] = %c\n", str2, j, bs2[j], str1, i+j, bs1[i+j])

			if bs1[i+j] == bs2[j] {
				// fmt.Printf("%v[%d](%v) == %v[%d](%v)\n", bs1, i+j+1, bs1[i+j+1], bs2, j, bs2[j])
				if j == 0 {
					res = append(res, i)
					i++
					break
				}
			} else {
				var (
					skip int = -1
					ni   int = i
				)
				for k := i; k < i+j+1; k++ {
					fmt.Printf("bs1[%d] = %c   bs2[%d] = %c\n", k, bs1[k], k-i, bs2[k-i])
					if idx, existed := right[bs1[k]]; existed {
						if skip < idx {
							skip = idx
							ni = k
						}
					}
				}
				if skip < 0 {
					skip = 0
				}
				fmt.Printf("ni: %d, j: %d, skip: %d\n", ni, j, skip)
				i = ni + j - skip + 1
				break
			}
		}
	}

	return res
}
