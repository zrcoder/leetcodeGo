package decode_string

/*
394. 字符串解码 https://leetcode-cn.com/problems/decode-string/
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例 1：
输入：s = "3[a]2[bc]"
输出："aaabcbc"

示例 2：
输入：s = "3[a2[c]]"
输出："accaccacc"

示例 3：
输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"

示例 4：
输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"
*/
/*
遍历s中字符，不断追加到结果数组，直到遇到 ']',
这时候要从结果数组向前查找 '['
然后统计左右方括号间的字符以及方括号前边的数字，结果数组收缩到数字前边
根据数字和方括号中字符集形成重复串，追加到结果数组
*/
func decodeString(s string) string {
	var result []byte
	for i := range s {
		if s[i] != ']' || result == nil {
			result = append(result, s[i])
			continue
		}
		index := len(result) - 1
		for index >= 0 && result[index] != '[' {
			index--
		}
		sub := copyBytes(result[index+1:])
		result = result[:index]
		index--
		for index >= 0 && isNum(result[index]) {
			index--
		}
		num := 0
		for j := index + 1; j < len(result); j++ {
			num = num*10 + int(result[j]-'0')
		}
		result = result[:index+1]
		result = appendSubs(num, sub, result)
	}
	return string(result)
}

func copyBytes(b []byte) []byte {
	r := make([]byte, len(b))
	_ = copy(r, b)
	return r
}

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func appendSubs(repeat int, sub, result []byte) []byte {
	for repeat > 0 {
		for _, v := range sub {
			result = append(result, v)
		}
		repeat--
	}
	return result
}
