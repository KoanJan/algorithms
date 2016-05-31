package main

import (
	"fmt"
	"strings"
)

func main() {
	//test
	pat := "ABCAC"
	raw := "ABCABBAC."
	yes := pattern(pat, raw)
	fmt.Printf("\"%s\" is in \"%s\": %v\n", pat, raw, yes)
}

func pattern(pat, raw string) bool {
	rawArray := strings.Split(raw, "")
	patArray := strings.Split(pat, "")
	if !(len(patArray) < len(rawArray)) {
		return false
	}
	var (
		k1        int
		k2        int
		patBarrel []int = barrel(pat)
		patLen    int   = len(patArray)
		rawLen    int   = len(rawArray)
	)
	fmt.Printf("patBarrel : %v\n", patBarrel)
	fmt.Printf("patArray : %v\n", patArray)
	if len(patBarrel) != len(patArray) {
		fmt.Printf("patBarrel len: %d, patArray len: %d\n", len(patBarrel), len(patArray))
		return false
	}
	for {
	loop:
		if !(k1 < rawLen && k2 < patLen) {
			break
		}
		fmt.Printf("k1=%d k2=%d\n", k1, k2)
		fmt.Printf("v1: %s, v2: %s \n", rawArray[k1], patArray[k2])
		if patArray[k2] != rawArray[k1] {
			k1++
			if patBarrel[k2] > 1 {
				k1 -= patBarrel[k2]
			}
			k2 = 0
			goto loop
		} else if k2+1 == len(patArray) {
			return true
		} else {
			k1++
			k2++
		}
	}
	return false
}

func barrel(val string) []int {
	bar := []int{}
	tmp := map[string]int{}
	valArray := strings.Split(val, "")
	for _, v := range valArray {
		if _, ok := tmp[v]; ok {
			tmp[v] = tmp[v] + 1
		} else {
			tmp[v] = 1
		}
		bar = append(bar, tmp[v])
	}
	return bar
}
