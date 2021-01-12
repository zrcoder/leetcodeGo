# [5. 最长回文子串](https://leetcode-cn.com/problems/longest-palindromic-substring/)

## 题目
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
Example:
```
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
```
Example:
```
Input: "cbbd"
Output: "bb"
```
### 1. 朴素实现
时间复杂度O(n^3)，空间复杂度O(1)
```
func longestPalindrome1(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			sub := s[i : j+1]
			if isPalindromic(sub) && len(result) < j-i+1 {
				result = sub
			}
		}
	}
	return result
}

func isPalindromic(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

```
###v2.动态规划改进朴素实现
时间复杂度O(n^2)，空间复杂度O(n)<br>
```text
isPalindromic的复杂度可以改进，用动态规划
定义dp（i, j)表示s[i:j]是否为回文串，有个明显的递推关系：dp(i, j) = s[i]==s[j] && dp(i+1, j-1)
注意到dp(i)依赖于dp(i+1)，可以倒着遍历来确定dp
用一个二维dp数组来记录dp函数的结，并注意当i、j相等、相邻或间隔一个字符(即j-i<3)的初始边界
```
```
func longestPalindrome(s string) string {
	result := ""
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s)-1; i>=0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = s[i] == s[j] && (j-i<3 || dp[i+1][j-1])
			if dp[i][j] && j-i+1 > len(result) {
				result = s[i:j+1]
			}
		}
	}
	return result
}
```
目前时间复杂度O(n^2)，空间复杂度O(n^2)
```text
dp数组可以转为一维，消除i的影响;注意需要倒序遍历j——正序的话会得到错误结果
```
```
func longestPalindrome3(s string) string {
	result := ""
	dp := make([]bool, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		for j := len(s) - 1; j >= 0; j-- {
			dp[j] = s[i] == s[j] && (j-i < 3 || dp[j-1])
			if dp[j] && j-i+1 > len(result) {
				result = s[i : j+1]
			}
		}
	}
	return result
}
```
时间复杂度O(n^2)，空间复杂度O(n)
### 扩展中心
时间复杂度O(n^2)，空间复杂度O(1)
```text
回文串是对称的，可以每次循环选择一个中心，左右扩展，判断左右字符是否相等即可。
需要注意长度为奇数和偶数的情况不同
```
```
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left, right := expandAroundCenter(s, i, i)
		if right-left > end-start {
			start, end = left, right
		}
		left, right = expandAroundCenter(s, i, i+1)
		if right-left > end-start {
			start, end = left, right
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
```
### Manacher 算法
值得一提的是有个Manacher 算法，中国程序员戏称为马拉车算法，可以将时间复杂度降到O(n), 空间复杂度也是O(n)

[参考实现](imp.go)