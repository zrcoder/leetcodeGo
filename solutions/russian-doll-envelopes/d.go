/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package russian_doll_envelopes

import "sort"

/*
354. 俄罗斯套娃信封问题 https://leetcode-cn.com/problems/russian-doll-envelopes/
给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h) 出现。
当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

说明:
不允许旋转信封。

示例:
输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出: 3
解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
*/
// 同问题[面试题 17.08. 马戏团人塔] 二维lis问题; 时间复杂度O(n*lgn)
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] > envelopes[j][0]
	})
	length := 0
	for _, v := range envelopes {
		j := sort.Search(length, func(i int) bool {
			c := envelopes[i]
			return c[0] <= v[0] || c[1] <= v[1]
		})
		envelopes[j] = v
		if j == length {
			length++
		}
	}
	return length
}

// dp, 时间复杂度O(n^2)
func maxEnvelopes1(envelopes [][]int) int {
	// 先按一个维度排序（宽度或高度都行）；只是把问题降维
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] > envelopes[j][0]
	})
	result := 0
	dp := make([]int, len(envelopes))
	for i, v := range envelopes {
		dp[i] = 1
		for j := 0; j < i; j++ {
			vj := envelopes[j]
			if vj[0] > v[0] && vj[1] > v[1] {
				dp[i] = max(dp[i], dp[j]+1)
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
