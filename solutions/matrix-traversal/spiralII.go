package matrix_traversal

func generateMatrix(n int) [][]int {
	if n <= 0 {return nil}
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}
	const right, down, left, up = 0, 1, 2, 3
	r, c, direct := 0, 0, right
	rMin, rMax, cMin, cMax := 0, n-1, 0, n-1 // 上下左右边界
	max := n*n
	for x := 1; x <= max; x++ {
		result[r][c] = x
		switch direct {
		case right:
			if c < cMax {
				c++
			} else {
				direct, rMin, r = down, rMin+1, r+1
			}
		case down:
			if r < rMax {
				r++
			} else {
				direct, cMax, c = left, cMax-1, c-1
			}
		case left:
			if c > cMin {
				c--
			} else {
				direct, rMax, r = up, rMax-1, r-1
			}
		case up:
			if r > rMin {
				r--
			} else {
				direct, cMin, c = right, cMin+1, c+1
			}
		}
	}
	return result
}