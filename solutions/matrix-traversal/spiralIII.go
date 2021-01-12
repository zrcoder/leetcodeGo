package matrix_traversal

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	dirs := [][]int{{0,1}, {1,0}, {0, -1}, {-1, 0}} // 右、下、左、上
	steps := 0 // 步长，向同一个方向走 steps 步
	result := make([][]int, R*C)

	r, c := r0, c0
	result[0] = []int{r, c}
	i := 1
	for i < len(result) {
		for j, d := range dirs {
			if j == 0 || j == 2 { // 根据走的规律，在合适时机增加步长
				steps++
			}
			for k := 0; k < steps; k++ {
				r, c = r+d[0], c+d[1]
				if r >= 0 && r < R && c >= 0 && c < C {
					result[i] = []int{r, c}
					i++
				}
			}
		}
	}
	return result
}