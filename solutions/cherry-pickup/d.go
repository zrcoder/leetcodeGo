/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package cherry_pickup

import "math"

/*
741. 摘樱桃 https://leetcode-cn.com/problems/cherry-pickup
一个N x N的网格(grid) 代表了一块樱桃地，每个格子由以下三种数字的一种来表示：

0 表示这个格子是空的，所以你可以穿过它。
1 表示这个格子里装着一个樱桃，你可以摘到樱桃然后穿过它。
-1 表示这个格子里有荆棘，挡着你的路。
你的任务是在遵守下列规则的情况下，尽可能的摘到最多樱桃：

从位置 (0, 0) 出发，最后到达 (N-1, N-1) ，只能向下或向右走，并且只能穿越有效的格子（即只可以穿过值为0或者1的格子）；
当到达 (N-1, N-1) 后，你要继续走，直到返回到 (0, 0) ，只能向上或向左走，并且只能穿越有效的格子；
当你经过一个格子且这个格子包含一个樱桃时，你将摘到樱桃并且这个格子会变成空的（值变为0）；
如果在 (0, 0) 和 (N-1, N-1) 之间不存在一条可经过的路径，则没有任何一个樱桃能被摘到。

示例 1:
输入: grid =
[[0, 1, -1],
 [1, 0, -1],
 [1, 1,  1]]
输出: 5
解释：
玩家从（0,0）点出发，经过了向下走，向下走，向右走，向右走，到达了点(2, 2)。
在这趟单程中，总共摘到了4颗樱桃，矩阵变成了[[0,1,-1],[0,0,-1],[0,0,0]]。
接着，这名玩家向左走，向上走，向上走，向左走，返回了起始点，又摘到了1颗樱桃。
在旅程中，总共摘到了5颗樱桃，这是可以摘到的最大值了。

说明:
grid 是一个 N * N 的二维数组，N的取值范围是1 <= N <= 50。
每一个 grid[i][j] 都是集合 {-1, 0, 1}中的一个数。
可以保证起点 grid[0][0] 和终点 grid[N-1][N-1] 的值都不会是 -1。
*/

/*
问题简化：如果只是从起点到终点摘一遍呢？这将是一个典型的动态规划

定义dp(r,c)表示从起点走到(r,c)摘到的最大樱桃数； 则dp[r][c] = max(dp[r-1][c], dp[r][c-1]) + grid[r][c]
注意边界情况即(r,c)处本身为荆棘或其上边和左边一格均无法到达的情况
或者dp多申请一行一列，0行0列都是0，不参与结果计算，只是方便少判断边界
*/
func cherryPickupOnce(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return -1
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {
			cherryNum := grid[r-1][c-1]
			if cherryNum == -1 {
				dp[r][c] = -1
				continue
			}
			dp[r][c] = cherryNum
			topLeftMax := max(dp[r][c-1], dp[r-1][c])
			if topLeftMax > 0 {
				dp[r][c] += topLeftMax
			}
		}
	}
	return dp[n][n]
}

/*
可以考虑执行cherryPickup两次，且第一次把摘掉的樱桃格子里的值置为0
但这样会导致过多的格子被置为0；
可以考虑记录第一次摘过樱桃的路径，但显然又会漏掉最优解，
如下边的情况，为了明显，将两个特别的樱桃用★表示了：
				{1,1,1,1,0,0,0},
				{0,0,0,1,0,0,0},
				{0,0,0,1,0,0,★},
				{★,0,0,1,0,0,0},
				{0,0,0,1,0,0,0},
				{0,0,0,1,0,0,0},
				{0,0,0,1,1,1,1},
可见最优解是所有值为1的格子都能被摘掉，共15
但如果采用上面的办法，结果会是14，左下或右上的★会被漏掉一个
*/

/* ---------------------------------------------------
问题转化：两个人同时从左上角走到右下角
*/

