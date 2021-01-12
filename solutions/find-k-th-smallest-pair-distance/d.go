/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package find_k_th_smallest_pair_distance

import "sort"

/*
719. 找出第 k 小的距离对 https://leetcode-cn.com/problems/find-k-th-smallest-pair-distance

给定一个整数数组，返回所有数对之间的第 k 个最小距离。
一对 (A, B) 的距离被定义为 A 和 B 之间的绝对差值。

示例 1:
输入：
nums = [1,3,1]
k = 1
输出：0
解释：
所有数对如下：
(1,3) -> 2
(1,1) -> 0
(3,1) -> 2
因此第 1 个最小距离的数对是 (1,1)，它们之间的距离为 0。

提示:
2 <= len(nums) <= 10000.
0 <= nums[i] < 1000000.
1 <= k <= len(nums) * (len(nums) - 1) / 2.
*/

/*
排序后方便计算最小距离，二分搜索
*/
func smallestDistancePair(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	max := nums[len(nums)-1] - nums[0]
	return sort.Search(max+1, func(i int) bool {
		return countLowers(nums, i) >= k
	})
}

// 返回nums中数对距离不大于v的个数；nums已经排序
func countLowers(nums []int, v int) int {
	count := 0
	for left, right := 0, 0; right < len(nums); right++ {
		for nums[right]-nums[left] > v {
			left++
		}
		count += right - left
	}
	return count
}
