package _sum

import "sort"

/*
15. 三数之和
https://leetcode-cn.com/problems/3sum

给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
/*
先简化下问题：
如果在一个已经排序的数组里，找到两个数字，其和为定值呢？
可以用左右两个指针，不断向中间凑：
两个指针处的值相加，如果和相等，则找到了一对；
如果和小于目标，左指针向右移动
如果和大于目标，右指针向左移动
当然，为了避免重复结果，指针移动不一定是一步，只要和当前处的值相同就一直移动，直到到达不同的值
比如对于[2,2,2,5,8], 假设当前左指针在最左边，需要向右移动，那么一直移动到5

扩展到这个三数之和问题，可以先将数组排序，遍历数组，注意重复元素跳过
对于遍历到索引i处， 在nums[i+1:]里边使用上面的双指针技巧找到两个和为-nums[i]的就行了
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}
		target := -v
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{v, nums[left], nums[right]})
				left = moveLeft(left, nums)
				right = moveRight(right, nums)
			} else if sum < target {
				left = moveLeft(left, nums)
			} else {
				right = moveRight(right, nums)
			}
		}
	}
	return result
}

func moveLeft(left int, nums []int) int {
	v := nums[left]
	for left < len(nums) && v == nums[left] {
		left++
	}
	return left
}

func moveRight(right int, nums []int) int {
	v := nums[right]
	for right >= 0 && v == nums[right] {
		right--
	}
	return right
}
