/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package wiggle_subsequence

/*
376. 摆动序列 https://leetcode-cn.com/problems/wiggle-subsequence/

如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为摆动序列。
第一个差（如果存在的话）可能是正数或负数。少于两个元素的序列也是摆动序列。

例如， [1,7,4,9,2,5] 是一个摆动序列，因为差值 (6,-3,5,-7,3) 是正负交替出现的。
相反, [1,4,7,2,5] 和 [1,7,4,5,5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。

给定一个整数序列，返回作为摆动序列的最长子序列的长度。
通过从原始序列中删除一些（也可以不删除）元素来获得子序列，剩下的元素保持其原始顺序。

示例 1:
输入: [1,7,4,9,2,5]
输出: 6
解释: 整个序列均为摆动序列。

示例 2:
输入: [1,17,5,10,13,15,10,5,16,8]
输出: 7
解释: 这个序列包含几个长度为 7 摆动序列，其中一个可为[1,17,10,13,10,16,8]。

示例 3:
输入: [1,2,3,4,5,6,7,8,9]
输出: 2
进阶:
你能否用 O(n) 时间复杂度完成此题?
*/
/*
动态规划
用两个数组来 dp ，分别记作 up 和 down 。
up[i] 存的是目前为止最长的以第i个元素结尾的上升摆动序列的长度。
类似的， down[i]记录的是目前为止最长的以第 i 个元素结尾的下降摆动序列的长度。

初始状态：up[0] = down[0] = 1
状态转移：
若nums[i] > nums[i-1]:
up[i] = down[i-1]+1; down[i]=down[i-1]
若nums[i] < nums[i-1]:
down[i] = up[i-1]+1; up[i]=up[i-1
若nums[i] == nums[i-1]:
up[i],down[i]与up[i-1],down[i-1]相同

可以看到每次状态只跟上次状态有关，可以用两个变量代替数组

时间复杂度O(n), 空间复杂度O(1)
*/
func wiggleMaxLength1(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	up, down := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}

	return max(up, down)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
贪心
首先，最长摆动序列一定可以以原始数组的第一个数作为开始，证明略，想想容易明白

当序列有一段连续的递增(或递减)时，为了形成摇摆子序列，只需要保留这段连续的递增(或递减)的首尾元素，
这样更可能使得尾部的后一个元素成为摇摆子序列的下一个元素。
*/
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	const begin, up, down = 0, 1, 2
	state := begin
	result := 1
	for i := 1; i < len(nums); i++ {
		switch state {
		case begin:
			if nums[i] > nums[i-1] {
				state = up
				result++
			} else if nums[i] < nums[i-1] {
				state = down
				result++
			}
		case up:
			if nums[i] < nums[i-1] {
				state = down
				result++
			}
		case down:
			if nums[i] > nums[i-1] {
				state = up
				result++
			}
		}
	}
	return result
}
