package main

import (
	"math/rand"
	"time"
)

func RandIntn(max int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(max)
}

func InSlice(i int, ia []int) bool {
	for _, _i := range ia {
		if i == _i {
			return true
		}
	}
	return false
}

func AppendUniq(ia []int, i int) []int {
	if InSlice(i, ia) {
		return ia
	}
	return append(ia, i)
}
