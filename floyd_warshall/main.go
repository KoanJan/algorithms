package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// test
	picture := newPicture(5)
	fmt.Println("Before: \n")
	printPicture(picture)
	floydWarshall(picture)
	fmt.Println("After: \n")
	printPicture(picture)
	floydWarshall(picture)
	fmt.Println("Again: \n")
	printPicture(picture)
}

func printPicture(picture [][]int) {
	picStr := ""
	tab := ""
	for _, col := range picture {
		for _, ele := range col {
			eleStr := fmt.Sprintf("%d", ele)
			if len(eleStr) == 2 {
				tab = " "
			} else {
				tab = "  "
			}
			picStr = fmt.Sprintf("%s%s%s", picStr, tab, eleStr)
		}
		picStr += "\n"
	}
	fmt.Println(picStr)
}

func floydWarshall(picture [][]int) {
	length := len(picture)
	if length == 0 {
		return
	}
	// invalid picture
	for _, col := range picture {
		if len(col) != length {
			return
		}
	}
	// Floyd-Warshall
	for k := 0; k < length; k++ {
		for i := 0; i < length; i++ {
			for j := 0; j < length; j++ {
				if picture[i][j] > picture[i][k]+picture[k][j] {
					fmt.Printf("[%d][%d] = %d, [%d][%d] = %d, [%d][%d] = %d\n", i, j, picture[i][j], i, k, picture[i][k], k, j, picture[k][j])
					picture[i][j] = picture[i][k] + picture[k][j]
				}
			}
		}
	}
	fmt.Println("\n")
	return
}

func newPicture(length int) [][]int {
	picture := [][]int{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		col := []int{}
		for j := 0; j < length; j++ {
			col = append(col, r.Intn(100))
		}
		picture = append(picture, col)
	}
	return picture
}
