/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package excel_sheet_column_title

import "bytes"

/*
168. Excel表列名称 https://leetcode-cn.com/problems/excel-sheet-column-title
给定一个正整数，返回它在 Excel 表中相对应的列名称。

例如，

    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB
    ...

示例 1:
输入: 1
输出: "A"

示例 2:
输入: 28
输出: "AB"

示例 3:
输入: 701
输出: "ZY"
*/

/*
十进制转二十六进制
需要注意不是0-25，而是1-26
*/
func convertToTitle(n int) string {
	const base = 26
	var buf []byte
	for n > 0 {
		r := n % base
		if r == 0 {
			r = base
			n--
		}
		buf = append([]byte{'A' + byte(r-1)}, buf...)
		n /= base
	}
	return string(buf)
}

func convertToTitle1(n int) string {
	const base = 26
	var buf []byte
	for n > 0 {
		r := n % base
		if r == 0 {
			r = base
			n--
		}
		buf = append(buf, 'A'+byte(r-1))
		n /= base
	}
	res := bytes.Buffer{}
	for i := len(buf) - 1; i >= 0; i-- {
		res.WriteByte(buf[i])
	}
	return res.String()
}
