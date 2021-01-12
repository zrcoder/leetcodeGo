/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package search_in_rotated_sorted_array

/*
33. 搜索旋转排序数组
https://leetcode-cn.com/problems/search-in-rotated-sorted-array

假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1:
输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
示例 2:
输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1
*/

// 经典二分法（模板一）
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[0] <= target: // target is in the left part
			if nums[mid] >= nums[0] && nums[mid] < target { // only nums[mid] in the left part, move left, else move right
				left = mid + 1
			} else {
				right = mid - 1
			}
		default: // target is in the right part
			if nums[mid] < nums[0] && nums[mid] > target { // only nums[mid] in the right part, move right, else move left
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

// 高级二分法（模板二）
func search1(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[0] <= target:
			if nums[mid] >= nums[0] && nums[mid] < target {
				left = mid + 1
			} else {
				right = mid
			}
		default:
			if nums[mid] < nums[0] && nums[mid] > target {
				right = mid
			} else {
				left = mid + 1
			}
		}
	}
	if left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

// 二分法（模板三）
func search2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[0] <= target:
			if nums[mid] >= nums[0] && nums[mid] < target {
				left = mid
			} else {
				right = mid
			}
		default:
			if nums[mid] < nums[0] && nums[mid] > target {
				right = mid
			} else {
				left = mid
			}
		}
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	return -1
}
