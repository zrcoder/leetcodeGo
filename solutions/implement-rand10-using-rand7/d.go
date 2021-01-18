/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package implement_rand10_using_rand7

import "math/rand"

/*
470. 用 Rand7() 实现 Rand10() https://leetcode-cn.com/problems/implement-rand10-using-rand7
已有方法 rand7 可生成 1 到 7 范围内的均匀随机整数，试写一个方法 rand10 生成 1 到 10 范围内的均匀随机整数。

不要使用系统的 Math.random() 方法。

示例 1:
输入: 1
输出: [7]

示例 2:
输入: 2
输出: [8,4]

示例 3:
输入: 3
输出: [8,1,10]

提示:
rand7 已定义。
传入参数: n 表示 rand10 的调用次数。

进阶:
rand7()调用次数的 期望值 是多少 ?
你能否尽量少调用 rand7() ?
*/
func rand10() int {
	a, b := rand7(), rand7()
	// a 在[1,5]中取值， 概率1/5
	for a > 5 {
		a = rand7()
	}
	// b在[1,6]中取值
	for b == 7 {
		b = rand7()
	}
	// b为偶数和奇数的概率相当，都是1/2；最终概率为1/5 * 1/2 = 1/10
	if b%2 == 0 {
		return a
	}
	return a + 5
}

// 一个假定的rand7实现
func rand7() int {
	return rand.Intn(7) + 1
}
