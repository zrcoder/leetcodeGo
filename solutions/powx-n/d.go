/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package powx_n

/*
50. Pow(x, n)
https://leetcode-cn.com/problems/powx-n

实现 pow(x, n) ，即计算 x 的 n 次幂函数。

示例 1:

输入: 2.00000, 10
输出: 1024.00000
示例 2:

输入: 2.10000, 3
输出: 9.26100
示例 3:

输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25
说明:

-100.0 < x < 100.0
n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。
*/

func myPow(x float64, n int) float64 {
	// 先将n<0的情况转换为n为正数的情况
	if n < 0 {
		return pow(1/x, -n)
	}
	return pow(x, n)
}

/*
二分递归
*/
// n >= 0
func pow1(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	halfPow := pow1(x, n/2)
	if n%2 == 1 {
		return halfPow * halfPow * x
	}
	return halfPow * halfPow
}

/*
二分迭代
来自于上边递归的思想

可以这样理解：
从计算机的角度看， n 就是0和1组成的串： 00011010100...
所以
x^n  = x^(2^ix + 2^iy+...+2^iz ) = x^(2^ix)* x( 2^iy)*...*x(2^iz )

func pow(x float64, n int) float64 {
	res := 1.0
	for mask := 1; mask <= n; mask <<= 1 {
		if n&mask != 0 {
			res *= x
		}
		x *= x
	}
	return res
}

甚至可以去掉 mask 变量，直接将n折半来循环
*/
// n >= 0
func pow(x float64, n int) float64 {
	res := 1.0
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
	}
	return res
}
