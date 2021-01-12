/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package top_k_frequent_elements

/*
347. 前 K 个高频元素 https://leetcode-cn.com/problems/top-k-frequent-elements

给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
示例 1:
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]

示例 2:
输入: nums = [1], k = 1
输出: [1]

说明：
你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
*/
/*
首先用一个哈希表统计各元素的个数，key为元素值，value为个数
接下来怎么利用这个统计结果？
注意到统计结果中元素的个数是有边界的：[0, maxCount]，且可以肯定maxCount不大于n
可以再用一个哈希表，以各个元素的个数为key，
因可能有不同的元素个数相同，可以将个数相同的元素组织成一个列表，新哈希表的值就是一个列表
最后遍历取出前k个元素即可
时间复杂度O(n), 空间复杂度O(n)
*/
func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	count := make(map[int]int, 0)
	for _, v := range nums {
		count[v]++
	}
	maxCount := 0
	m := make(map[int][]int, 0)
	for num, c := range count {
		m[c] = append(m[c], num)
		if c > maxCount {
			maxCount = c
		}
	}
	result := make([]int, k)
	i := 0
	for j := maxCount; j > 0; j-- {
		for _, v := range m[j] { // m[j]可能并不存在，不用特意判断，Go语法兼容range一个nil切片
			result[i] = v
			i++
			if i == k {
				return result
			}
		}
	}
	return result
}
