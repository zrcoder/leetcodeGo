/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package longest_continuous_subarray_with_absolute_diff_less_than_or_equal_to_limit

import (
	"container/heap"
	"github.com/zrcoder/leetcodeGo/util/intheap"
)

/*
1438. 绝对差不超过限制的最长连续子数组 https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
给你一个整数数组 nums ，和一个表示限制的整数 limit，
请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于 limit 。

如果不存在满足条件的子数组，则返回 0 。

示例 1：
输入：nums = [8,2,4,7], limit = 4
输出：2
解释：所有子数组如下：
[8] 最大绝对差 |8-8| = 0 <= 4.
[8,2] 最大绝对差 |8-2| = 6 > 4.
[8,2,4] 最大绝对差 |8-2| = 6 > 4.
[8,2,4,7] 最大绝对差 |8-2| = 6 > 4.
[2] 最大绝对差 |2-2| = 0 <= 4.
[2,4] 最大绝对差 |2-4| = 2 <= 4.
[2,4,7] 最大绝对差 |2-7| = 5 > 4.
[4] 最大绝对差 |4-4| = 0 <= 4.
[4,7] 最大绝对差 |4-7| = 3 <= 4.
[7] 最大绝对差 |7-7| = 0 <= 4.
因此，满足题意的最长子数组的长度为 2 。

示例 2：
输入：nums = [10,1,2,4,7,2], limit = 5
输出：4
解释：满足题意的最长子数组是 [2,4,7,2]，其最大绝对差 |2-7| = 5 <= 5 。

示例 3：
输入：nums = [4,2,2,2,4,4,2,2], limit = 0
输出：3

提示：
1 <= nums.length <= 10^5
1 <= nums[i] <= 10^9
0 <= limit <= 10^9
*/
/*
参考 [239] 滑动窗口最大值
朴素滑动窗口实现：
用左右两个指针left，right，两个指针确定了一个窗口[left, right]
起初两个指针都指向最左侧，根据情况向右移动左指针或右指针
如果窗口里的元素满足题意，即所有值的绝对查不超过limit，则移动right，否则移动left
为了判断窗口里的元素是否满足题意，需要遍历一遍窗口统计到最小值和最大值求差
最坏时间复杂度O(n^2),最后一个用例超时;空间复杂度O(1)
*/
func longestSubarray0(nums []int, limit int) int {
	result := 0
	left, right := 0, 0
	for right < len(nums) {
		if isValid(nums, left, right, limit) {
			result = max(result, right-left+1)
			right++
		} else {
			left++
		}
	}
	return result
}

func isValid(nums []int, left, right, limit int) bool {
	lo, hi := nums[left], nums[left]
	for i := left + 1; i <= right; i++ {
		if nums[i] < lo {
			lo = nums[i]
		} else if nums[i] > hi {
			hi = nums[i]
		}
	}
	return hi-lo <= limit
}

/*
isValid方法时间复杂度可以优化
要迅速找到窗口中的最大值和最小值，可以借助堆
用两个堆，一个大顶堆一个小顶堆，两个堆里的元素完全相同，都是窗口里的元素
每次根据两个堆堆顶元素的差值能够迅速获知窗口中所有元素是否满足题意
时间复杂度:平均O(n*lg(n))，最坏O(n^2)；空间复杂度O(n);实测双百通过
*/
func longestSubarray(nums []int, limit int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	minHeap := &intheap.Heap{}
	minHeap.InitWithCmp(func(i, j int) bool {
		return minHeap.Get(i) < minHeap.Get(j)
	})
	maxHeap := &intheap.Heap{}
	maxHeap.InitWithCmp(func(i, j int) bool {
		return maxHeap.Get(i) > maxHeap.Get(j)
	})
	heap.Push(minHeap, nums[0])
	heap.Push(maxHeap, nums[0])
	result := 1
	left, right := 0, 1
	for right < len(nums) {
		if maxHeap.Peek()-minHeap.Peek() <= limit {
			result = max(result, minHeap.Len())
			heap.Push(minHeap, nums[right])
			heap.Push(maxHeap, nums[right])
			right++
		} else {
			minHeap.Remove(nums[left])
			maxHeap.Remove(nums[left])
			left++
		}
	}
	if maxHeap.Peek()-minHeap.Peek() <= limit {
		result = max(result, minHeap.Len())
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
