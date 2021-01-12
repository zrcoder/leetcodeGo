/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package number_of_longest_increasing_subsequence

/*
给定一个未排序的整数数组，找到最长递增子序列的个数。

示例 1:

输入: [1,3,5,4,7]
输出: 2
解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
示例 2:

输入: [2,2,2,2,2]
输出: 5
解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
注意: 给定的数组长度不超过 2000 并且结果一定是32位有符号整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
动态规划，时间复杂度O(n^2),空间复杂度O(n)
*/
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	lengths := make([]int, n) // lengths[i] 代表 nums[:i+1] 最长递增子序列的长度
	counts := make([]int, n)  // counts[i] 代表 nums[:i+1] 最长递增子序列的个数
	longest := 0
	for i := 0; i < n; i++ {
		counts[i], lengths[i] = 1, 1
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}
			if lengths[j] >= lengths[i] {
				counts[i] = counts[j]
			} else if lengths[j]+1 == lengths[i] {
				counts[i] += counts[j]
			}
			lengths[i] = lengths[j] + 1
		}
		if lengths[i] > longest {
			longest = lengths[i]
		}
	}
	result := 0
	for i, v := range lengths {
		if v == longest {
			result += counts[i]
		}
	}
	return result
}
