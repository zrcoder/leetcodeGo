package valid_parentheses

/*
20. 有效的括号
https://leetcode-cn.com/problems/valid-parentheses

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
*/

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	dic := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, v := range s {
		if left, ok := dic[v]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != left {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack) == 0
}
