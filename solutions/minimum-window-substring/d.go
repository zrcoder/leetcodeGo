/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package minimum_window_substring

import "math"

/*
76. 最小覆盖子串
https://leetcode-cn.com/·problems/minimum-window-substring

给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。

示例：

输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"
说明：

如果 S 中不存这样的子串，则返回空字符串 ""。
如果 S 中存在这样的子串，我们保证它是唯一的答案。
*/

/*
注意T中可能有重复字符，这意味着结果中该字符的个数至少与T中相等

暴力法：O(n^2)
func minWindow(s string, t string) string {
	for i:=0; i<len(s)-len(t); i++ {
		for j:= i+len(t); i <len(s); j ++ {
			if s[i:j] 包含t中所有字符 { //  用map实现
				更新答案
			}
		}
	}
}

滑动窗口解法：
1、在字符串 S 中使用双指针中的左右指针技巧，初始化 left = right = 0，把索引闭区间 [left, right] 称为一个「窗口」。
2、先不断地增加 right 指针扩大窗口 [left, right]，直到窗口中的字符串符合要求（包含了 T 中的所有字符）。
3、此时，停止增加 right，转而不断增加 left 指针缩小窗口 [left, right]，直到窗口中的字符串不再符合要求（不包含 T 中的所有字符了）。
同时，每次增加 left，我们都要更新一轮结果。
4、重复第 2 和第 3 步，直到 right 到达字符串 S 的尽头。
怎么判断window里是否有t中所有的字符？分别用两个map计数器即可，key为字符，value为字符在串里的个数；
时空复杂度都是(m+n)，m，n分别是s和t的长度
*/
func minWindow(s string, t string) string {
	needs := make(map[byte]int, 0) // 记录t中字符的个数；
	for i := 0; i < len(t); i++ {
		needs[t[i]]++
	}
	found := make(map[byte]int, len(needs)) // 记录滑动窗口中在t中存在的字符及其个数

	matched := 0                   // 记录窗口中有多少个字符符合要求：在t中存在且个数大于等于t中的个数;len(needs) == matched即说明窗口里包含了t中所有字符
	start, end := 0, math.MaxInt32 // 记录结果的起始和结束位置
	// left、right为窗口的左右边界, window = s[left:right+1]
	for left, right := 0, 0; right < len(s); right++ { // 先移动 right 寻找可行解
		c := s[right]
		if needs[c] == 0 {
			continue
		}
		found[c]++
		if found[c] == needs[c] {
			matched++
		}
		for ; matched == len(needs); left++ { // 找到可行解后，开始移动 left 缩小窗口
			if right-left < end-start {
				start, end = left, right
			}
			c := s[left]
			if needs[c] == 0 {
				continue
			}
			found[c]--
			if found[c] < needs[c] {
				matched--
			}
		}
	}
	if end == math.MaxInt32 {
		return ""
	}
	return s[start : end+1]
}
