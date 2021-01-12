/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package missing_element_in_sorted_array

import "sort"

/*
1060. 有序数组中的缺失元素 https://leetcode-cn.com/problems/missing-element-in-sorted-array
给出一个有序数组 A，数组中的每个数字都是 独一无二的，找出从数组最左边开始的第 K 个缺失数字。

示例 1：
输入：A = [4,7,9,10], K = 1
输出：5
解释：
第一个缺失数字为 5 。

示例 2：
输入：A = [4,7,9,10], K = 3
输出：8
解释：
缺失数字有 [5,6,8,...]，因此第三个缺失数字为 8 。

示例 3：
输入：A = [1,2,4], K = 3
输出：6
解释：
缺失数字有 [3,5,6,7,...]，因此第三个缺失数字为 6 。


提示：
1 <= A.length <= 50000
1 <= A[i] <= 1e7
1 <= K <= 1e8
*/

// 计算每个元素之前缺失几个元素，再做计算；注意k超出最后一个元素前缺失元素个数的情况
func missingElement1(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	missCount := make([]int, n) //1. 记录第i个元素之前缺少的元素个数；missCount[0] = 0
	for i := 1; i < n; i++ {
		// 递推公式， 由这个递推公式可以推出通项公式，详见后续countMiss函数
		missCount[i] = nums[i] - nums[i-1] - 1 + missCount[i-1]
	}
	for i, v := range missCount {
		if v >= k {
			return nums[i-1] + k - missCount[i-1]
		}
	}
	return nums[n-1] + k - missCount[n-1]
}

// 1.处的数组可以优化为函数
func missingElement2(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	lastMissCount := countMiss(n-1, nums)
	if lastMissCount < k {
		return nums[n-1] + k - lastMissCount //2.
	}
	i := 1 // 找到合适的i，使得countMiss(i-1) < k <= countMiss(i)
	for countMiss(i, nums) < k {
		i++
	}
	return nums[i-1] + k - countMiss(i-1, nums) //3.
}

// 2.3.处的形式其实类似，可以归一
func missingElement3(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	i := 1 // 找到合适的i，使得countMiss(i-1) < k <= countMiss(i)
	for i < n && countMiss(i, nums) < k {
		i++
	}
	return nums[i-1] + k - countMiss(i-1, nums) //3.
}

// 寻找合适的i可以用二分法; 不过要注意k超出最后一个元素前缺失元素个数的情况
func missingElement(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	lastMissCount := countMiss(n-1, nums)
	if lastMissCount < k {
		return nums[n-1] + k - lastMissCount
	}
	i := sort.Search(n-1, func(i int) bool {
		return countMiss(i, nums) >= k
	})
	return nums[i-1] + k - countMiss(i-1, nums)
}

// index处元素之前缺失的元素个数
func countMiss(index int, nums []int) int {
	// 递推公式：	countMiss(i) = nums[i] - nums[i-1] - 1 + countMiss(i-1)
	// =>
	// 通项公式：	countMiss(i) = nums[i] - nums[0] - i
	return nums[index] - nums[0] - index
}
