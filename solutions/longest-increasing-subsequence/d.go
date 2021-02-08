/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package longest_increasing_subsequence

import "sort"

/*
300. 最长上升子序列 https://leetcode-cn.com/problems/longest-increasing-subsequence

给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:
输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。

说明:
可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n^2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
*/

/*
动态规划，时间复杂度O(n^2),  空间复杂度O(n)
*/
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums)) // dp[i]代表以nums[i]结尾的子序列长度
	maxLen := 0
	for i, v := range nums {
		dp[i] = 1 // 一个元素算递增长度为1
		for j := 0; j < i; j++ {
			if nums[j] < v {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
贪心+二分搜索，二维的（最长上升子序列）lis问题：
如果要使上升子序列尽可能长，则需要让序列上升得尽可能慢，因此我们希望每次在上升子序列最后加上的那个数尽可能的小。
建立 memo 数组，memo[i]代表长度为 i+1 的递增子序列末尾数字
遍历 nums，对于当前元素：
如果大于结果数组最后元素，直接追加到结果数组最后；
否则，在结果数组里找到第一个不小于当前元素的元素，并将其更新为当前元素。

时间复杂度O(nlogn), 空间复杂度O(n)
*/
func lengthOfLIS0(nums []int) int {
	memo := make([]int, len(nums))
	length := 0
	for _, v := range nums {
		j := sort.Search(length, func(i int) bool {
			return memo[i] >= v
		})
		memo[j] = v
		if j == length {
			length++
		}
	}
	return length
}

// 如果允许修改nums，result数组可以省略
func lengthOfLIS00(nums []int) int {
	length := 0
	for _, v := range nums {
		j := sort.Search(length, func(i int) bool {
			return nums[i] >= v
		})
		nums[j] = v
		if j == length {
			length++
		}
	}
	return length
}
