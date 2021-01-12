/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package find_the_duplicate_number

/*
给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。
假设只有一个重复的整数，找出这个重复的数。

示例 1:

输入: [1,3,4,2,2]
输出: 2
示例 2:

输入: [3,1,3,4,2]
输出: 3
说明：

不能更改原数组（假设数组是只读的）。
只能使用额外的 O(1) 的空间。
时间复杂度小于 O(n^2) 。
数组中只有一个重复的数字，但它可能不止重复出现一次。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-the-duplicate-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/* 如果将数组排序，或者借助一个set，问题就会容易解决
但是题目里的约束条件限制了这些做法

1. 一个二分的思想：
数组的值域是[1, n]
对于一个在区间内的数字x，统计nums里<=x的元素个数count
各个小于等于x的数字在数组里最多有一个(可能没有) <=> count <= x <=> 重复元素在(x, n]中；
否则count > x <=> 重复元素在[1, x]中

注意这里的二分是针对值域[1, n]，而不是对数组nums
nums只是一个集合，每次统计小于某个数字x的时候用

时间复杂度O(nlogn), 空间复杂度O(1)
*/
func findDuplicate(nums []int) int {
	left, right := 1, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if countLowers(nums, mid) <= mid {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// 返回nums中不大于x的元素的个数
func countLowers(nums []int, x int) int {
	count := 0
	for _, v := range nums {
		if v <= x {
			count++
		}
	}
	return count
}

/*
2. 弗洛伊德的兔子和乌龟
*/
func findDuplicate1(nums []int) int {
	// Find the intersection point of the two runners.
	slow := nums[nums[0]]
	fast := nums[slow]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	// Find the "entrance" to the cycle.
	p := nums[0]
	q := slow
	for p != q {
		p = nums[p]
		q = nums[q]
	}
	return p
}
