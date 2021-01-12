package paths

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	// 初始化dp第0行
	for c := 1; c < n && obstacleGrid[0][c] == 0; c++ {
		dp[0][c] = 1
	}
	// 初始化dp第0列
	for r := 1; r < m && obstacleGrid[r][0] == 0; r++ {
		dp[r][0] = 1
	}
	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			if obstacleGrid[r][c] == 0 {
				dp[r][c] = dp[r-1][c] + dp[r][c-1]
			}
		}
	}
	return dp[m-1][n-1]
}
