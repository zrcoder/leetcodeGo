package palindrome

import "strings"

/*
https://leetcode-cn.com/problems/shortest-palindrome

给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。

示例 1:

输入: "aacecaaa"
输出: "aaacecaaa"
示例 2:

输入: "abcd"
输出: "dcbabcd"
*/

// 找到最长的回文前缀， 剩余的部分逆序后加到原字符串前边即可

// 朴素实现，复杂度O(n^2), 略

/*
马拉车方法求最长回文前缀

时空复杂度都是 O(n)
*/
func shortestPalindrome1(s string) string {
	t := insertSeparators(s)
	index := -1
	center, end := -1, -1
	armLens := make([]int, len(t))
	for i := range t {
		if i <= end {
			i1 := 2*center - i
			minArmLen := min(armLens[i1], end-i)
			armLens[i] = expand(t, i-minArmLen, i+minArmLen)
		} else {
			armLens[i] = expand(t, i, i)
		}
		if i+armLens[i] > end {
			center = i
			end = i + armLens[i]
		}
		if i-armLens[i] == 0 {
			index = i
		}
	}
	b := []byte(s[index:])
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - 1 - i
		b[i], b[j] = b[j], b[i]
	}
	return string(b) + s
}

func insertSeparators(s string) string {
	buf := strings.Builder{}
	buf.Grow(len(s)*2 + 1)
	buf.WriteRune('#')
	for _, v := range s {
		buf.WriteRune(v)
		buf.WriteRune('#')
	}
	return buf.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func expand(s string, i, j int) int {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return (j - i - 1) / 2
}

/*
Robin-Karp 编码方法求最长回文前缀, 如 1392.最长快乐前缀

时间复杂度 O(n), 空间复杂度O(1)
*/
func shortestPalindrome(s string) string {
	prefixCode, reverseCode := 0, 0
	base, mod, mul := 131, int(1e9+7), 1
	index := -1
	for i := range s {
		prefixCode = (prefixCode*base%mod + int(s[i])) % mod
		reverseCode = (int(s[i])*mul%mod + reverseCode) % mod
		mul = mul * base % mod
		if prefixCode == reverseCode {
			index = i
		}
	}
	b := []byte(s[index+1:])
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - 1 - i
		b[i], b[j] = b[j], b[i]
	}
	return string(b) + s
}

// kmp算法，略
