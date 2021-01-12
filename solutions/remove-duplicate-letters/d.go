/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package remove_duplicate_letters

/* 316. 去除重复字母 https://leetcode-cn.com/problems/remove-duplicate-letters/
给定一个仅包含小写字母的字符串，去除字符串中重复的字母，使得每个字母只出现一次。
需保证返回结果的字典序最小（要求不能打乱其他字符的相对位置）。

示例 1:
输入: "bcabc"
输出: "abc"
解释: 原字符串中 ‘b’, ‘c’ 均有多个，在不改变原有字符位置前提下，去掉最前面的 “bc”，使得输出字符串字典序最小

示例 2:
输入: "cbacdcbc"
输出: "acdb"
*/

/*
遍历s， 挑选合适的字母追加到result
对于遍历到的当前字母C，和已经放入result的尾部字母T：

C已经在result里，什么也不做
C不在result里：
	C<T，则看T在原字符串s中C之后是不是还有
		没有, 将C追加到result
		有, 可以从result中删除T，并接着对result尾部字母T'与C做相同的判断处理，最后将C追加到result
		——这里的玩法像是栈！！！
	C>t 追加C

借助两个map：
count首先记录每个字母在s中出现的次数；在修改result时根据情况增减字母个数
inResult记录字母是否已经在result中；在修改result时根据情况标记字母是否在result中
*/
func removeDuplicateLetters(s string) string {
	letterNums := countLetters(s)
	inResult := make(map[rune]bool, 26)
	var result []rune
	for _, c := range s {
		letterNums[c]--
		if inResult[c] {
			continue
		}
		for n := len(result); n > 0 && c < result[n-1] && letterNums[result[n-1]] > 0; n-- {
			last := result[n-1]
			inResult[last] = false
			result = result[:n-1]
		}
		result = append(result, c)
		inResult[c] = true
	}
	return string(result)
}

func countLetters(s string) map[rune]int {
	r := make(map[rune]int, 26)
	for _, c := range s {
		r[c]++
	}
	return r
}

// 因为全是小写字母，两个map的大小最多为26，map可以优化为数组，读写更迅捷~代码略
