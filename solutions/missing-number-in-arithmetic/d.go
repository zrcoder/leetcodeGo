/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package missing_number_in_arithmetic

/*
1228. 等差数列中缺失的数字 https://leetcode-cn.com/problems/missing-number-in-arithmetic-progression
有一个数组，其中的值符合等差数列的数值规律，也就是说：

在 0 <= i < arr.length - 1 的前提下，arr[i+1] - arr[i] 的值都相等。
我们会从该数组中删除一个 既不是第一个 也 不是最后一个的值，得到一个新的数组  arr。

给你这个缺值的数组 arr，请你帮忙找出被删除的那个数。

示例 1：

输入：arr = [5,7,11,13]
输出：9
解释：原来的数组是 [5,7,9,11,13]。

示例 2：
输入：arr = [15,13,12]
输出：14
解释：原来的数组是 [15,14,13,12]。

提示：
3 <= arr.length <= 1000
0 <= arr[i] <= 10^5
*/
func missingNumber(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	n := len(arr)
	d := (arr[n-1] - arr[0]) / n
	for i := 0; i < n-1; i++ {
		if arr[i]+d != arr[i+1] {
			return arr[i] + d
		}
	}
	return 0
}
