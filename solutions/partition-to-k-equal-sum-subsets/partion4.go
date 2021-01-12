/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package partition_to_k_equal_sum_subsets

func makesquare(nums []int) bool {
	const squareEdges = 4
	if len(nums) < squareEdges {
		return false
	}
	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	target := sum / squareEdges
	if sum%squareEdges != 0 || max > target {
		return false
	}
	used := make([]bool, len(nums))
	k := squareEdges
	return backTracking(k, 0, 0, target, nums, used)
}
