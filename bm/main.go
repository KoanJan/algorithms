package main

import (
	"fmt"
)

func main() {
	// test
	str1 := "ABCSAKDFFEFKJDDEFKLD"
	str2 := "DDEFK"

	res := BM(str1, str2)

	if len(res) == 0 {
		fmt.Printf("%s is not in %s\n", str2, str1)
	} else {
		fmt.Printf("%s is in %s, and indexes are %v", str2, str1, res)
	}
}
