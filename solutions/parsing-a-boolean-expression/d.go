package expression

/*
1106. 解析布尔表达式
https://leetcode-cn.com/problems/parsing-a-boolean-expression

给你一个以字符串形式表述的 布尔表达式（boolean） expression，返回该式的运算结果。

有效的表达式需遵循以下约定：

"t"，运算结果为 True
"f"，运算结果为 False
"!(expr)"，运算过程为对内部表达式 expr 进行逻辑 非的运算（NOT）
"&(expr1,expr2,...)"，运算过程为对 2 个或以上内部表达式 expr1, expr2, ... 进行逻辑 与的运算（AND）
"|(expr1,expr2,...)"，运算过程为对 2 个或以上内部表达式 expr1, expr2, ... 进行逻辑 或的运算（OR）


示例 1：

输入：expression = "!(f)"
输出：true
示例 2：

输入：expression = "|(f,t)"
输出：true
示例 3：

输入：expression = "&(t,f)"
输出：false
示例 4：

输入：expression = "|(&(t,f,t),!(t))"
输出：false


提示：

1 <= expression.length <= 20000
expression[i] 由 {'(', ')', '&', '|', '!', 't', 'f', ','} 中的字符组成。
expression 是以上述形式给出的有效表达式，表示一个布尔值。
*/
func parseBoolExpr(expression string) bool {
	var vals, operators []rune
	for _, r := range expression {
		switch r {
		case '|', '&', '!':
			operators = append(operators, r)
		case 't', 'f', '(':
			vals = append(vals, r)
		case ')':
			trues, falses := 0, 0
			for len(vals) > 0 {
				top := vals[len(vals)-1]
				vals = vals[:len(vals)-1]
				if top == 't' {
					trues++
				} else if top == 'f' {
					falses++
				} else { // top == '('
					break
				}
			}
			op := operators[len(operators)-1]
			operators = operators[:len(operators)-1]

			t := 'f'
			if op == '!' && trues == 0 ||
				op == '&' && falses == 0 ||
				op == '|' && trues > 0 {
				t = 't'
			}
			vals = append(vals, t)
		}
	}
	// vals最后仅含一个元素
	return vals[0] == 't'
}
