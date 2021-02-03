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
	m, n int
	// 机关、石堆
	mPoses, oPoses [][2]int
	// 起点、 终点
	sx, sy, tx, ty int
	// 下、上、右、左四个方向
	dirs = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
)

const inf = math.MaxInt32

func minimalSteps(maze []string) int {
	if !Init(maze) {
		return -1
	}

	distMemo, res := calDist(maze)
	if res != 0 {
		return res
	}

	return dp(distMemo)
}

func Init(maze []string) bool {
	m, n = len(maze), len(maze[0])
	mPoses, oPoses = nil, nil
	sx, sy, tx, ty = -1, -1, -1, -1
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			switch maze[r][c] {
			case 'M':
				mPoses = append(mPoses, [2]int{r, c})
			case 'O':
				oPoses = append(oPoses, [2]int{r, c})
			case 'S':
				sx, sy = r, c
			case 'T':
				tx, ty = r, c
			}
		}
	}
	return sx != -1 && tx != -1
}

// 返回机关经过某个石堆到达另一个机关（起点/终点）的最小距离
// 发现不可能完成提前返回 -1；发现没有机关，直接返回从起点到终点的最短距离
func calDist(maze []string) ([][]int, int) {
	mLen, oLen := len(mPoses), len(oPoses)
	sDist := bfs(sx, sy, maze)
	// 边界情况：没有机关
	if mLen == 0 {
		if sDist[tx][ty] == inf {
			return nil, -1
		}
		return nil, sDist[tx][ty]
	}
	// 边界情况：有机关没石堆
	if oLen == 0 {
		return nil, -1
	}
	// distMemo[i][j] 代表从机关 i 经过某个石堆到达机关 j 的最短距离
	// j 如果是 mLen，代表的是起点，如果是 mLen+1 代表的是终点
	distMemo := genMemo(mLen, mLen+2)
	mDist := make([][][]int, mLen)
	for i, M := range mPoses {
		mDist[i] = bfs(M[0], M[1], maze)
		// 机关 -> 终点
		distMemo[i][mLen+1] = mDist[i][tx][ty]
		if distMemo[i][mLen+1] == inf {
			return nil, -1
		}
		// 机关 -> 石头 -> 起点
		distMemo[i][mLen] = calByStoneDist(mDist[i], sDist)
		if distMemo[i][mLen] == inf {
			return nil, -1
		}
	}
	// 机关 -> 石头 -> 另一个机关
	for i := 0; i < mLen-1; i++ {
		for j := i + 1; j < mLen; j++ {
			dist := calByStoneDist(mDist[i], mDist[j])
			if dist == inf {
				return nil, -1
			}
			distMemo[i][j] = dist
			distMemo[j][i] = dist
		}
	}
	return distMemo, 0
}

// 返回的矩阵记录从(x, y) 点到达每个点的最短距离
func bfs(x, y int, maze []string) [][]int {
	res := genMemo(m, n)
	res[x][y] = 0
	queue := [][]int{{x, y}}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		x, y = p[0], p[1]
		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]
			if check(nx, ny, maze, res) {
				res[nx][ny] = res[x][y] + 1
				queue = append(queue, []int{nx, ny})
			}
		}
	}
	return res
}

func check(r, c int, maze []string, res [][]int) bool {
	return inBound(r, c) && maze[r][c] != '#' && res[r][c] == inf
}

func inBound(r, c int) bool {
	return r >= 0 && r < m && c >= 0 && c < n
}

func genMemo(rows, columns int) [][]int {
	res := make([][]int, rows)
	for r := 0; r < rows; r++ {
		res[r] = make([]int, columns)
		for c := 0; c < columns; c++ {
			res[r][c] = inf
		}
	}
	return res
}

func calByStoneDist(dist1, dist2 [][]int) int {
	res := inf
	for _, stone := range oPoses {
		r, c := stone[0], stone[1]
		if dist1[r][c] == inf || dist2[r][c] == inf {
			continue
		}
		res = min(res, dist1[r][c]+dist2[r][c])
	}
	return res
}

func dp(distMemo [][]int) int {
	mLen := len(mPoses)
	total := 1 << mLen
	// memo(state, i)表示在机关i处，触发状态为state的最小步数
	memo := genMemo(total, mLen)
	for i := range mPoses {
		// 起点经过某个石头堆到机关i的最小距离
		memo[1<<i][i] = distMemo[i][mLen]
	}
	for curState := 1; curState < total; curState++ {
		for i := range mPoses {
			if curState&(1<<i) == 0 {
				continue
			}
			for j := range mPoses {
				if curState&(1<<j) != 0 {
					continue
				}
				nextState := curState | (1 << j)
				memo[nextState][j] = min(memo[nextState][j], memo[curState][i]+distMemo[i][j])
			}
		}
	}
	res := inf
	final := total - 1
	for i := range mPoses {
		res = min(res, memo[final][i]+distMemo[i][mLen+1])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
