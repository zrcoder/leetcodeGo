/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package shortest_way_to_form_string

/*
1055. 形成字符串的最短路径 https://leetcode-cn.com/problems/shortest-way-to-form-string/
对于任何字符串，我们可以通过删除其中一些字符（也可能不删除）来构造该字符串的子序列。

给定源字符串 source 和目标字符串 target，找出源字符串中能通过串联形成目标字符串的子序列的最小数量。
如果无法通过串联源字符串中的子序列来构造目标字符串，则返回 -1。

示例 1：
输入：source = "abc", target = "abcbc"
输出：2
解释：目标字符串 "abcbc" 可以由 "abc" 和 "bc" 形成，它们都是源字符串 "abc" 的子序列。

示例 2：
输入：source = "abc", target = "acdbc"
输出：-1
解释：由于目标字符串中包含字符 "d"，所以无法由源字符串的子序列构建目标字符串。

示例 3：
输入：source = "xyz", target = "xzyxz"
输出：3
解释：目标字符串可以按如下方式构建： "xz" + "y" + "xz"。

提示：
source 和 target 两个字符串都只包含 "a"-"z" 的英文小写字母。
source 和 target 两个字符串的长度介于 1 和 1000 之间。
*/
/*
主体是遍历target
用s、t两个指针分别指向source、target
只有 source[s] == target[t]时，t才后移一位，而s每次都后移一位，移动到source末尾时，结果+1，s循环到开头
注意遍历时判断target里存在source中没有的字符时，直接返回 -1
*/
func shortestWay(source string, target string) int {
	inSource := make(map[byte]bool, 26)
	for i := range source {
		inSource[source[i]] = true
	}
	result := 0
	s := 0
	t := 0
	for t < len(target) {
		if !inSource[target[t]] {
			return -1
		}
		if source[s] == target[t] {
			t++
		}
		s++
		if s == len(source) {
			result++
			s = 0
		}
	}
	if s != 0 {
		result++
	}
	return result
}
