/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package subset

/*
78. 子集
https://leetcode-cn.com/problems/subsets

给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

/*
朴素实现
时空复杂度均为O(n*2^n)
*/
func subsets(nums []int) [][]int {
	result := [][]int{{}} // 空集也是子集之一
	for _, num := range nums {
		for _, r := range result {
			tmp := make([]int, len(r)+1)
			_ = copy(tmp, r)
			tmp[len(tmp)-1] = num
			result = append(result, tmp)
		}
	}
	return result
}

/*
回溯1
参考 491. 递增子序列 解法二，使用递归枚举子序列的通用模板
*/
func subsets1(nums []int) [][]int {
	var result [][]int
	var curr []int
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(nums) {
			result = append(result, copySlice(curr))
			return
		}
		// 不选择当前元素
		backtrack(i + 1)
		// 选择当前元素
		curr = append(curr, nums[i])
		backtrack(i + 1)
		curr = curr[:len(curr)-1]
	}
	backtrack(0)
	return result
}

// 回溯2
func subsets2(nums []int) [][]int {
	var result [][]int
	var curr []int
	var backtrack func(start int)
	backtrack = func(start int) {
		result = append(result, copySlice(curr))
		for i := start; i < len(nums); i++ {
			curr = append(curr, nums[i])
			backtrack(i + 1)
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(0)
	return result
}

func copySlice(s []int) []int {
	r := make([]int, len(s))
	_ = copy(r, s)
	return r
}

/*
二进制枚举

nums 里的每个元素，要么在结果中，要么不在结果中
用一个 n 位的 bitset 来表示各个元素在不在结果中，
如 000...000 表示所有元素都不在结果中，000..011 表示后边两个元素在结果中

局限：len(nums)不能大于64， 否则无法用一个int做mask

时空复杂度均为O(n*2^n)
*/
func subsets3(nums []int) [][]int {
	var result [][]int
	end := 1 << len(nums)
	for mask := 0; mask < end; mask++ {
		var cur []int
		for i, v := range nums {
			if (1<<i)&mask != 0 {
				cur = append(cur, v)
			}
		}
		result = append(result, cur)
	}
	return result
}
