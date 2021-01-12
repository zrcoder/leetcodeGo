/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package closed_number_lcci

/*
面试题 05.04. 下一个数
下一个数。给定一个正整数，找出与其二进制表达式中1的个数相同且大小最接近的那两个数（一个略大，一个略小）。

示例1:
 输入：num = 2（或者0b10）
 输出：[4, 1] 或者（[0b100, 0b1]）

示例2:
 输入：num = 1
 输出：[2, -1]
提示:
num的范围在[1, 2147483647]之间；
如果找不到前一个或者后一个满足条件的正数，那么输出 -1。
*/

func findClosedNumbers(num int) []int {
	find := 0
	ans := []int{-1, -1}
	sn := 0
	on := uint(0) // sn中1的个数
	for num > 0 {
		pn := num & (-num)
		num &= num - 1
		// 1<<30判断是否超出范围
		if pn&(1<<30) == 0 && num&(pn<<1) == 0 && ans[0] == -1 {
			ans[0] = num | (pn << 1) | (1<<on - 1)
			find++
		}
		// pn<=1无略小值
		if pn > 1 && sn&(pn>>1) == 0 && ans[1] == -1 {
			ans[1] = num | (pn >> 1) | ((pn>>1 - 1) ^ (pn>>1-1)>>on)
			find++
		}
		if find == 2 {
			break
		}
		sn |= pn
		on++
	}
	return ans
}
