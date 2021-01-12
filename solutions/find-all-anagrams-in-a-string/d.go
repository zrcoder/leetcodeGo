/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package find_all_anagrams_in_a_string

/*
242. 有效的字母异位词 https://leetcode-cn.com/problems/valid-anagram
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
*/
/*
1.借助两个哈希表（因只是小写字母，也可用长度26的数组）
分别统计s、t中字符的个数，比较是否相同即可
*/
func isAnagram(s string, t string) bool {
	const letterNums = 26
	const firstLetter = 'a'
	a := [letterNums]int{}
	b := [letterNums]int{}
	for _, v := range s {
		a[v-firstLetter]++
	}
	for _, v := range t {
		b[v-firstLetter]++
	}
	return a == b
}

/*
也可以只借助一个哈希表
先统计s里字符的个数，再遍历t，没一个字符在统计结果里将个数减1
最后遍历一遍统计结果，如果个数出现非0，大于0说明s里有比t多的字符，小于0说明t里有比s多的字符，直接返回false
遍历完毕，即所有字符个数为0，说明s与t是字符异位词
*/
func isAnagram1(s string, t string) bool {
	const letterNums = 26
	const firstLetter = 'a'
	m := [letterNums]int{}
	for _, v := range s {
		m[v-firstLetter]++
	}
	for _, v := range t {
		m[v-firstLetter]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

/*
438. 找到字符串中所有字母异位词 https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。

说明：
字母异位词指字母相同，但排列不同的字符串。
不考虑答案输出的顺序。

示例 1:
输入:
s: "cbaebabacd" p: "abc"
输出:
[0, 6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。

示例 2:
输入:
s: "abab" p: "ab"
输出:
[0, 1, 2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
*/

/*
双指针滑动窗口
*/
func findAnagrams(s string, p string) []int {
	m, n := len(s), len(p)
	if m < n {
		return nil
	}
	needed := make(map[byte]int, 0) // 统计p中字符个数
	for i := 0; i < n; i++ {
		needed[p[i]]++
	}
	found := make(map[byte]int, len(needed))
	matched := 0
	var res []int
	for left, right := 0, 0; right < m; right++ {
		c := s[right]
		if needed[c] > 0 {
			found[c]++
			if found[c] == needed[c] {
				matched++
			}
		}
		for matched == len(needed) {
			if right-left == n-1 {
				res = append(res, left)
			}
			c := s[left]
			if needed[c] > 0 {
				found[c]--
				if found[c] < needed[c] {
					matched--
				}
			}
			left++
		}
	}
	return res
}
