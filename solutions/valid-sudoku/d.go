/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package valid_sudoku

/*

36. 有效的数独 https://leetcode-cn.com/problems/valid-sudoku
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

示例 1:
输入:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: true

示例 2:
输入:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: false

解释: 除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。
	 但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。

说明:
一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
给定数独序列只包含数字 1-9 和字符 '.' 。
给定数独永远是 9x9 形式的。
*/

type BSet map[byte]struct{}

func (s BSet) Has(v byte) bool {
	_, ok := s[v]
	return ok
}
func (s BSet) Add(v byte) {
	s[v] = struct{}{}
}

// 1.
// primary implemention
func isValidSudoku(board [][]byte) bool {
	const n = 9
	// check rows
	for _, row := range board {
		set := make(BSet, 0)
		for _, v := range row {
			if ok := check(set, v); !ok {
				return false
			}
		}
	}
	//check clomns
	for c := 0; c < n; c++ {
		set := make(BSet, 0)
		for r := 0; r < n; r++ {
			if ok := check(set, board[r][c]); !ok {
				return false
			}
		}
	}
	//check grids
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// (i, j) aims at one of total 9 grids
			set := make(BSet, 0)
			if ok := checkGrid(board, i, j, set); !ok {
				return false
			}
		}
	}
	return true
}

func checkGrid(board [][]byte, i, j int, set BSet) bool {
	minRow, minClonm := i*3, j*3
	for r := minRow; r < minRow+3; r++ {
		for c := minClonm; c < minClonm+3; c++ {
			if ok := check(set, board[r][c]); !ok {
				return false
			}
		}
	}
	return true
}

func check(set BSet, v byte) bool {
	if v == '.' {
		return true
	}
	if set.Has(v) {
		return false
	}
	set.Add(v)
	return true
}

// 2.
// we can generate all sets at begin
// and then range the board only once
func isValidSudoku1(board [][]byte) bool {
	const n = 9
	rowSets, clonmSets, gridSets := make([]BSet, n), make([]BSet, n), make([]BSet, n)
	for i := 0; i < n; i++ {
		rowSets[i], clonmSets[i], gridSets[i] = make(BSet, 0), make(BSet, 0), make(BSet, 0)
	}
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			ch := board[r][c]
			if ch == '.' {
				continue
			}
			gridIndex := (r/3)*3 + c/3
			if rowSets[r].Has(ch) ||
				clonmSets[c].Has(ch) ||
				gridSets[gridIndex].Has(ch) {
				return false
			}
			rowSets[r].Add(ch)
			clonmSets[c].Add(ch)
			gridSets[gridIndex].Add(ch)
		}
	}
	return true
}

// 3.
// wo have rows in [0, 8], clonms in [0, 8] and grids in [0, 8]
// for each row/clonm/grid, we can find out an area who's rows in [r1, r2] and clonms in [c1, c2]
// and we can define a function to check the area
func isValidSudoku2(board [][]byte) bool {
	const n = 9
	for i := 0; i < n; i++ {
		minR := i / 3 * 3
		minC := i % 3 * 3
		if !isAreaValid(i, i, 0, n-1, board) || // check row i
			!isAreaValid(0, n-1, i, i, board) || // check clonm i
			!isAreaValid(minR, minR+2, minC, minC+2, board) { // check grid i
			return false
		}
	}
	return true
}

// check the area who's rows in [r1, r2] and clonms in [c1, c2]
func isAreaValid(r1, r2, c1, c2 int, board [][]byte) bool {
	set := make(BSet, 0)
	for r := r1; r <= r2; r++ {
		for c := c1; c <= c2; c++ {
			ch := board[r][c]
			if ch == '.' {
				continue
			}
			if set.Has(ch) {
				return false
			}
			set.Add(ch)
		}
	}
	return true
}
