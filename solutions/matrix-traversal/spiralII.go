package matrix_traversal

func generateMatrix(n int) [][]int {
	if n <= 0 {
		return nil
	}
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	levels := (n + 1) / 2
	num := 1
	for i := 0; i < levels; i++ {
		for c := i; c < n-i; c++ {
			res[i][c] = num
			num++
		}
		for r := i + 1; r < n-i; r++ {
			res[r][n-1-i] = num
			num++
		}
		for c := n - 2 - i; i != n-1-i && c >= i; c-- {
			res[n-1-i][c] = num
			num++
		}
		for r := n - 2 - i; i != n-1-i && r > i; r-- {
			res[r][i] = num
			num++
		}
	}
	return res
}
