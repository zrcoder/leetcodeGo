package solver

/*
37. 解数独
https://leetcode-cn.com/problems/sudoku-solver

编写一个程序，通过已填充的空格来解决数独问题。

一个数独的解法需遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
空白格用 '.' 表示。

Note:

给定的数独序列只包含数字 1-9 和字符 '.' 。
你可以假设给定的数独只有唯一解。
给定数独永远是 9x9 形式的。
*/
const n = 9

var grids, rows, cols [n][n + 1]bool
var blanks [][]int

func solveSudoku(board [][]byte) {
	// 每个用例不独立，全局变量在每次都需要重新初始化
	blanks = nil
	grids, rows, cols = [n][n + 1]bool{}, [n][n + 1]bool{}, [n][n + 1]bool{}

	initSlices(board)
	_ = dfs(0, board)
}

func initSlices(board [][]byte) {
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if board[r][c] == '.' {
				blanks = append(blanks, []int{r, c})
			} else {
				num := int(board[r][c] - '0')
				rows[r][num] = true
				cols[c][num] = true
				grids[(r/3)*3+c/3][num] = true
			}
		}
	}
}

func dfs(i int, board [][]byte) bool {
	if i == len(blanks) {
		return true
	}
	r, c := blanks[i][0], blanks[i][1]
	for num := 1; num <= n; num++ {
		if rows[r][num] || cols[c][num] || grids[(r/3)*3+c/3][num] {
			continue
		}
		rows[r][num] = true
		cols[c][num] = true
		grids[(r/3)*3+c/3][num] = true
		board[r][c] = byte(num + int('0'))
		if dfs(i+1, board) {
			return true
		}
		board[r][c] = '.'
		rows[r][num] = false
		cols[c][num] = false
		grids[(r/3)*3+c/3][num] = false
	}
	return false
}
