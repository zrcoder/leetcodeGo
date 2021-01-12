/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package longest_substring_without_repeating_characters

import "math"

/*
3. 无重复字符的最长子串 https://leetcode-cn.com/problems/longest-substring-without-repeating-characters

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
/*
滑动窗口
时间复杂度：O(2n)=O(n)，在最糟糕的情况下，每个字符将被i 和j 访问两次。
空间复杂度：O(k)，Set 的大小。取决于字符串的大小n 以及字符集 / 字母的大小m 。
*/
func lengthOfLongestSubstring1(s string) int {
	result := 0
	set := make(map[byte]struct{}, 0)
	for l, r := 0, 0; r < len(s); {
		c := s[r]
		if _, found := set[c]; !found {
			set[c] = struct{}{}
			r++
			if r-l > result {
				result = r - l
			}
		} else {
			delete(set, s[l])
			l++
		}
	}
	return result
}

/*
优化的滑动窗口
上述的方法最多需要执行 2n 个步骤。事实上，它可以被进一步优化为仅需要 n 个步骤。
可以定义字符到索引的映射，而不是使用集合来判断一个字符是否存在。 当找到重复的字符时，可以立即跳过该窗口。
*/
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int, 0)
	result := 0
	for left, right := 0, 0; right < len(s); right++ {
		c := s[right]
		if index, found := m[c]; found {
			left = max(left, index+1)
		}
		result = max(result, right-left+1)
		m[c] = right
	}
	return result
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

/*
另一个修改版，map里记录字符的个数，通用的滑动窗口解法
*/
func lengthOfLongestSubstring10(s string) int {
	found := make(map[byte]int, 0)
	result := 0
	for l, r := 0, 0; r < len(s); {
		c := s[r]
		found[c]++
		r++
		for found[c] > 1 {
			found[s[l]]--
			l++
		}
		if r-l > result {
			result = r - l
		}
	}
	return result
}
