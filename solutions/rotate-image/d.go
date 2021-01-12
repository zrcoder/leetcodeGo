/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package rotate_image

func rotate(matrix [][]int) {
	n := len(matrix)
	for r := 0; r < (n+1)/2; r++ {
		for c := 0; c < n/2; c++ {
			// p1, p4, p3, p2
			matrix[r][c], matrix[n-1-c][r], matrix[n-1-r][n-1-c], matrix[c][n-1-r] =
				// p4, p3, p2, p1
				matrix[n-1-c][r], matrix[n-1-r][n-1-c], matrix[c][n-1-r], matrix[r][c]
		}
	}
}

func rotateAnticlockwise(s [][]int) {
	n := len(s)
	for r := 0; r < (n+1)/2; r++ {
		for c := 0; c < n/2; c++ {
			s[r][c], s[n-1-c][r], s[n-1-r][n-1-c], s[c][n-1-r] =
				s[c][n-1-r], s[r][c], s[n-1-c][r], s[n-1-r][n-1-c]
		}
	}
}
