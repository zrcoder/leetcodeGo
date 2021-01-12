package multiply_strings

/*
43. 字符串相乘
https://leetcode-cn.com/problems/multiply-strings

给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
*/
/*
假设两个字符串的长度分别是 m、n，则最终结果的长度是 m+n 或 m+n-1， 因为：
假设 num1 和 num2 分别是 m、n 长度的最小值，即 1000...000 (m-1 个 0)和 1000...000 (n-1 个 0)，即 10^(m-1) 和 10(n-1)
乘积为 10^(m+n-2)
假设 num1 和 num2 分别是 m、n 长度的最大值，即 999...999 (m 个)和 999...999 (n 个), 即 10^(m+1)-1 和 10(n+1)-1
乘积为 10^(m+n+2) - 10^m - 10^n + 1

所以结果在闭区间 [10^(m+n-2), 10^(m+n+2) - 10^m - 10^n + 1], 其中又端点 小于10^(m+n+2)
*/
func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	if num1 == "0" || num2 == "0" || m*n == 0 {
		return "0"
	}
	res := make([]int, m+n)
	for i, x := range num1 {
		for j, y := range num2 {
			res[i+j+1] += (int(x) - '0') * int(y-'0')
		}
	}
	for i := m + n - 1; i > 0; i-- {
		res[i-1] += res[i] / 10
		res[i] %= 10
	}

	if res[0] == 0 {
		res = res[1:]
	}
	return parse(res)
}

func parse(src []int) string {
	buf := make([]byte, len(src))
	for i, v := range src {
		buf[i] = byte(v) + '0'
	}
	return string(buf)
}
