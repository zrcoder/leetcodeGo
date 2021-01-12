package minimum_size_subarray_sum

/*
209. 长度最小的子数组 https://leetcode-cn.com/problems/minimum-size-subarray-sum/
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组，
并返回其长度。如果不存在符合条件的连续子数组，返回 0。

示例:

输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
进阶:

如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
*/
/*
滑动窗口，用left和right两个指针指向窗口的左右边界
窗口总体向右滑动，每次right加1， 窗口中所有元素和大于等于s后，更新结果，
并开始让left向右移动，直到和小于s，移动过程中也需要更新结果

时间复杂度 O(n), 最坏情况left和right都遍历一遍数组
空间复杂度 O(1)
*/
func minSubArrayLen(s int, nums []int) int {
	result := len(nums) + 1
	sum := 0
	left, right := 0, 0
	for right < len(nums) {
		sum += nums[right]
		for sum >= s {
			result = min(result, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if result == len(nums)+1 {
		return 0
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
