/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package construct_multiple_array

/*
面试题66. 构建乘积数组 https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/
给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，
其中 B 中的元素 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。

不能使用除法。

示例:

输入: [1,2,3,4,5]
输出: [120,60,40,30,24]


提示：
所有元素乘积之和不会溢出 32 位整数
a.length <= 100000
*/

/*
不能使用除法，逼我想其他办法
*/
func constructArr(a []int) []int {
	n := len(a)
	if n == 0 {
		return nil
	}
	result := make([]int, n)
	// 从左到右，确定b中每个元素的一部分
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * a[i-1]
	}
	// 从右向左，确定b中每个元素的另一部分，并与已有的一部分相乘得到最终结果
	rightProduct := a[n-1]
	for i := n - 2; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= a[i]
	}
	return result
}
