/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package plus_without_arithmetic_operators

/*
面试题65. 不用加减乘除做加法
写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。

示例:
输入: a = 1, b = 1
输出: 2

提示：
a, b 均可能是负数或 0
结果不会溢出 32 位整数
*/
// 位运算，模拟加法器实现
func add(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1 // 分别为不进位相加的结果和进位
	}
	return a
}
