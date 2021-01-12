/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package check_if_a_number_is_majority_element_in_a_sorted_array

/*
1150. 检查一个数是否在数组中占绝大多数
https://leetcode-cn.com/problems/check-if-a-number-is-majority-element-in-a-sorted-array

给出一个按 非递减 顺序排列的数组 nums，和一个目标数值 target。
假如数组 nums 中绝大多数元素的数值都等于 target，则返回 True，否则请返回 False。
所谓占绝大多数，是指在长度为 N 的数组中出现必须 超过 N/2 次。

示例 1：

输入：nums = [2,4,5,5,5,5,5,6,6], target = 5
输出：true
解释：
数字 5 出现了 5 次，而数组的长度为 9。
所以，5 在数组中占绝大多数，因为 5 次 > 9/2。

示例 2：

输入：nums = [10,100,101,101], target = 101
输出：false
解释：
数字 101 出现了 2 次，而数组的长度是 4。
所以，101 不是 数组占绝大多数的元素，因为 2 次 = 4/2。
*/
/*
用两次二分法找到数组中最左边target的索引和最右边的索引，两个索引距离+1即target出现的个数
时间复杂度O(lgn)，空间复杂度O(1)
*/
func isMajorityElement(nums []int, target int) bool {
	left := search(nums, target) // 也可以用标准库sort包的Search方法，达到的效果一样
	if left == len(nums) || nums[left] != target {
		return false
	}
	right := searchFromRight(nums, target)
	return right-left > len(nums)/2
}

/*
在nums里搜索target，返回新target应该插入的位置；如果nums里已经有target，则在第一个已有target元素之前插入
nums已经排序，但可能有重复元素
功能同标准库里的sort.SearchInts(nums, target)
*/
func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] >= target:
			right = mid
		}
	}
	return left
}

/*
在nums里从右向左搜索target，返回新target应该插入的位置；如果nums里已经有target，则在最后一个target元素之后插入
nums已经排序，但可能有重复元素
*/
func searchFromRight(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] <= target:
			left = mid + 1
		case nums[mid] > target:
			right = mid
		}
	}
	return left
}
