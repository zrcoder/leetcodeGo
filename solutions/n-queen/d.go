/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package n_queen

/*
51. N皇后
https://leetcode-cn.com/problems/n-queens

n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
示例:
输入: 4
输出: [
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]
解释: 4 皇后问题存在两个不同的解法。
*/
/*
常规回溯
时间复杂度O(n!)；空间复杂度O(n)
*/
func solveNQueens(n int) [][]string {
	var res [][]string
	board := makeBoard(n)
	var backtrack func(r int)
	// 在行r找到合适的列放置皇后
	backtrack = func(r int) {
		if r == len(board) {
			res = append(res, parse(board))
			return
		}
		for c := 0; c < len(board); c++ {
			if !canSetQueen(board, r, c) {
				continue
			}
			board[r][c] = 'Q'
			backtrack(r + 1)
			board[r][c] = '.'
		}
	}
	backtrack(0)
	return res
}

func makeBoard(n int) [][]byte {
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	return board
}

func parse(board [][]byte) []string {
	r := make([]string, len(board))
	for i := range r {
		r[i] = string(board[i])
	}
	return r
}

func canSetQueen(board [][]byte, r, c int) bool {
	var i, j int
	for i = 0; i < r; i++ { // top
		if board[i][c] == 'Q' {
			return false
		}
	}
	for i, j = r-1, c-1; i >= 0 && j >= 0; i, j = i-1, j-1 { // topLeft
		if board[i][j] == 'Q' {
			return false
		}
	}
	for i, j = r-1, c+1; i >= 0 && j < len(board); i, j = i-1, j+1 { // topRight
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}

/*
优化
canSetQueen() 函数耗费的时间可以优化掉
用三个哈希表，分别记录左上右下对角线及列上是不是有皇后会将判断某个位置能否放置皇后变成常数级复杂度
column 记录列上是否已经有皇后，列的取值范围是[0,n]
dia1 记录左上右下方向对角线上是否有皇后，这样的对角线上位置的行列坐标差是定值
（可以想象，从这样的对角线左上端点向右下端点走，每次都是行+1且列+1，所以线上所有位置行号-列号是个定值）
当然每条对角线行号-列号的结果会不同；这样可以用行号-列号的值作为这类对角线的唯一标识，且其取值范围是[1-n, n-1]
类似，dia2 记录左下右上对角线上是否有皇后, 可以用行号+列号唯一标识这样的对角线，且其取值范围是[0, 2n-2]

进一步，这三个哈希表，可以改成数组实现，这样速度更快，在 n 较大时，这个优化效果会很显著

总的时空复杂度同上一解法
*/
func solveNQueens1(n int) [][]string {
	var res [][]string
	memo := make([]int, n)

	column := make([]bool, n)
	dia1 := make([]bool, 2*n-1)
	dia2 := make([]bool, 2*n-1)

	var backtrack func(r int)
	backtrack = func(r int) {
		if r == n {
			res = append(res, parseMemo(memo))
			return
		}
		for c := 0; c < n; c++ {
			if column[c] || dia1[r-c+n-1] || dia2[r+c] {
				continue
			}
			memo[r] = c // 这一步可以不回溯，后边如果回溯会被重新赋值
			column[c], dia1[r-c+n-1], dia2[r+c] = true, true, true
			backtrack(r + 1)
			column[c], dia1[r-c+n-1], dia2[r+c] = false, false, false
		}
	}

	backtrack(0)
	return res
}

func parseMemo(memo []int) []string {
	n := len(memo)
	res := make([]string, n)
	for i := range res {
		row := make([]byte, n)
		for j := range row {
			row[j] = '.'
		}
		row[memo[i]] = 'Q'
		res[i] = string(row)
	}
	return res
}

/*
52. N皇后 II
https://leetcode-cn.com/problems/n-queens-ii

与问题51类似，更简单些，只需要返回不同的解决方案的数量
*/
func totalNQueens(n int) int {
	total := 0

	column := make([]bool, n)
	dia1 := make([]bool, 2*n-1)
	dia2 := make([]bool, 2*n-1)

	var backtrack func(r int)
	backtrack = func(r int) {
		if r == n {
			total++
			return
		}
		for c := 0; c < n; c++ {
			if column[c] || dia1[r-c+n-1] || dia2[r+c] {
				continue
			}
			column[c], dia1[r-c+n-1], dia2[r+c] = true, true, true
			backtrack(r + 1)
			column[c], dia1[r-c+n-1], dia2[r+c] = false, false, false
		}
	}

	backtrack(0)
	return total
}
