/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package longest_common_prefix

import trie "github.com/zrcoder/leetcodeGo/solutions/implement-trie-prefix-tree"

/*
14. 最长公共前缀
https://leetcode-cn.com/problems/longest-common-prefix

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/

/*1,
自然的想法,将所有字符串排成一列左对齐，从左到右比较即可
时间复杂度O(n*m), m为最短字符串的长度
空间复杂度O(1)
*/
func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return strs[0]
	}
	for i := 0; i < len(strs[0]); i++ {
		ch := strs[0][i]
		for r := 1; r < n; r++ {
			if i == len(strs[r]) || strs[r][i] != ch {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

/*2.
值得一提的是，有一个前缀树的解法(所有字符都在小写字母范围内)：
 时间复杂度和空间复杂度都是O(s), s为所有字符串长度和
*/
func longestCommonPrefix1(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return strs[0]
	}
	tree := trie.Constructor()
	for i := 1; i < len(strs); i++ {
		tree.Insert(strs[i])
	}
	return tree.SearchLongestPrefixOf(strs[0])
}
