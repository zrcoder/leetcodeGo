/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package add_binary

import "math/big"

/*
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-binary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
方法一：
各个进制的通用解法，只需修改base常量，这个问题中base为2
如： 415. 字符串相加 https://leetcode-cn.com/problems/add-strings/
时空复杂度都是O(max(m, n))， 其中m，n为两个字符串的长度
*/
func addBinary(a string, b string) string {
	m, n := len(a), len(b)
	if m > n {
		return addBinary(b, a)
	}
	const base = 2
	buf := make([]byte, n+1)
	carry := 0
	for i, j := n-1, m-1; i >= 0; i-- {
		if j >= 0 {
			carry += int(a[j] - '0')
			j--
		}
		carry += int(b[i] - '0')
		buf[i+1] = byte(carry%base) + '0'
		carry /= base
	}
	if carry == 0 {
		return string(buf[1:])
	}
	buf[0] = '1'
	return string(buf)
}

/*
方法二：位操作
思路
如果不允许使用加法运算，则可以使用位操作。
XOR 操作得到两个数字无进位相加的结果。
进位和两个数字与操作结果左移一位对应。
两个整数相加可以用如下方法计算：
func add(a, b int) int {
	for b != 0 {
		a, b = a ^ b, (a & b) << 1
	}
	return a
}
现在问题被简化为：首先计算两个数字的无进位相加结果和进位，然后计算无进位相加结果与进位之和。
同理求和问题又可以转换成上一步，直到进位为 0 结束。

算法

把a 和b 转换成整型数字x和y，x保存结果，y保存进位。
当进位不为0即y != 0：
计算当前x 和y 的无进位相加结果：answer = x^y。
计算当前x 和y 的进位：carry = (x & y) << 1。
完成本次循环，更新 x = answer，y = carry。

返回x 的二进制形式。
性能分析:
如果输入数字大于 1 << 100(2^100)必须使用效率较低的 BigInteger。

时间复杂度：O(N+M)，其中N 和M 是输入字符串的长度。
空间复杂度：O(max(N,M))，存储计算结果。

作者：LeetCode
链接：https://leetcode-cn.com/problems/add-binary/solution/er-jin-zhi-qiu-he-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func addBinary1(a string, b string) string {
	const base = 2
	x, y := big.NewInt(0), big.NewInt(0)
	x.SetString(a, base)
	y.SetString(b, base)
	zero := big.NewInt(0)
	for y.Cmp(zero) != 0 {
		answer, carry := big.NewInt(0), big.NewInt(0)
		x, y = answer.Xor(x, y), carry.And(x, y).Lsh(carry, 1)
	}
	return x.Text(base)
}

// 返回a+b的结果，a，b可正可负
func add(a, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}

/*
方法三：
转化成数字再运算;防止越界，用大数；不要用strconv.Atoi和strconv.ParseInt
*/
func addBinary11(a string, b string) string {
	const base = 2
	x, y := big.NewInt(0), big.NewInt(0)
	x.SetString(a, base)
	y.SetString(b, base)
	return x.Add(x, y).Text(base)
}
