package ways

import (
	"strconv"
)

/*
91. 解码方法 https://leetcode-cn.com/problems/decode-ways
一条包含字母 A-Z 的消息通过以下方式进行了编码：

'A' -> 1
'B' -> 2
...
'Z' -> 26
给定一个只包含数字的非空字符串，请计算解码方法的总数。

示例 1:
输入: "12"
输出: 2
解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。

示例 2:
输入: "226"
输出: 3
解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
*/
/*
递归解法
这个问题的递归写起来比自底向上的动态规划还要难些
*/
// 从后向前递归
func numDecodings0(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	return help(s)
}

func help(s string) int {
	if len(s) < 2 {
		return 1
	}
	// 从后向前递归，可以处理含有"00"的情况；也可以从前向后
	preLast, last := s[len(s)-2], s[len(s)-1]
	if last == '0' {
		if preLast != '1' && preLast != '2' {
			return 0
		}
		return help(s[:len(s)-2])
	}
	if preLast == '1' || preLast == '2' && last <= '6' {
		return help(s[:len(s)-1]) + help(s[:len(s)-2])
	}
	return help(s[:len(s)-1])
}

/*
动态规划
dp数组可以优化为有限变量
*/
func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	pre, cur := 1, 1
	for i := 1; i < len(s); i++ {
		tmp := cur
		if cannotDecode(i, s) {
			return 0
		}
		if s[i] == '0' {
			cur = pre
		} else if canDecodeTwoWays(i, s) {
			cur += pre
		}
		pre = tmp
	}
	return cur
}

func cannotDecode(i int, s string) bool {
	return s[i-1] != '1' && s[i-1] != '2' && s[i] == '0'
}

func canDecodeTwoWays(i int, s string) bool {
	return s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6')
}

/*
面试题46. 把数字翻译成字符串 https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。
一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

示例 1:
输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"

提示：
0 <= num < 2^31
*/
/*
问题91的简化版本
dp解法
*/
func translateNum1(num int) int {
	if num == 0 {
		return 1
	}
	numStr := strconv.Itoa(num)
	pre, cur := 1, 1
	for i := 1; i < len(numStr); i++ {
		tmp := cur
		if numStr[i-1] == '1' || numStr[i-1] == '2' && numStr[i] < '6' {
			cur += pre
		}
		pre = tmp
	}
	return cur
}

/*
递归解法
*/
func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	lastTwo := num % 100
	if lastTwo > 9 && lastTwo < 26 {
		return translateNum(num/10) + translateNum(num/100)
	}
	return translateNum(num / 10)
}
