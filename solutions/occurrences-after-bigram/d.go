/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package occurrences_after_bigram

/*
https://leetcode-cn.com/problems/occurrences-after-bigram

给出第一个词 first 和第二个词 second，考虑在某些文本 text 中可能以 "first second third" 形式出现的情况，
其中 second 紧随 first 出现，third 紧随 second 出现。

对于每种这样的情况，将第三个词 "third" 添加到答案中，并返回答案。

示例 1：
输入：text = "alice is a good girl she is a good student", first = "a", second = "good"
输出：["girl","student"]

示例 2：
输入：text = "we will we will rock you", first = "we", second = "will"
输出：["we","rock"]
*/
import "strings"

func findOcurrences(text string, first string, second string) []string {
	words := strings.Fields(text)
	var res []string
	for i := 0; i < len(words)-2; i++ {
		if words[i] == first && words[i+1] == second {
			res = append(res, words[i+2])
		}
	}
	return res
}

func findOcurrences1(text string, first string, second string) []string {
	var result []string
	dest := first + " " + second + " "
	index := strings.Index(text, dest)
	for index != -1 {
		newText := text[index+len(dest):]
		if index == 0 || text[index-1] == ' ' {
			result = append(result, getFirstWord(newText))
		}
		text = newText
		index = strings.Index(text, dest)
	}
	return result
}

func getFirstWord(s string) string {
	index := strings.Index(s, " ")
	if index == -1 {
		return s
	}
	return s[:index]
}
