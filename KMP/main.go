package main

import (
	"fmt"
	"strings"
)

func main() {
	//test
	yes := pattern("pat", "A pattern.")
	fmt.Println("\"pat\" is in \"A pattern\": ", yes)
}

// no dictionary
func pattern(pat, raw string) bool {
	rawArray := strings.Split(raw, "")
	patArray := strings.Split(pat, "")
	if !(len(patArray) < len(rawArray)) {
		return false
	}
	var k1 int
	for k2, v2 := range patArray {
	loop:
		fmt.Printf("v1: %s, v2: %s \n", rawArray[k1], v2)
		if v2 != rawArray[k1] {
			k1++
			goto loop
		} else if k2+1 == len(patArray) {
			return true
		} else {
			k1++
		}
	}
	return false
}
