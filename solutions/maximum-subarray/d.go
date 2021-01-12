/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package subarray

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp, res := 0, nums[0]
	for _, v := range nums {
		dp = max(0, dp) + v
		res = max(res, dp)
	}
	return res
}

func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp, res := 0, nums[0]
	for _, v := range nums {
		dp += v
		res = max(res, dp)
		if dp < 0 {
			dp = 0
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