/*
[解法1：动态规划，自顶向下]

假设有2个人，在走了t步后；分别居于(r1, c1), (r2, c2)位置
因r1+c1=r2+c2=t；所以r2=r1+c1-c2，这意味着r1，c1，c2唯一地决定了2个走了t步的人，以这个条件来做动态规划：
定义dp[r1][c1][c2]为从(r1, c1), (r2, c2)开始，走到终点（n-1，n-1）所能摘到的最多樱桃数量；其中r2=r1+c1-c2
如果(r1, c1), (r2, c2)处不是荆棘；那么dp[r1][c1][c2]的值这样计算：
先得到(r1, c1), (r2, c2)两处的樱桃总数（如果位置重复则只算一次）； 再加上
max(
dp(r1+1, c1, c2, dp, grid),	// a, b都向下
dp(r1, c1+1, c2, dp, grid),	// a右b下
dp(r1+1, c1, c2+1, dp, grid),	// a下b右
dp(r1, c1+1, c2+1, dp, grid))	// a，b都向右
*/
func cherryPickup(grid [][]int) int {
	n := len(grid)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
			for k := 0; k < n; k++ {
				dp[i][j][k] = math.MinInt32
			}
		}
	}
	return max(0, pickup(0, 0, 0, dp, grid))
}

func pickup(r1, c1, c2 int, dp [][][]int, grid [][]int) int {
	r2 := r1 + c1 - c2
	n := len(grid)
	if n == r1 || n == c1 || n == r2 || n == c2 ||
		grid[r1][c1] == -1 || grid[r2][c2] == -1 {
		return math.MinInt32
	}
	if r1 == n-1 && c1 == n-1 {
		return grid[r1][c1]
	}
	if dp[r1][c1][c2] != math.MinInt32 {
		return dp[r1][c1][c2]
	}
	val := grid[r1][c1]
	if c1 != c2 {
		val += grid[r2][c2]
	}
	val += max(
		pickup(r1+1, c1, c2, dp, grid),
		pickup(r1, c1+1, c2, dp, grid),
		pickup(r1+1, c1, c2+1, dp, grid),
		pickup(r1, c1+1, c2+1, dp, grid))
	dp[r1][c1][c2] = val
	return val
}

/*
[解法2：动态规划，自顶向下]
与解法1类似，只是这里定义的dp[r1][c1][c2]为从起点走到(r1, c1), (r2, c2)所能摘到的最多樱桃数量；其中r2=r1+c1-c2

一开始调用的时候r1,c1,c2的值传n-1
*/
func pickup1(r1, c1, c2 int, dp [][][]int, grid [][]int) int {
	r2 := r1 + c1 - c2
	if r1 < 0 || c1 < 0 || r2 < 0 || c2 < 0 ||
		grid[r1][c1] == -1 || grid[r2][c2] == -1 {
		return math.MinInt32
	}

	if r1 == 0 && c1 == 0 && c2 == 0 {
		return grid[r1][c1]
	}

	if dp[r1][c1][c2] != math.MinInt32 {
		return dp[r1][c1][c2]
	}

	val := grid[r1][c1]
	if c1 != c2 {
		val += grid[r2][c2]
	}
	val += max(
		pickup1(r1-1, c1, c2, dp, grid),
		pickup1(r1, c1-1, c2, dp, grid),
		pickup1(r1-1, c1, c2-1, dp, grid),
		pickup1(r1, c1-1, c2-1, dp, grid))
	dp[r1][c1][c2] = val
	return val
}

/*
[解法3：动态规划，自底向上]

定义dp[c1][c2]为第t步，从起点走到(r1,c1)和从起点走到(r2,c2)能摘到的最多樱桃数；其中r1=t-c1，r2=t-c2
一个人从左上角走到右下角共需n-1 + n-1 即2n-2步
*/
func cherryPickup1(grid [][]int) int {
	n := len(grid)
	dp := genDp(n)
	dp[0][0] = grid[0][0]
	for t := 1; t <= 2*n-2; t++ {
		dp2 := genDp(n)
		from, end := max(0, t-(n-1)), min(n-1, t)
		for i := from; i <= end; i++ {
			if grid[i][t-i] == -1 {
				continue
			}
			for j := from; j <= end; j++ {
				if grid[j][t-j] == -1 {
					continue
				}
				val := grid[i][t-i]
				if i != j {
					val += grid[j][t-j]
				}
				for pi := i - 1; pi <= i; pi++ {
					for pj := j - 1; pj <= j; pj++ {
						if pi >= 0 && pj >= 0 {
							dp2[i][j] = max(dp2[i][j], dp[pi][pj]+val)
						}
					}
				}
			}
		}
		dp = dp2
	}
	return max(0, dp[n-1][n-1])
}

func genDp(n int) [][]int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = math.MinInt32
		}
	}
	return dp
}

func max(nums ...int) int {
	r := math.MinInt32
	for _, v := range nums {
		if v > r {
			r = v
		}
	}
	return r
}

func min(nums ...int) int {
	r := math.MaxInt64
	for _, v := range nums {
		if v < r {
			r = v
		}
	}
	return r
}
