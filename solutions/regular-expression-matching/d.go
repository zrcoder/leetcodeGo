/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package regular_expression_matching

import "strings"

/*
10. 正则表达式匹配 https://leetcode-cn.com/problems/regular-expression-matching
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。

示例 1:
输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。

示例 2:
输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。

示例 3:
输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。

示例 4:
输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。

示例 5:
输入:
s = "mississippi"
p = "mis*is*p*."
输出: false
*/
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	firstMatch := len(s) > 0 && (p[0] == s[0] || p[0] == '.')

	if len(p) > 1 && p[1] == '*' {
		return isMatch(s, p[2:]) || firstMatch && isMatch(s[1:], p)
	}
	return firstMatch && isMatch(s[1:], p[1:])
}

/*
动态规划
定义dp[i][j]表示s[:i]和p[:j]是否匹配
*/
func isMatch1(s string, p string) bool {
	dp := initDp1(s, p)
	for i := range s {
		for j := range p {
			judge1(i, j, s, p, dp)
		}
	}
	return dp[len(s)][len(p)]
}

func initDp1(s, p string) [][]bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for j := 1; j < len(p); j++ {
		if p[j] == '*' {
			dp[0][j+1] = dp[0][j-1]
		}
	}
	return dp
}

func judge1(i, j int, s, p string, dp [][]bool) {
	vs, vp := s[i], p[j]
	if vs == vp || vp == '.' {
		dp[i+1][j+1] = dp[i][j]
		return
	}
	if vp != '*' || j == 0 {
		return
	}
	if p[j-1] == s[i] || p[j-1] == '.' {
		dp[i+1][j+1] = dp[i][j+1] || dp[i+1][j-1]
	} else {
		dp[i+1][j+1] = dp[i+1][j-1]
	}
}

/*
44. 通配符匹配 https://leetcode-cn.com/problems/wildcard-matching/
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。

'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
两个字符串完全匹配才算匹配成功。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。

示例 1:
输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。

示例 2:
输入:
s = "aa"
p = "*"
输出: true
解释: '*' 可以匹配任意字符串。

示例 3:
输入:
s = "cb"
p = "?a"
输出: false
解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
示例 4:

输入:
s = "adceb"
p = "*a*b"
输出: true
解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".

示例 5:
输入:
s = "acdcb"
p = "a*c?b"
输入: false
*/
/*
朴素递归，会超时
可以加上备忘录优化，优化代码略
*/
func isMatch2(s string, p string) bool {
	if s == p || p == "*" {
		return true
	}
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(p) == 0 {
		return false
	}
	if len(s) == 0 {
		return strings.Count(p, "*") == len(p)
	}
	if isFirstMatch(p, s) {
		return isMatch(s[1:], p[1:])
	}
	if p[0] == '*' {
		return isMatch(s, p[1:]) || isMatch(s[1:], p)
	}
	return false
}

func isFirstMatch(p, s string) bool {
	return p[0] == '?' || p[0] == s[0]
}

/*
自底向上动态规划
定义dp[i][j]表示s[:i]和p[:j]是否匹配
*/
func isMatch4(s string, p string) bool {
	dp := initDp(s, p)
	for i := range s {
		for j := range p {
			judge4(i, j, s, p, dp)
		}
	}
	return dp[len(s)][len(p)]
}

func initDp(s, p string) [][]bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 0; i < len(p) && p[i] == '*'; i++ {
		dp[0][i+1] = true
	}
	return dp
}

func judge4(i, j int, s, p string, dp [][]bool) {
	vs, vp := s[i], p[j]
	if vs == vp || vp == '?' {
		dp[i+1][j+1] = dp[i][j]
	} else if vp == '*' {
		dp[i+1][j+1] = dp[i][j+1] || dp[i+1][j]
	}
}
