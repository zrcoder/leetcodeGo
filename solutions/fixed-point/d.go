/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package fixed_point

import "sort"

/*
1064. 不动点
https://leetcode-cn.com/problems/fixed-point

给定已经按升序排列、由不同整数组成的数组 A，返回满足 A[i] == i 的最小索引 i。如果不存在这样的 i，返回 -1。

示例 1：
输入：[-10,-5,0,3,7]
输出：3
解释：
对于给定的数组，A[0] = -10，A[1] = -5，A[2] = 0，A[3] = 3，因此输出为 3 。

示例 2：
输入：[0,2,5,8,17]
输出：0
A[0] = 0，因此输出为 0 。

示例 3：
输入：[-10,-5,3,4,7,9]
输出：-1
解释：
不存在这样的 i 满足 A[i] = i，因此输出为 -1 。

提示：
1 <= A.length < 10^4
-10^9 <= A[i] <= 10^9
*/

// 朴素解法， 完全没有利用数组已经排序的特点
func fixedPoint1(A []int) int {
	for i, v := range A {
		if i == v {
			return i
		}
	}
	return -1
}

// 一个隐藏的二分搜索问题：如果位置i处元素大于i，那么目标在i左侧；如果i处元素小于i，目标在i右侧；
// 如果等于i，不要立即返回，有可能左侧还有，题目求的是最左侧的
func fixedPoint2(A []int) int {
	left, right := 0, len(A)
	for left < right {
		mid := left + (right-left)/2
		switch {
		case A[mid] >= mid: // 可能左边还有符合题意的
			right = mid
		case A[mid] < mid:
			left = mid + 1
		}
	}
	if left == len(A) || A[left] != left {
		return -1
	}
	return left
}

// 使用标准库，减少代码量
func fixedPoint(A []int) int {
	i := sort.Search(len(A), func(i int) bool {
		return A[i] >= i
	})
	if i == len(A) || A[i] != i {
		return -1
	}
	return i
}
