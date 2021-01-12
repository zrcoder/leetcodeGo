/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package majority_element

/*
169. 多数元素 https://leetcode-cn.com/problems/majority-element/
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1:
输入: [3,2,3]
输出: 3

示例 2:
输入: [2,2,1,1,1,2,2]
输出: 2
*/
/*
1.可以用一个哈希表统计各个元素出现的次数， 当某个元素的个数超过n/2则得到结果
时空复杂度都是O(n)
2.可以将数组排序，中间元素即为结果
时间复杂度O(nlgn), 空间复杂度O(1)

3. 投票算法
时间复杂度O(n)， 空间复杂度O(1)的解法

如果我们把众数记为 +1，把其他数记为 −1，将它们全部加起来，显然和大于 0，从结果本身我们可以看出众数比其他数多。
可以先假设第一个元素是结果r，遍历数组，和r相等则投票+1，不等则投票-1；当投票数为0，则改选当前元素为结果；一直到遍历完即可找出真正的结果
可以简单理解如下：
如果候选人不是maj 则 maj会和其他非候选人一起反对,所以候选人一定会下台(r==0时发生换届选举)
如果候选人是maj , 则maj 会支持自己，其他候选人会反对，同样因为maj 票数超过一半，所以maj 一定会成功当选
*/
func majorityElement(nums []int) int {
	var r, volts int
	for _, v := range nums {
		if volts == 0 {
			r = v
		}
		if r == v {
			volts++
		} else {
			volts--
		}
	}
	return r
}

/*
由于题目明确说明：给定的数组总是存在多数元素 ，因此上面的解法没有考虑 数组中不存在众数 的情况。
如果题目改变，确实会有不存在众数的情况，该怎么办？
在最后返回前增加验证即可
*/
func majorityElement1(nums []int) int {
	var r, volts int
	for _, v := range nums {
		if volts == 0 {
			r = v
		}
		if r == v {
			volts++
		} else {
			volts--
		}
	}
	// 验证
	count := 0
	for _, v := range nums {
		if r == v {
			count++
		}
	}
	if count > len(nums)/2 {
		return r
	}
	return 0 // 没有众数，返回0
}
