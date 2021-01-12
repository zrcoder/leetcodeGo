/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package find_repeat_number

/*
找出数组中重复的数字。

在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
容易想到两个方法：
1.先排序，再遍历看相邻元素是否相同——时间复杂度O(nlgn),空间复杂度O(1)
2.借助一个set，遍历数组，如果元素在set里已存在就找到了结果，返回，不存在则存入set——时空复杂度都是O(n);
特别地，在数组的长度不超过 32 的时候，使用位运算的技巧（异或）可以实现 O(1) 空间复杂度判重

以上方法都没有利用题目的一个限定条件：所有数字都在 0～n-1 的范围内
针对这个限定挖掘下：
如果数组是排序的，且没有重复元素，那么数组必然是这样的：
{0, 1, 2, ..., n-2, n-1}
即每个元素的索引和值相等
对于打乱了顺序的数组，可以有这样一个排序方法：
注意到元素的值不会超过数组的长度，元素的值可以作为索引来用
从左到右遍历
对于索引i处，预期nums[i] == i; 如果不相等，可以把nums[i]的值(假设为j)放到索引为j的地方
为了不丢失j处的值，我们可以交换i和j这两处索引的值；
一直检查索引i处的值，直到i == nums[i]，i可以后移一位
可以看到在搞定i处元素的时候，其实同时搞定了其他位置的若干个元素
这样排序, 虽然有两层循环，但对于位置i如果内层循环的次数多，则意味着后边几个已经放好位置的地方不会再有内层循环，
这里用到的是均摊复杂度分析的方法，发生交换操作的次数最多是n，加上遍历数组本身，总的时间复杂度是O(n)

如果有重复元素会怎么样？那就是排序过程中会出现i处的元素nums[i] == j, 想将j放到索引j处时，发现那里已经放好了j，即j是重复元素
*/
func findRepeatNumber(nums []int) int {
	for i, v := range nums {
		for i != v {
			if nums[v] == v {
				return v
			}
			nums[i], nums[v] = nums[v], v
		}
	}
	return -1
}

/*
另一个思路:
遍历数组，每个元素值代表的索引处元素值+n
如果没有重复元素，则操作后所有元素的值都是小于2n的
有重复元素的话，假设重复元素是v，那么索引v处在之前的操作会导致至少加了两次n，其值会大于等于2n
注意一开始加n可能导致索引越界，取模解决
*/
func findRepeatNumber1(nums []int) int {
	n := len(nums)
	for _, v := range nums {
		nums[v%n] += n
	}
	for i, v := range nums {
		if v >= 2*n {
			return i
		}
	}
	return -1
}
