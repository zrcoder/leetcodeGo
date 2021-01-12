/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package find_minimum_in_rotated_sorted_array

/*
33. 搜索旋转排序数组 https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

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
/*
二分法， 对于mid，根据nums[0]和mid的大小关系，能方便获知mid左半侧有序还是右半侧有序
再根据target是否在有序一侧决定left和right的移动
*/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		curr := nums[mid]
		if curr == target {
			return mid
		}
		if nums[0] <= curr { // nums[0:mid+1] 有序; mid可能为0
			if nums[0] <= target && target < curr {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // nums[mid+1:] 有序
			if curr < target && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

/*
153. 寻找旋转排序数组中的最小值 https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array

假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

请找出其中最小的元素。

你可以假设数组中不存在重复元素。

示例 1:

输入: [3,4,5,1,2]
输出: 1
示例 2:

输入: [4,5,6,7,0,1,2]
输出: 0
*/
// 朴素实现
func findMin00(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return nums[i]
		}
	}
	return nums[0]
}

/*
二分法1，每次将mid和right处的值比较，以判断mid落在旋转点左侧还是右侧
*/
func findMin0(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	if nums[0] <= nums[right] {
		return nums[0]
	}
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}
		if nums[mid] > nums[right] { // mid落在旋转点左侧; 改成nums[mid] > nums[0]也行
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 二分法2
func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] { // mid落在旋转点左侧
			left = mid + 1
		} else { // mid和right在旋转点同侧，但因为一开始right在整个数组最右，所以当前只可能同在旋转点右侧
			right = mid
		}
	}
	return nums[left]
}

// 分治法，类似二分法
func findMin01(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}
	end := len(nums) - 1
	mid := end / 2
	if nums[mid] > nums[end] {
		return findMin(nums[mid+1:])
	}
	if nums[mid] < nums[end] {
		return findMin(nums[:mid+1])
	}
	return -1 // 实际到不了这里
}

/*
题目变体：

154. 寻找旋转排序数组中的最小值 II https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii
如果nums里有重复元素呢？如：
[3, 1, 2, 3]
[3, 1, 2, 2, 3]
*/
func findMin1(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] > nums[right]: // mid落在旋转点左侧
			left = mid + 1
		case nums[mid] < nums[right]: // mid和right在旋转点同侧，但因为一开始right在整个数组最右，所以当前只可能同在旋转点右侧
			right = mid
		default: // 相等时保守缩进，避免遗漏一些元素
			right--
		}
	}
	return nums[left]
}

// 或者用分治法，本质与二分法类似
func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	end := len(numbers) - 1
	mid := end / 2
	if numbers[mid] > numbers[end] {
		return minArray(numbers[mid+1:])
	}
	if numbers[mid] < numbers[end] {
		return minArray(numbers[:mid+1])
	}
	// 相等时保守缩进，避免遗漏一些元素
	return minArray(numbers[:end])
}
