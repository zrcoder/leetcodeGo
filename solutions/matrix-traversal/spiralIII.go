package matrix_traversal

var row, column int

func spiralMatrixIII(R int, C int, r int, c int) [][]int {
	steps := 0
	row, column = R, C
	res := make([][]int, 0, R*C)
	res = append(res, []int{r, c})
	for len(res) < cap(res) {
		originR, originC := r, c
		// right
		steps++
		for k := 0; c < C && k < steps && len(res) < cap(res); k++ {
			c++
			res = checkToAppend(res, r, c)
		}
		c = originC + steps
		// down
		for k := 0; r < R && k < steps && len(res) < cap(res); k++ {
			r++
			res = checkToAppend(res, r, c)
		}
		r = originR + steps
		// left
		steps++
		for k := 0; c >= 0 && k < steps && len(res) < cap(res); k++ {
			c--
			res = checkToAppend(res, r, c)
		}
		c = originC - 1
		// up
		for k := 0; r >= 0 && k < steps && len(res) < cap(res); k++ {
			r--
			res = checkToAppend(res, r, c)
		}
		r = originR - 1
	}
	return res
}

func checkToAppend(res [][]int, r, c int) [][]int {
	if isValid(r, c) {
		res = append(res, []int{r, c})
	}
	return res
}

func isValid(r, c int) bool {
	return r >= 0 && r < row && c >= 0 && c < column
}

// 代码可以优化：
var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func spiralMatrixIIILoop(R int, C int, r int, c int) [][]int {
	// right, down, left, up
	steps := 0
	row, column = R, C
	res := make([][]int, 0, R*C)
	res = append(res, []int{r, c})
	for len(res) < cap(res) {
		for i, d := range dirs {
			if i == 0 || i == 2 {
				steps++
			}
			for k := 0; k < steps; k++ {
				r, c = r+d[0], c+d[1]
				res = checkToAppend(res, r, c)
			}
		}
	}
	return res
}
