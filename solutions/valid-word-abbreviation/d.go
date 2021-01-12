/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package valid_word_abbreviation

/*
408. 有效单词缩写 https://leetcode-cn.com/problems/valid-word-abbreviation
给一个 非空 字符串 s 和一个单词缩写 abbr ，判断这个缩写是否可以是给定单词的缩写。

字符串 "word" 的所有有效缩写为：

["word", "1ord", "w1rd", "wo1d", "wor1", "2rd", "w2d", "wo2", "1o1d", "1or1", "w1r1", "1o2", "2r1", "3d", "w3", "4"]
注意单词 "word" 的所有有效缩写仅包含以上这些。任何其他的字符串都不是 "word" 的有效缩写。

注意:
假设字符串 s 仅包含小写字母且 abbr 只包含小写字母和数字。

示例 1:

给定 s = "internationalization", abbr = "i12iz4n":

函数返回 true.


示例 2:

给定 s = "apple", abbr = "a2e":

函数返回 false.
*/
func validWordAbbreviation(word string, abbr string) bool {
	i, num := 0, 0
	for _, b := range []byte(abbr) {
		if b >= '0' && b <= '9' {
			if b == '0' && num == 0 {
				return false
			}
			num = num*10 + int(b-'0')
		} else {
			i += num
			if i >= len(word) || b != word[i] {
				return false
			}
			i, num = i+1, 0
		}
	}
	return i+num == len(word)
}

func validWordAbbreviation1(word string, abbr string) bool {
	wordIndex, abbrIndex := 0, 0
	for wordIndex < len(word) && abbrIndex < len(abbr) {
		if word[wordIndex] == abbr[abbrIndex] {
			wordIndex, abbrIndex = wordIndex+1, abbrIndex+1
		} else if abbr[abbrIndex] >= '0' && abbr[abbrIndex] <= '9' {
			if abbr[abbrIndex] == '0' {
				return false
			}
			num := 0
			for abbrIndex < len(abbr) && abbr[abbrIndex] >= '0' && abbr[abbrIndex] <= '9' {
				num = num*10 + int(abbr[abbrIndex]-'0')
				abbrIndex++
			}
			wordIndex += num
		} else {
			return false
		}
	}
	return wordIndex == len(word) && abbrIndex == len(abbr)
}
