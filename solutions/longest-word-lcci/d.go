/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package longest_word_lcci

import "sort"

/*
面试题 17.15. 最长单词
给定一组单词words，编写一个程序，找出其中的最长单词，且该单词由这组单词中的其他单词组合而成。
若有多个长度相同的结果，返回其中字典序最小的一项，若没有符合要求的单词则返回空字符串。

示例：

输入： ["cat","banana","dog","nana","walk","walker","dogwalker"]
输出： "dogwalker"
解释： "dogwalker"可由"dog"和"walker"组成。

提示：
0 <= len(words) <= 100
1 <= len(words[i]) <= 100
*/

/* 自顶向下递归
1.（同上）
2. 将words排序：长度降序，长度相同的单词按照字典序
3. 遍历words， 将当前单词word在集合里删除，再在集合剩余的元素里查看word的子串是否存在
*/
func longestWord(words []string) string {
	set := make(map[string]bool, len(words))
	for _, s := range words {
		set[s] = true
	}
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return words[i] < words[j]
		}
		return len(words[i]) > len(words[j])
	})
	for _, word := range words {
		delete(set, word)
		if findAllSubs(set, word) {
			return word
		}
	}
	return ""
}

func findAllSubs(set map[string]bool, word string) bool {
	if len(word) == 0 {
		return true
	}
	for i := 1; i <= len(word); i++ {
		if !set[word[:i]] {
			continue
		}
		if findAllSubs(set, word[i:]) {
			return true
		}
	}
	return false
}
