/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package gifts_max_value

import "math"

/*
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

示例 1:
输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 12
解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
典型的动态规划，时空复杂度都是O(m*n)
*/
func maxValue(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m+1) // 多加一行一列， 第一行和第一列元素全为0，方便后边代码，不用判断太多边界情况
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			dp[r+1][c+1] = max(dp[r][c+1], dp[r+1][c]) + grid[r][c]
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
