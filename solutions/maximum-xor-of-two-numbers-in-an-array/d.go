/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package maximum_xor_of_two_numbers_in_an_array

/*
给定一个非空数组，数组中元素为 a0, a1, a2, … , an-1，其中 0 ≤ ai < 231 。

找到 ai 和aj 最大的异或 (XOR) 运算结果，其中0 ≤ i,  j < n 。

你能在O(n)的时间解决这个问题吗？

示例:

输入: [3, 10, 5, 25, 2, 8]

输出: 28

解释: 最大的结果是 5 ^ 25 = 28.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
暴力法, 时间复杂度O(n^2)，空间复杂度O(1)
*/
func findMaximumXOR0(nums []int) int {
	r := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]^nums[j] > r {
				r = nums[i] ^ nums[j]
			}
		}
	}
	return r
}

/*
时空复杂度都是vO(n*32) = O(n)
*/
func findMaximumXOR(nums []int) int {
	result := 0
	mask := 0
	const maxBit uint = 31
	for i := maxBit; i >= 0; i-- {
		mask = mask | (1 << i)
		prefixSet := make(map[int]struct{}, 0)
		for _, v := range nums {
			prefixSet[v&mask] = struct{}{}
		}
		temp := result | (1 << i)
		for k := range prefixSet {
			if _, found := prefixSet[k^temp]; found {
				result = temp
				break
			}
		}
		if i == 0 {
			break
		}
	}
	return result
}
