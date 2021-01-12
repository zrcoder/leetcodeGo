/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package search_a_2d_matrix

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	i := sort.Search(m*n, func(i int) bool {
		return matrix[i/n][i%n] >= target
	})
	return i != m*n && matrix[i/n][i%n] == target
}
