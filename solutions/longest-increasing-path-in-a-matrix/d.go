/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package longest_increasing_path_in_a_matrix

/*
329.矩阵中的最长递增路径 https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix
给定一个整数矩阵，找出最长递增路径的长度。

对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外（即不允许环绕）。

示例 1:

输入: nums =
[
[9,9,4],
[6,6,8],
[2,1,1]
]
输出: 4
解释: 最长递增路径为 [1, 2, 6, 9]。
示例 2:

输入: nums =
[
[3,4,5],
[3,2,6],
[2,2,1]
]
输出: 4
解释: 最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。
*/
/*
朴素dfs

时间复杂度 ：O(2^(m+n))。对每个有效递增路径均进行搜索。在最坏情况下，会有2^(m+n)次调用。例如：
1 2 3 . . . n
2 3 . . .   n+1
3 . . .     n+2
.           .
.           .
.           .
m m+1 . . . n+m-1
空间复杂度 ： O(mn)。 对于每次深度优先搜索，系统栈需要 O(h)空间，其中 h 为递归的最深深度。最坏情况下， O(h) = O(mn)。
*/
func longestIncreasingPath1(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		result := 1 // 一个元素自身可作为一个递增序列
		for _, d := range dirs {
			x, y := r+d[0], c+d[1]
			if x >= 0 && x < m && y >= 0 && y < n && matrix[x][y] > matrix[r][c] {
				result = max(result, 1+dfs(x, y))
			}
		}
		return result
	}
	result := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			result = max(result, dfs(r, c))
		}
	}
	return result
}

/*
可以用一个备忘录存储dfs函数里已经计算的结果，减少重复计算
时间复杂度降为O((mn)^2)；空间复杂度依然是O(mn)
*/
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n) // 可以初始化全部元素为1，但也可在dfs过程中初始化为1
	}
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if memo[r][c] > 0 {
			return memo[r][c]
		}
		memo[r][c] = 1
		for _, d := range dirs {
			x, y := r+d[0], c+d[1]
			if x >= 0 && x < m && y >= 0 && y < n && matrix[x][y] > matrix[r][c] {
				memo[r][c] = max(memo[r][c], 1+dfs(x, y))
			}
		}
		return memo[r][c]
	}
	result := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			result = max(result, dfs(r, c))
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
