/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package valid_mountain_array

/*
941. 有效的山脉数组 https://leetcode-cn.com/problems/valid-mountain-array
给定一个整数数组 A，如果它是有效的山脉数组就返回 true，否则返回 false。

让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：

A.length >= 3
在 0 < i < A.length - 1 条件下，存在 i 使得：
A[0] < A[1] < ... A[i-1] < A[i]
A[i] > A[i+1] > ... > A[B.length - 1]


示例 1：

输入：[2,1]
输出：false
示例 2：

输入：[3,5,5]
输出：false
示例 3：

输入：[0,3,2,1]
输出：true


提示：

0 <= A.length <= 10000
0 <= A[i] <= 10000
*/
func validMountainArray(A []int) bool {
	n := len(A)
	if n < 3 {
		return false
	}
	i := 0
	for i+1 < n && A[i] < A[i+1] {
		i++
	}
	if i == 0 || i == n-1 {
		return false
	}
	for i+1 < n && A[i] > A[i+1] {
		i++
	}
	return i == n-1
}
