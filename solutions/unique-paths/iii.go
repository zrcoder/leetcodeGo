package paths

func uniquePathsIII(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	startR, startC, steps := prepair(grid)
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	result := 0
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n ||
			visited[r][c] || grid[r][c] == -1 {
			return
		}
		if grid[r][c] == 2 {
			if steps == 0 {
				result++
			}
			return
		}
		steps--
		visited[r][c] = true
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c+1)
		dfs(r, c-1)
		steps++
		visited[r][c] = false
	}
	dfs(startR, startC)
	return result
}

func prepair(grid [][]int) (int, int, int) {
	r, c, steps := 0, 0, 1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				steps++
			}
			if grid[i][j] == 1 {
				r, c = i, j
			}
		}
	}
	return r, c, steps
}
