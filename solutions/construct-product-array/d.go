/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package construct_product_array

/*
给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B 中的元素 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。
示例:
输入: [1,2,3,4,5]
输出: [120,60,40,30,24]
提示：
所有元素乘积之和不会溢出 32 位整数
a.length <= 100000
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
如果允许除法，则非常简单，先遍历数组得到所有所有元素的乘积totalProduct
再遍历一遍，结果数组b[i] = totalProduct / a[i]

不允许除法，可以遍历b，对于b[i],遍历a，得到结果，不过这个复杂度太高，是O(n^2)
实际上，每次计算b[i]，重复计算了，在计算b[i-1]的时候已经计算了b[i]的一部分
从左到右， b[i]可以借助前一次计算结果b[i-1]得到一部分；从右到左，b[i]可以借助前一次计算b[i+1]得到另一部分结果

时间复杂度O(n),空间复杂度O(n)
*/
func constructArr(a []int) []int {
	n := len(a)
	if n == 0 {
		return nil
	}
	b := make([]int, n)
	// 从左到右，确定b中每个元素的一部分
	b[0] = 1
	for i := 1; i < n; i++ {
		b[i] = b[i-1] * a[i-1]
	}
	// 从右向左，确定b中每个元素的另一部分，并与已有的一部分相乘得到最终结果
	rightProduct := a[n-1]
	for i := n - 2; i >= 0; i-- {
		b[i] *= rightProduct
		rightProduct *= a[i]
	}
	return b
}

/*
可以稍作修改，让代码更容易读
也可以引入leftProduct和rightProduct并让初始值为1
进一步，两个变量可以合并为一个
*/
func constructArr1(a []int) []int {
	n := len(a)
	if n == 0 {
		return nil
	}
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = 1
	}
	// 从左到右，确定b中每个元素的一部分
	b[0] = 1
	product := 1
	for i := 0; i < n; i++ {
		b[i] *= product
		product *= a[i]
	}
	// 从右向左，确定b中每个元素的另一部分，并与已有的一部分相乘得到最终结果
	product = 1
	for i := n - 1; i >= 0; i-- {
		b[i] *= product
		product *= a[i]
	}
	return b
}
