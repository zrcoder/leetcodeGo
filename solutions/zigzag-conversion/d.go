package zigzag_conversion

import (
	"strings"
)

/*
6. Z 字形变换 https://leetcode-cn.com/problems/zigzag-conversion/
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);

示例 1:
输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"

示例 2:
输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:
L     D     R
E   O E   I I
E C   I H   N
T     S     G
*/
/*
近似模拟：不需要用二维矩阵模拟，只需要确定每一行有哪些字母即可
在遍历s的时候，需要确定当前的行号，附加一个表示方向的变量，容易计算当前行号
时空复杂度都是O(n)， n 是字符串的长度
*/
func convert(s string, numRows int) string {
	if numRows < 1 {
		return ""
	}
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	rows := make([]strings.Builder, numRows)
	curRow, isDown := 0, true
	for i := range s {
		rows[curRow].WriteByte(s[i])
		if isDown && curRow == numRows-1 || !isDown && curRow == 0 {
			isDown = !isDown
		}
		if isDown {
			curRow++
		} else {
			curRow--
		}
	}
	result := strings.Builder{}
	result.Grow(numRows)
	for _, b := range rows {
		result.WriteString(b.String())
	}
	return result.String()
}
