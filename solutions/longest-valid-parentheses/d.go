package d

/*
32.最长有效括号
https://leetcode-cn.com/problems/longest-valid-parentheses

给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
*/
// 在遍历的时候统计左右括号的数量
func longestValidParentheses(s string) int {
	return max(calFromLowToHight(s), calFromHiToLow(s))
}
func calFromLowToHight(s string) int {
	left, right, res := 0, 0, 0
	for i := range s {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			res = max(res, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}
	return res
}
func calFromHiToLow(s string) int {
	left, right, res := 0, 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			res = max(res, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return res
}

// 借助栈
func longestValidParentheses1(s string) int {
	res := 0
	// 记录索引，维持栈底元素表示最后一个没有被匹配的右括号的索引
	// 初始-1表示在索引-1处有个右括号，帮助后边边界处理
	stack := []int{-1}
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}
		stack = stack[:len(stack)-1] // 当前右括号和栈顶记录的左括号匹配
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			res = max(res, i-stack[len(stack)-1])
		}
	}
	return res
}

// 动态规划
func longestValidParentheses2(s string) int {
	res := 0
	// dp[i]表示以i结尾的子串包含的最大有效括号长度
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == '(' { // 左括号结尾，没有有效括号，最大有效括号长度为0
			continue
		}
		if s[i-1] == '(' {
			dp[i] = 2
			if i >= 2 {
				dp[i] += dp[i-2]
			}
		} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
			dp[i] = dp[i-1] + 2
			if i-dp[i-1] >= 2 {
				dp[i] += dp[i-dp[i-1]-2]
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
