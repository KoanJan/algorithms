package main

// LCS Longest Common Subsequence
func LCS(str1, str2 string) ([][]int, [][]string) {

	x := []byte(str1)
	y := []byte(str2)

	// no common subsequence
	if len(x) == 0 || len(y) == 0 {
		return nil, nil
	}

	// init b
	b := make([][]string, len(x)+1)
	for i := 0; i < len(x)+1; i++ {
		b[i] = make([]string, len(y)+1)
	}

	// init c
	c := make([][]int, len(x)+1)
	for i := 0; i < len(x)+1; i++ {
		c[i] = make([]int, len(y)+1)
	}
	for i := 1; i < len(x)+1; i++ {
		c[i][0] = 0
	}
	for i := 1; i < len(y)+1; i++ {
		c[0][i] = 0
	}

	// loop
	for i := 1; i < len(x)+1; i++ {
		for j := 1; j < len(y)+1; j++ {
			if x[i-1] == y[j-1] {
				c[i][j] = c[i-1][j-1] + 1
				b[i][j] = "LU"
			} else if c[i-1][j] >= c[i][j-1] {
				c[i][j] = c[i-1][j]
				b[i][j] = "U"
			} else {
				c[i][j] = c[i][j-1]
				b[i][j] = "L"
			}
		}
	}

	// return
	return c, b
}
