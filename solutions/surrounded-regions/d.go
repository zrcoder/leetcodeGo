package regions

import "container/list"

/*
130. 被围绕的区域


给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。

找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

示例:

X X X X
X O O X
X X O X
X O X X
运行你的函数后，矩阵变为：

X X X X
X X X X
X X X X
X O X X

解释:

被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。
任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。
如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
*/
/*
从四个边界开始，找到‘O’并向里扩展标记连成一片的‘O’们，用dfs或bfs都可以
*/
// dfs
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	marked := make([][]bool, m) // 标记与边界相连的‘O’们
	for i := range marked {
		marked[i] = make([]bool, n)
	}
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n ||
			board[r][c] == 'X' || marked[r][c] {
			return
		}
		marked[r][c] = true
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	// 上下边界
	for c := 0; c < n; c++ {
		dfs(0, c)
		dfs(m-1, c)
	}
	// 左右边界
	for r := 1; r < m-1; r++ {
		dfs(r, 0)
		dfs(r, n-1)
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if board[r][c] == 'O' && !marked[r][c] {
				board[r][c] = 'X'
			}
		}
	}
}

/*
上边用了一个矩阵标记与边界‘O'相连的’O'，可以优化，
直接在原矩阵里将与边界相连的‘O'改成另一个字符，不和’X‘、’O‘相同即可,最后还原成‘O'就行
*/
func solve1(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n ||
			board[r][c] == 'X' || board[r][c] == 'I' {
			return
		}
		board[r][c] = 'I'
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	// 上下边界
	for c := 0; c < n; c++ {
		dfs(0, c)
		dfs(m-1, c)
	}
	// 左右边界
	for r := 1; r < m-1; r++ {
		dfs(r, 0)
		dfs(r, n-1)
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			} else if board[r][c] == 'I' {
				board[r][c] = 'O'
			}
		}
	}
}

// bfs
func solve2(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	marked := make([][]bool, m) // 标记与边界相连的‘O’们
	for i := range marked {
		marked[i] = make([]bool, n)
	}
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	bfs := func(r, c int) {
		queue := list.New()
		queue.PushBack([]int{r, c})
		marked[r][c] = true
		for queue.Len() > 0 {
			pos := queue.Remove(queue.Front()).([]int)
			for _, d := range dirs {
				nextR, nextC := pos[0]+d[0], pos[1]+d[1]
				if nextR < 0 || nextC < 0 || nextR >= m || nextC >= n ||
					board[nextR][nextC] == 'X' || marked[nextR][nextC] {
					continue
				}
				queue.PushBack([]int{nextR, nextC})
				marked[nextR][nextC] = true
			}
		}
	}
	// 上下边界
	for c := 0; c < n; c++ {
		if board[0][c] == 'O' {
			bfs(0, c)
		}
		if board[m-1][c] == 'O' {
			bfs(m-1, c)
		}
	}
	// 左右边界
	for r := 1; r < m-1; r++ {
		if board[r][0] == 'O' {
			bfs(r, 0)
		}
		if board[r][n-1] == 'O' {
			bfs(r, n-1)
		}
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if board[r][c] == 'O' && !marked[r][c] {
				board[r][c] = 'X'
			}
		}
	}
}
