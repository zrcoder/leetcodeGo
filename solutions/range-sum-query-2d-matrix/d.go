/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package range_sum_query_2d_matrix

type NumMatrix struct {
	rowPrefixSum [][]int // 记录每个元素在每一行的前缀和
	matrix       [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}
	m, n := len(matrix), len(matrix[0])
	rowSum := make([][]int, m)
	for r := 0; r < m; r++ {
		rowSum[r] = make([]int, n)
		rowSum[r][0] = matrix[r][0]
		for c := 1; c < n; c++ {
			rowSum[r][c] = rowSum[r][c-1] + matrix[r][c]
		}
	}
	return NumMatrix{
		rowPrefixSum: rowSum,
		matrix:       matrix,
	}
}

func (n *NumMatrix) Update(row int, col int, val int) {
	n.matrix[row][col] = val
	c := col
	if col == 0 {
		n.rowPrefixSum[row][col] = val
		c = 1
	}
	for ; c < len(n.matrix[0]); c++ {
		n.rowPrefixSum[row][c] = n.rowPrefixSum[row][c-1] + n.matrix[row][c]
	}
}

func (n *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for r := row1; r <= row2; r++ {
		if col1 == 0 {
			sum += n.rowPrefixSum[r][col2]
		} else {
			sum += n.rowPrefixSum[r][col2] - n.rowPrefixSum[r][col1-1]
		}
	}
	return sum
}
