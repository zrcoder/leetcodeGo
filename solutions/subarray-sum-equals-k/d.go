/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package subarray_sum_equals_k

/*
560. 和为K的子数组
给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

示例 1 :

输入:nums = [1,1,1], k = 2
输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
说明 :

数组的长度为 [1, 20,000]。
数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
*/
/*
O(n^2)的暴力解
*/
func subarraySum0(nums []int, k int) int {
	res := 0
	for i := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				res++
			}
		}
	}
	return res
}

/*
前缀和技巧
sum(nums[i:j+1]) == k
<=> prefixSum(j)-prefixSum(i-1) == k
<=> prefixSum(i-1) = prefixSum(j) - k
*/
func subarraySum(nums []int, k int) int {
	res, prefixSum := 0, 0
	counts := map[int]int{}
	counts[0] = 1
	for _, v := range nums {
		prefixSum += v
		res += counts[prefixSum-k]
		counts[prefixSum]++
	}
	return res
}
