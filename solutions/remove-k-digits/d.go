/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package remove_k_digits

import "strings"

/*
402. 移掉K位数字 https://leetcode-cn.com/problems/remove-k-digits/

给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。

注意:
num 的长度小于 10002 且 ≥ k。
num 不会包含任何前导零。

示例 1 :
输入: num = "1432219", k = 3
输出: "1219"
解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。

示例 2 :
输入: num = "10200", k = 1
输出: "200"
解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。

示例 3 :
输入: num = "10", k = 2
输出: "0"
解释: 从原数字移除所有的数字，剩余为空就是0。
*/

/*
贪心
对于一个数字序列 [D1D2D3…Dn]，如果数字 D2小于其左邻居D1，则应该删除左邻居（D1）以获得最小结果。
如果D1<=D2<=...<=Di, 且Di>Di+1则可以把Di及其前边的数字删除（在k的限度内），以使最终结果最小
通过一个接一个的删除数字，逐步的接近最优解。
如果遍历完成后，删除的数字个数m小于k，则直接把最后的k-m的数字删除即可
*/
func removeKdigits(num string, k int) string {
	var result []byte
	for i := range num {
		// result 在 k 的限度内维持单调非递减
		for isLastBiggerThanCur(k, result, num[i]) {
			result = result[:len(result)-1]
			k--
		}
		result = append(result, num[i])
	}
	result = result[:len(result)-k] // 末尾的数字删除
	s := string(result)
	s = strings.TrimLeft(s, "0") // 处理“00000xxx”这样的情况
	if s == "" {
		return "0"
	}
	return s
}

func isLastBiggerThanCur(k int, result []byte, c byte) bool {
	return k > 0 && len(result) > 0 && result[len(result)-1] > c
}
