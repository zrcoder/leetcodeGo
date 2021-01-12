package paths

/*
576. 出界的路径数 https://leetcode-cn.com/problems/out-of-boundary-paths/
给定一个 m × n 的网格和一个球。球的起始坐标为 (i,j) ，你可以将球移到相邻的单元格内，或者往上、下、左、右四个方向上移动使球穿过网格边界。
但是，你最多可以移动 N 次。找出可以将球移出边界的路径数量。
答案可能非常大，返回 结果 mod 109 + 7 的值。

示例 1：
输入: m = 2, n = 2, N = 2, i = 0, j = 0
输出: 6
解释:

示例 2：
输入: m = 1, n = 3, N = 3, i = 0, j = 1
输出: 12
解释:

说明:

球一旦出界，就不能再被移动回网格内。
网格的长度和高度在 [1,50] 的范围内。
N 在 [0,50] 的范围内。
*/
/*
从dfs到dp
*/
/*
1. 朴素dfs实现，76 / 94 个通过测试用例；第77个超时
*/
func findPaths0(m int, n int, N int, i int, j int) int {
	if i >= N && (m-i) > N && j >= N && (n-j) > N {
		return 0
	}
	const max = 1000000007
	var dfs func(r, c, rest int) int
	dfs = func(r, c, rest int) int {
		if rest < 0 {
			return 0
		}
		if r == -1 || r == m || c == -1 || c == n {
			return 1
		}
		rest--
		return (dfs(r+1, c, rest)%max + dfs(r-1, c, rest)%max + dfs(r, c-1, rest)%max + dfs(r, c+1, rest)%max) % max
	}
	return dfs(i, j, N)
}

/*
2. 加备忘录，减少递归
*/
func findPaths1(m int, n int, N int, i int, j int) int {
	if i >= N && (m-i) > N && j >= N && (n-j) > N {
		return 0
	}
	const max = 1000000007
	const init = -1
	memo := make3d(m, n, N+1, init)
	var dfs func(r, c, rest int) int
	dfs = func(r, c, rest int) int {
		if rest < 0 {
			return 0
		}
		if r == -1 || r == m || c == -1 || c == n {
			return 1
		}
		if memo[r][c][rest] != init {
			return memo[r][c][rest]
		}
		rest--
		down := dfs(r+1, c, rest) % max
		up := dfs(r-1, c, rest) % max
		left := dfs(r, c-1, rest) % max
		right := dfs(r, c+1, rest) % max
		rest++
		memo[r][c][rest] = (down + up + left + right) % max
		return memo[r][c][rest]
	}
	return dfs(i, j, N)
}

/*
3. 自底向上到动态规划
上面到解法是自顶向下递归，不难改为自底向上动态规划；

参考：
三维dp的动态规划：https://leetcode-cn.com/problems/out-of-boundary-paths/solution/zhuang-tai-ji-du-shi-zhuang-tai-ji-by-christmas_wa/
二维dp的动态规划：https://leetcode-cn.com/problems/out-of-boundary-paths/solution/javade-dfsyu-dong-tai-gui-hua-by-zackqf/

代码略
*/

func make3d(m, n, k, fill int) [][][]int {
	r := make([][][]int, m)
	for i := range r {
		r[i] = make([][]int, n)
		for j := range r[i] {
			r[i][j] = make([]int, k)
			for k := range r[i][j] {
				r[i][j][k] = fill
			}
		}
	}
	return r
}
