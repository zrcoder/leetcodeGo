package interleaving_string

/*
97. 交错字符串 https://leetcode-cn.com/problems/interleaving-string

给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。

示例 1:

输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出: true
示例 2:

输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
输出: false
*/
// 朴素实现，递归
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if len(s1) == 0 {
		return s2 == s3
	}
	if len(s2) == 0 {
		return s1 == s3
	}
	if s1[0] != s3[0] && s2[0] != s3[0] {
		return false
	}
	if s3[0] == s1[0] && s3[0] == s2[0] {
		return isInterleave(s1[1:], s2, s3[1:]) ||
			isInterleave(s1, s2[1:], s3[1:])
	}
	if s3[0] == s1[0] {
		return isInterleave(s1[1:], s2, s3[1:])
	}
	return isInterleave(s1, s2[1:], s3[1:])
}

/*
可以在朴素实现基础上增加备忘录，形成自顶向下的记忆化搜索解法
看朴素解法，核心是递推的过程；也可以自底向上用动态规划
*/
func isInterleaveDp(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if len(s1) == 0 {
		return s2 == s3
	}
	if len(s2) == 0 {
		return s1 == s3
	}
	n1, n2 := len(s1), len(s2)
	// dp[i][j]代表 s1[:i], s2[:j]是否可交错成字符串s3[:i+j]
	dp := make([][]bool, n1+1)
	for i := range dp {
		dp[i] = make([]bool, n2+1)
	}
	dp[0][0] = true            // 相当于s1、s2、s3都为空
	for i := 1; i <= n1; i++ { // 相当于s2为空
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for i := 1; i <= n2; i++ { // 相当于s1为空
		dp[0][i] = dp[0][i-1] && s2[i-1] == s3[i-1]
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			k := i + j
			dp[i][j] = s1[i-1] == s3[k-1] && dp[i-1][j] ||
				s2[j-1] == s3[k-1] && dp[i][j-1]
		}
	}
	return dp[n1][n2]
}
