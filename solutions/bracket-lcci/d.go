/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package bracket_lcci

/*
面试题 08.09 括号 https://leetcode-cn.com/problems/bracket-lcci/
括号。设计一种算法，打印n对括号的所有合法的（例如，开闭一一对应）组合。

说明：解集不能包含重复的子集。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/
// 经典回溯
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	curr := make([]byte, 0, 2*n) // 记录当前构建的一个可能串
	left, right := 0, 0          // 记录当前构建的串里左右括号的个数
	var backtrace func()
	// 结果中已经出现left个左括号，right个右括号
	backtrace = func() {
		if left > n || right > left { // 递归结束条件1
			return
		}
		if len(curr) == cap(curr) { // 等价于： left == right == n； 递归结束条件2
			res = append(res, string(curr))
			return
		}
		// 尝试追加一个左括号
		left++
		curr = append(curr, '(')
		backtrace()
		// 回溯
		left--
		curr = curr[:len(curr)-1]

		// 尝试追加一个右括号
		right++
		curr = append(curr, ')')
		backtrace()
		// 回溯
		right--
		curr = curr[:len(curr)-1]
	}
	backtrace()
	return res
}

// 将left、right作为参数，这样可以减少回溯代码。（curr也可以作为参数，后边见）
func generateParenthesis0(n int) []string {
	res := make([]string, 0)
	curr := make([]byte, 0, 2*n) // 记录当前构建的一个可能串
	var backtrace func(int, int)
	// 结果中已经出现left个左括号，right个右括号
	backtrace = func(left, right int) {
		if left > n || right > left {
			return
		}
		if len(curr) == cap(curr) {
			res = append(res, string(curr))
			return
		}
		curr = append(curr, '(')
		backtrace(left+1, right)
		curr = curr[:len(curr)-1]

		curr = append(curr, ')')
		backtrace(left, right+1)
		curr = curr[:len(curr)-1]
	}
	backtrace(0, 0)
	return res
}

// left，right，cur都可以作为参数传递；cur每次复制，可以用string代替[]byte，空间占用稍微多点
func generateParenthesis1(n int) []string {
	res := make([]string, 0)
	var backtrace func(left, right int, curr string)
	backtrace = func(left, right int, curr string) {
		if left > n || right > left {
			return
		}
		if len(curr) == 2*n {
			res = append(res, curr)
			return
		}
		backtrace(left+1, right, curr+"(")
		backtrace(left, right+1, curr+")")
	}
	backtrace(0, 0, "")
	return res
}
