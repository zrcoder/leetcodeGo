/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package search_in_a_sorted_array_of_unknown_size

/*
702. 搜索长度未知的有序数组
https://leetcode-cn.com/problems/search-in-a-sorted-array-of-unknown-size

给定一个升序整数数组，写一个函数搜索 nums 中数字 target。如果 target 存在，返回它的下标，否则返回 -1。
注意，这个数组的大小是未知的。你只可以通过 ArrayReader 接口访问这个数组，ArrayReader.get(k) 返回数组中第 k 个元素（下标从 0 开始）。

你可以认为数组中所有的整数都小于 10000。如果你访问数组越界，ArrayReader.get 会返回 2147483647。

样例 1：
输入: array = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 存在在 nums 中，下标为 4
样例 2：
输入: array = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不在数组中所以返回 -1


注释 ：

你可以认为数组中所有元素的值互不相同。
数组元素的值域是 [-9999, 9999]。
*/

/*
因为数组中所有元素的值不同，且元素值域为[-9999, 9999]
则数组中最多有2*9999+1个元素，这样索引的范围就是[0, 2*9999]
现在可以用二分法来寻找target

定义left和right分别为左右边界
实际上left和right的初始值可以大大缩小
一开始分别让left和right取0，1两个值
然后判断reader.get(right), 如果比target大，说明范围需要扩大
这时候可让left=right，为了尽快锁定范围，right可以翻倍
一直循环执行以上两步，直到找到 >= target 的right（此时left是前一个right，正好是right的一半）
*/
const (
	outOfRange = 2147483647
	limited    = 9999
)

func search(reader ArrayReader, target int) int {
	if target > limited || target < -limited {
		return -1
	}
	// 确定大致范围
	left, right := 0, 1
	for reader.get(right) < target && target <= 2*limited {
		left = right
		right *= 2
	}
	// 在确定的范围里二分搜索
	for left <= right {
		mid := left + (right-left)/2
		val := reader.get(mid)
		switch {
		case val == target:
			return mid
		case val < target:
			left = mid + 1
		case val > target, val == outOfRange:
			right = mid - 1
		}
	}
	return -1
}

type ArrayReader []int

func (r ArrayReader) get(i int) int {
	if i < 0 || i >= len(r) {
		return outOfRange
	}
	return r[i]
}
