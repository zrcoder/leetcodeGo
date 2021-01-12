package letter_combinations_of_a_phone_number

/*
17. 电话号码的字母组合 https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例:
输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
*/
/*
常规递归回溯
*/
var dic = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	res := []string{}
	if digits == "" {
		return res
	}
	buf := make([]byte, 0, len(digits))
	var helper func(int)
	helper = func(i int) {
		if i == len(digits) {
			res = append(res, string(buf))
			return
		}
		letters := dic[int(digits[i]-'0')-2]
		for j := range letters {
			buf = append(buf, letters[j])
			helper(i + 1)
			buf = buf[:len(buf)-1]
		}
	}
	helper(0)
	return res
}
