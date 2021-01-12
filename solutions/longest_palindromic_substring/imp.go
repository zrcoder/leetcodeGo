/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package longest_palindromic_substring

import "strings"

// 基于中心扩展的马拉车算法
// 参考 https://www.cxyxiaowu.com/2665.html
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	// 先在每个字符中间插入占位符，使得 str 为一个奇数长度的串，方便处理；且 str 不影响s的回文性质
	str := addPlaceholders(s, '#')
	n := len(str)
	p := make([]int, n)
	maxRight, center := 0, 0 // 分别记录当前扩展最右边界及其对应的回文子串的中心位置

	result := ""
	for i := range str {
		if i < maxRight { // 马拉车算法的核心：在这个分支里，利用了p已有记录，减少了以i为中心的扩展次数
			//   --- maxRight1 --- i1 --- center --- i --- maxRight ---
			i1 := 2*center - i            // i1 是 i 关于 center 对称位置
			p[i] = min(maxRight-i, p[i1]) // maxRight-i = i1-maxRight1
		}
		// 一步一步向两端扩展
		left, right := i-p[i]-1, i+p[i]+1
		for left >= 0 && right < n && str[left] == str[right] {
			p[i]++
			left--
			right++
		}
		if i+p[i] > maxRight { // 发现了新的扩展右边界，需要更新 maxRight 和 center
			maxRight = i + p[i]
			center = i
		}

		if len(result) < p[i] { // 更新结果
			result = s[(i-p[i])/2 : (i+p[i])/2]
		}

	}
	return result
}

func addPlaceholders(s string, c byte) string {
	buf := strings.Builder{}
	buf.WriteByte(c)
	for _, v := range s {
		buf.WriteRune(v)
		buf.WriteByte(c)
	}
	return buf.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
