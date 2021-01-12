/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package divide_two_integers

import "math"

/*
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
返回被除数 dividend 除以除数 divisor 得到的商。

示例 1:
输入: dividend = 10, divisor = 3
输出: 3

示例 2:
输入: dividend = 7, divisor = -3
输出: -2

说明:
被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/divide-two-integers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
不能用除法运算符
除法转化为减法：让被除数不断减去除数，直到不够减。
朴素实现会超时
每次减完后除数翻倍，并记录当前为初始除数的几倍（time），若发现不够减且 time 不为 1 则让除数减半， time 也减半
用位移运算符达到翻倍和减半的效果：x += x <=> x *= 2 <=> x <<= 1; x /= 2 <=> x >>= 1
*/
func divide(dividend int, divisor int) int {
	if dividend == 0 || divisor == 1 {
		return dividend
	}
	result := div(abs(dividend), abs(divisor))
	switch {
	case dividend < 0 && divisor > 0 || dividend > 0 && divisor < 0:
		return -result
	default:
		if result > math.MaxInt32 {
			result = math.MaxInt32
		}
		return result
	}
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func div(a, b int) int {
	result, time := 0, 1
	for a >= b || time > 1 {
		if a >= b {
			result += time
			a -= b
			time <<= 1
			b <<= 1
		} else {
			time >>= 1
			b >>= 1
		}
	}
	return result
}
