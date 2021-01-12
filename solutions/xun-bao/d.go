package xun_bao

import "math"

/*
LCP 13. 寻宝
https://leetcode-cn.com/problems/xun-bao

我们得到了一副藏宝图，藏宝图显示，在一个迷宫中存在着未被世人发现的宝藏。

迷宫是一个二维矩阵，用一个字符串数组表示。它标识了唯一的入口（用 'S' 表示），和唯一的宝藏地点（用 'T' 表示）。
但是，宝藏被一些隐蔽的机关保护了起来。在地图上有若干个机关点（用 'M' 表示），只有所有机关均被触发，才可以拿到宝藏。

要保持机关的触发，需要把一个重石放在上面。迷宫中有若干个石堆（用 'O' 表示），每个石堆都有无限个足够触发机关的重石
。但是由于石头太重，我们一次只能搬一个石头到指定地点。

迷宫中同样有一些墙壁（用 '#' 表示），我们不能走入墙壁。剩余的都是可随意通行的点（用 '.' 表示）。
石堆、机关、起点和终点（无论是否能拿到宝藏）也是可以通行的。

我们每步可以选择向上/向下/向左/向右移动一格，并且不能移出迷宫。
搬起石头和放下石头不算步数。那么，从起点开始，我们最少需要多少步才能最后拿到宝藏呢？如果无法拿到宝藏，返回 -1 。

示例 1：

输入： ["S#O", "M..", "M.T"]

输出：16

解释：最优路线为： S->O, cost = 4,
去搬石头 O->第二行的M, cost = 3, M机关触发
第二行的M->O, cost = 3, 我们需要继续回去 O 搬石头。
O->第三行的M, cost = 4, 此时所有机关均触发
第三行的M->T, cost = 2，去T点拿宝藏。 总步数为16。

示例 2：

输入： ["S#O", "M.#", "M.T"]

输出：-1

解释：我们无法搬到石头触发机关

示例 3：

输入： ["S#O", "M.T", "M.."]

输出：17

解释：注意终点也是可以通行的。

限制：

1 <= maze.length<= 100
1 <= maze[i].length <= 100
maze[i].length == maze[j].length
S 和 T 有且只有一个
0 <= M的数量 <= 16
0 <= O的数量 <= 40，题目保证当迷宫中存在 M 时，一定存在至少一个 O 。
*/
var (
	// 迷宫行数、列数
	n, m int
	// 机关 & 石头
	buttons, stones [][]int
	// 起点 & 终点
	sx, sy, tx, ty int
	// 下、上、右、左四个方向
	dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
)

func minimalSteps(maze []string) int {
	preprocesses(maze)

	dist, result := calDist(maze)
	if result != 0 {
		return result
	}

	return binaryDp(dist)
}

func preprocesses(maze []string) {
	n, m = len(maze), len(maze[0])
	buttons, stones = nil, nil
	sx, sy, tx, ty = -1, -1, -1, -1
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			switch maze[r][c] {
			case 'M':
				buttons = append(buttons, []int{r, c})
			case 'O':
				stones = append(stones, []int{r, c})
			case 'S':
				sx, sy = r, c
			case 'T':
				tx, ty = r, c
			}
		}
	}
}

// 计算机关、石头、起点、终点间的距离；
// 发现不可能完成提前返回math.Max；发现没有机关，直接返回从起点到终点的最短距离
func calDist(maze []string) ([][]int, int) {
	nb := len(buttons)
	startDist := bfs(sx, sy, maze)
	// 边界情况：没有机关
	if nb == 0 {
		if startDist[tx][ty] == math.MaxInt32 {
			return nil, -1
		}
		return nil, startDist[tx][ty]
	}
	result := genMatrix(nb, nb+2)
	buttonDist := make([][][]int, nb)
	for i, button := range buttons {
		buttonDist[i] = bfs(button[0], button[1], maze)
		// 机关 -> 终点
		result[i][nb+1] = buttonDist[i][tx][ty]
		if result[i][nb+1] == math.MaxInt32 {
			return nil, -1
		}
		// 机关 -> 石头 -> 起点
		result[i][nb] = calFor(buttonDist[i], startDist)
		if result[i][nb] == math.MaxInt32 {
			return nil, -1
		}
	}
	for i := range buttons {
		// 机关 -> 石头 -> 另一个机关
		for j := i + 1; j < nb; j++ {
			dist := calFor(buttonDist[i], buttonDist[j])
			result[i][j] = dist
			result[j][i] = dist
		}
	}
	return result, 0
}

func calFor(dist1, dist2 [][]int) int {
	result := math.MaxInt32
	for _, stone := range stones {
		r, c := stone[0], stone[1]
		if dist1[r][c] == math.MaxInt32 || dist2[r][c] == math.MaxInt32 {
			continue
		}
		total := dist1[r][c] + dist2[r][c]
		result = min(result, total)
	}
	return result
}

/*
状态压缩动态规划

因为机关的个数不会超过16， 可以用一个一个16位的二进制数 state 表示状态
例如 0000110000010001 表示为机关1、5、11、12被触发，其他为 0 的位置对应的机关没有触发

定义dp(state, i)表示在机关i处，触发状态为state的最小步数

*/
func binaryDp(dist [][]int) int {
	nb := len(buttons)
	total := 1 << nb
	dp := genMatrix(total, nb)
	for i := range buttons {
		// 起点经过某个石头堆到机关i的最小距离
		dp[1<<i][i] = dist[i][nb]
	}
	// 由于更新的状态都比未更新的大，所以直接从小到大遍历即可
	for state := 1; state < total; state++ {
		for i := range buttons {
			dpFor(state, i, dist, dp)
		}
	}
	result := math.MaxInt32
	final := total - 1
	for i := range buttons {
		result = min(result, dp[final][i]+dist[i][nb+1])
	}
	if result == math.MaxInt32 {
		return -1
	}
	return result
}

func dpFor(state, button int, dist, dp [][]int) {
	if state&(1<<button) == 0 { // 机关i未被触发
		return
	}
	for j := range buttons {
		if state&(1<<j) != 0 { // 机关j被触发
			continue
		}
		next := state | (1 << j)
		steps := dp[state][button] + dist[button][j]
		dp[next][j] = min(dp[next][j], steps)
	}
}

// 返回的矩阵记录从 (r,c) 点到达每个点的最短距离
func bfs(r, c int, maze []string) [][]int {
	dist := genMatrix(n, m)
	dist[r][c] = 0
	var queue [][]int
	queue = append(queue, []int{r, c})
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		r, c := pos[0], pos[1]
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if canUpdate(nr, nc, maze, dist) {
				dist[nr][nc] = dist[r][c] + 1
				queue = append(queue, []int{nr, nc})
			}
		}
	}
	return dist
}

func canUpdate(r, c int, maze []string, result [][]int) bool {
	return inBound(r, c) && maze[r][c] != '#' && result[r][c] == math.MaxInt32
}

func inBound(r, c int) bool {
	return r >= 0 && r < n && c >= 0 && c < m
}

func genMatrix(rows, clomns int) [][]int {
	result := make([][]int, rows)
	for r := 0; r < rows; r++ {
		result[r] = make([]int, clomns)
		for c := 0; c < clomns; c++ {
			result[r][c] = math.MaxInt32
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
