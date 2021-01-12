package longest_happy_prefix

/*
1392. 最长快乐前缀
https://leetcode-cn.com/problems/longest-happy-prefix

「快乐前缀」是在原字符串中既是 非空 前缀也是后缀（不包括原字符串自身）的字符串。
给你一个字符串 s，请你返回它的 最长快乐前缀。
如果不存在满足题意的前缀，则返回一个空字符串。

示例 1：
输入：s = "level"
输出："l"
解释：不包括 s 自己，一共有 4 个前缀（"l", "le", "lev", "leve"）和 4 个后缀（"l", "el", "vel", "evel"）。最长的既是前缀也是后缀的字符串是 "l" 。
示例 2：
输入：s = "ababab"
输出："abab"
解释："abab" 是最长的既是前缀也是后缀的字符串。题目允许前后缀在原字符串中重叠。
示例 3：
输入：s = "leetcodeleet"
输出："leet"
示例 4：
输入：s = "a"
输出：""

提示：
1 <= s.length <= 10^5
s 只含有小写英文字母
*/
/*
朴素实现

时间复杂度 O(n^2)
*/
func longestPrefix0(s string) string {
	res := -1
	for i := 0; i < len(s)-1; i++ {
		prefix := s[:i+1]
		j := len(s) - 1 - i
		suffix := s[j:]
		if prefix == suffix { // 这个比较是 O(i)的复杂度
			res = i
		}
	}
	if res == -1 {
		return ""
	}
	return s[:res+1]
}

/*
Rabin-Karp 编码

Rabin-Karp 字符串编码是一种将字符串映射成整数的编码方式，可以看成是一种哈希算法。
具体地，假设字符串包含的字符种类不超过 ∣Σ∣（其中 Σ 表示字符集），那么我们选一个大于等于 ∣Σ∣ 的整数 base，
就可以将字符串看成 base进制 的整数，将其转换成十进制数后，就得到了字符串对应的编码。

时间复杂度 O(n)
*/
func longestPrefix(s string) string {
	const (
		base = 29
		mod  = int(1e9 + 7)
	)
	prefixVal, suffixVal := 0, 0
	times := 1
	res := -1
	for i := 0; i < len(s)-1; i++ {
		prefixVal = (prefixVal*base + int(s[i]-'a')) % mod
		j := len(s) - 1 - i
		suffixVal = (suffixVal + int(s[j]-'a')*times) % mod
		times = times * base % mod
		if prefixVal == suffixVal {
			res = i
		}
	}
	if res == -1 {
		return ""
	}
	return s[:res+1]
}
