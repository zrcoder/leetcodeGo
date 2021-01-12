package palindrome_pairs

/*
336. 回文对
https://leetcode-cn.com/problems/palindrome-pairs

给定一组 互不相同 的单词， 找出所有不同 的索引对(i, j)，使得列表中的两个单词， words[i] + words[j] ，可拼接成回文串。

示例 1：
输入：["abcd","dcba","lls","s","sssll"]
输出：[[0,1],[1,0],[3,2],[2,4]]
解释：可拼接成的回文串为 ["dcbaabcd","abcddcba","slls","llssssll"]

示例 2：
输入：["bat","tab","cat"]
输出：[[0,1],[1,0]]
解释：可拼接成的回文串为 ["battab","tabbat"]
*/
/*
朴素实现复杂度 O(n^2 * m) ，其中 n 为 words 长度， m 为所有单词平均长度
空间复杂度为 O(1)
*/
func palindromePairs0(words []string) [][]int {
	var result [][]int
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if isPalindrome(words[i] + words[j]) {
				result = append(result, []int{i, j})
			}
			if isPalindrome(words[j] + words[i]) {
				result = append(result, []int{j, i})
			}
		}
	}
	return result
}

/*

将一个单词X划分为 left、 right 两部分，
如果其中一部分本身是回文串，且另一部分的逆序列能在给定的 words 数组里找到，假设是单词 Y ，则能将 X 和 Y 拼接成一个回文串

时间复杂度是 O(n * m^2)

暴力解法的复杂度和这个方法的复杂度对比一下就发现：在n远大于m的时候显然这个方法较优，但反过来则暴力法较优~
如果单词不多，且每个单词比较长，用暴力法比较好~
*/
func palindromePairs(words []string) [][]int {
	n := len(words)
	wordsRev := make([]string, n)
	for i, w := range words {
		wordsRev[i] = reverse(w)
	}
	indices := make(map[string]int, n)
	for i, w := range wordsRev {
		indices[w] = i
	}

	var result [][]int
	for i, w := range words {
		m := len(w)
		for j := 0; j <= m; j++ {
			left, right := w[:j], w[j:m]
			if id, ok := indices[left]; ok && id != i && isPalindrome(right) {
				result = append(result, []int{i, id})
			}
			if id, ok := indices[right]; ok && id != i && j != 0 && isPalindrome(left) { // j ==  0 和 j == m的时候有重复
				result = append(result, []int{id, i})
			}
		}
	}
	return result
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func reverse(s string) string {
	n := len(s)
	b := []byte(s)
	for i := 0; i < n/2; i++ {
		b[i], b[n-i-1] = b[n-i-1], b[i]
	}
	return string(b)
}
