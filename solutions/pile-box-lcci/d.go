/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package pile_box_lcci

import "sort"

/*
面试题 08.13. 堆箱子 https://leetcode-cn.com/problems/pile-box-lcci/
堆箱子。给你一堆n个箱子，箱子宽 wi、深 di、高 hi。
箱子不能翻转，将箱子堆起来时，下面箱子的宽度、高度和深度必须大于上面的箱子。
实现一种方法，搭出最高的一堆箱子。箱堆的高度为每个箱子高度的总和。

输入使用数组[wi, di, hi]表示每个箱子。

示例1:

 输入：box = [[1, 1, 1], [2, 2, 2], [3, 3, 3]]
 输出：6
示例2:

 输入：box = [[1, 1, 1], [2, 3, 4], [2, 6, 7], [3, 4, 5]]
 输出：10
提示:

箱子的数目不大于3000个。
*/
/*
类似[354] 俄罗斯套娃信封问题
扩展到了三维
*/
func pileBox(box [][]int) int {
	// 先按照长宽高任意一个维度排序
	sort.Slice(box, func(i, j int) bool {
		bi, bj := box[i], box[j]
		return bi[0] > bj[0]
	})
	dp := make([]int, len(box))
	result := 0
	for i, v := range box {
		dp[i] = box[i][2]
		for j := 0; j < i; j++ {
			bj := box[j]
			if bj[0] > v[0] && bj[1] > v[1] && bj[2] > v[2] {
				dp[i] = max(dp[i], dp[j]+v[2])
			}
		}
		result = max(result, dp[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
