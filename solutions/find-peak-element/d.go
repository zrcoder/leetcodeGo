/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package find_peak_element

import (
	"sort"
)

/*
162. 寻找峰值
https://leetcode-cn.com/problems/find-peak-element

峰值元素是指其值大于左右相邻值的元素。

给定一个输入数组 nums，其中 nums[i] ≠ nums[i+1]，找到峰值元素并返回其索引。

数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。

你可以假设 nums[-1] = nums[n] = -∞。

示例 1:

输入: nums = [1,2,3,1]
输出: 2
解释: 3 是峰值元素，你的函数应该返回其索引 2。
示例 2:

输入: nums = [1,2,1,3,5,6,4]
输出: 1 或 5
解释: 你的函数可以返回索引 1，其峰值元素为 2；
     或者返回索引 5， 其峰值元素为 6。
说明:

你的解法应该是 O(logN) 时间复杂度的。
*/
/*
朴素实现
遍历nums，如果发现nums[i] > nums[i+1]，则i为所求；找不到则为最后一个索引
时间复杂度O(n）
*/
func findPeakElement0(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}

/*
二分法
对比nums[mid]和其右边的元素，如果nums[mid]大，说明峰值在mid左侧，包括mid；否则峰值在mid右侧
时间复杂度O(lgn)
*/
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid+1] > nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// 使用标准库，减少代码量
func findPeakElement4(nums []int) int {
	return sort.Search(len(nums)-1, func(i int) bool {
		return nums[i+1] <= nums[i]
	})
}
