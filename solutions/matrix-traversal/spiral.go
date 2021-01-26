/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package matrix_traversal

import "math"

// 1。 直觉遍历，时空复杂度都是O(m*n)
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	const right, down, left, up = 0, 1, 2, 3
	m, n := len(matrix), len(matrix[0])
	rMin, rMax, cMin, cMax := 0, m-1, 0, n-1 // 上下左右边界
	r, c, direct := 0, 0, right
	res := make([]int, m*n)
	for i := range res {
		res[i] = matrix[r][c]
		switch direct {
		case right:
			if c < cMax {
				c++
			} else {
				direct, rMin, r = down, rMin+1, r+1
			}
		case down:
			if r < rMax {
				r++
			} else {
				direct, cMax, c = left, cMax-1, c-1
			}
		case left:
			if c > cMin {
				c--
			} else {
				direct, rMax, r = up, rMax-1, r-1
			}
		case up:
			if r > rMin {
				r--
			} else {
				direct, cMin, c = right, cMin+1, c+1
			}
		}
	}
	return res
}

// 2。 另一个实现,思路同上一个实现，时空复杂度都是O(m*n)
func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	rMin, rMax, cMin, cMax := 0, m-1, 0, n-1 // 上下左右边界
	res := make([]int, 0, m*n)

	for levels := (min(m, n) + 1) / 2; levels > 0; levels-- { // 层数
		for i := cMin; i <= cMax; i++ { // 向右遍历
			res = append(res, matrix[rMin][i])
		}
		rMin++ //重新设定上边界，如果发现上边界还在下边界之下，说明遍历完成了，下同
		if rMin > rMax {
			break
		}
		for i := rMin; i <= rMax; i++ { // 向下遍历
			res = append(res, matrix[i][cMax])
		}
		cMax--
		if cMax < cMin {
			break
		}
		for i := cMax; i >= cMin; i-- { // 向左遍历
			res = append(res, matrix[rMax][i])
		}
		rMax--
		if rMax < rMin {
			break
		}
		for i := rMax; i >= rMin; i-- { // 向上遍历
			res = append(res, matrix[i][cMin])
		}
		cMin++
		if cMin > cMax {
			break
		}
		levels--
	}
	return res
}

// 3。从外部向内部逐层遍历打印矩阵，最外面一圈打印完，里面仍然是一个矩阵;时空复杂度都是O(m*n)
func spiralOrder2(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	levels := (min(m, n) + 1) / 2 //层数
	res := make([]int, 0, m*n)
	// 从外层向内层层层遍历
	for i := 0; i < levels; i++ {
		for c := i; c < n-i; c++ { // 向右
			res = append(res, matrix[i][c])
		}
		for r := i + 1; r < m-i; r++ { // 向下
			res = append(res, matrix[r][n-1-i])
		}
		for c := n - 1 - (i + 1); m-1-i != i && c >= i; c-- { // 向左；可能这一层只有一行，注意判断m-1-i 与 i是否相等
			res = append(res, matrix[m-1-i][c])
		}
		for r := m - 1 - (i + 1); n-1-i != i && r > i; r-- { // 向上；可能这一层只有一列，注意判断n-1-i 与 i是否相等
			res = append(res, matrix[r][i])
		}
	}
	return res
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
