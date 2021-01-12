package elimination

import (
	"container/list"
	"math"
)

/*
1293. 网格中的最短路径 https://leetcode-cn.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/

给你一个 m * n 的网格，其中每个单元格不是 0（空）就是 1（障碍物）。每一步，您都可以在空白单元格中上、下、左、右移动。
如果您 最多 可以消除 k 个障碍物，请找出从左上角 (0, 0) 到右下角 (m-1, n-1) 的最短路径，并返回通过该路径所需的步数。
如果找不到这样的路径，则返回 -1。

示例 1：
输入：
grid =
[[0,0,0],
 [1,1,0],
 [0,0,0],
 [0,1,1],
 [0,0,0]],
k = 1
输出：6
解释：
不消除任何障碍的最短路径是 10。
消除位置 (3,2) 处的障碍后，最短路径是 6 。该路径是 (0,0) -> (0,1) -> (0,2) -> (1,2) -> (2,2) -> (3,2) -> (4,2).


示例 2：
输入：
grid =
[[0,1,1],
 [1,1,1],
 [1,0,0]],
k = 1
输出：-1
解释：
我们至少需要消除两个障碍才能找到这样的路径。

提示：

grid.length == m
grid[0].length == n
1 <= m, n <= 40
1 <= k <= m*n
grid[i][j] == 0 or 1
grid[0][0] == grid[m-1][n-1] == 0
*/

/*
直接DFS会有用例超时，加上备忘录
另一个有意思的点是，需要先限定只向右或向下走计算下最少步数，如果无法到达才尝试每一步上下左右走的限定

当然，无论DFS还是BFS，都可以扩展为任意指定起点、终点，不一定就是左上角和右下角
注意任意指定后，对于DFS，要先判断终点在起点的哪个方向
*/
func shortestPath0(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	// 设想一种情况，只向右或向下走，将会是最短路径，共（m-1）+（n-1）步，已知起点终点都不是障碍，那么这条最短路径上最多有障碍m+n-3个
	// 如果k不小于m+n-3就可以走这条最短路径
	if k >= m+n-3 {
		return m + n - 2
	}

	visited := genVisited(m, n)
	memo := genMemo(m, n)
	dirs := [][]int{{1, 0}, {0, 1}} // 先尝试只向右或向下走

	// 返回从起点走到（r，c）处，再走到终点最少需要多少步
	// k 代表还能清理的障碍数， passed代表已经走过的步数
	var dfs func(r, c, k, passed int) int
	dfs = func(r, c, k, passed int) int {
		if r < 0 || r >= m || c < 0 || c >= n ||
			visited[r][c] || k < 0 || grid[r][c] == 1 && k == 0 {
			return math.MaxInt32
		}
		if r == m-1 && c == n-1 {
			return passed
		}
		if memo[r][c] < math.MaxInt32 {
			return memo[r][c]
		}
		visited[r][c] = true
		if grid[r][c] == 1 {
			k--
		}
		for _, d := range dirs {
			memo[r][c] = min(memo[r][c], dfs(r+d[0], c+d[1], k, passed+1))
		}
		// 回溯
		visited[r][c] = false
		return memo[r][c]
	}

	r := dfs(0, 0, k, 0)
	if r == math.MaxInt32 {
		visited = genVisited(m, n)
		memo = genMemo(m, n)
		// 尝试每次可以上下左右走
		dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
		r = dfs(0, 0, k, 0)
	}
	if r == math.MaxInt32 {
		return -1
	}
	return r
}

func genVisited(m, n int) [][]bool {
	r := make([][]bool, m)
	for i := range r {
		r[i] = make([]bool, n)
	}
	return r
}

func genMemo(m, n int) [][]int {
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = math.MaxInt32
		}
	}
	return memo
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
BFS，除了横坐标、纵坐标，需要把剩余能清理的障碍数量作为另一个状态传递给bfs的队列，方便判断是否能继续走
时空复杂度都是O(MN∗min(M+N,K))
*/
func shortestPath(grid [][]int, k int) int {
	const maxSize = 40
	const maxK = maxSize + maxSize - 3
	m, n := len(grid), len(grid[0])
	// 设想一种情况，只向右或向下走，将会是最短路径，共（m-1）+（n-1）步，已知起点终点都不是障碍，那么这条最短路径上最多有障碍m+n-3个
	// 如果k不小于m+n-3就可以走这条最短路径
	if k >= m+n-3 {
		return m + n - 2
	}
	visited := [maxSize][maxSize][maxK + 1]bool{}
	queue := list.New()
	queue.PushBack([]int{0, 0, k})
	visited[0][0][k] = true
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	passed := 0 // paased代表之前已经走了多少步
	for queue.Len() > 0 {
		levelSize := queue.Len()
		for i := 0; i < levelSize; i++ { // 当前层
			info := queue.Remove(queue.Front()).([]int)
			r, c, k := info[0], info[1], info[2]
			if r == m-1 && c == n-1 {
				return passed
			}
			if grid[r][c] == 1 {
				k--
			}
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr < 0 || nr >= m || nc < 0 || nc >= n || k < 0 || visited[nr][nc][k] {
					continue
				}
				queue.PushBack([]int{nr, nc, k})
				visited[nr][nc][k] = true
			}
		}
		passed++ // 一层处理完，步数加1
	}
	return -1
}
