package garden

import (
	"math"
)

/*
1326. 灌溉花园的最少水龙头数目
https://leetcode-cn.com/problems/minimum-number-of-taps-to-open-to-water-a-garden

在 x 轴上有一个一维的花园。花园长度为 n，从点 0 开始，到点 n 结束。

花园里总共有 n + 1 个水龙头，分别位于 [0, 1, ..., n] 。

给你一个整数 n 和一个长度为 n + 1 的整数数组 ranges ，
其中 ranges[i] （下标从 0 开始）表示：如果打开点 i 处的水龙头，可以灌溉的区域为 [i -  ranges[i], i + ranges[i]] 。

请你返回可以灌溉整个花园的 最少水龙头数目 。如果花园始终存在无法灌溉到的地方，请你返回 -1 。


示例 1：
输入：n = 5, ranges = [3,4,1,1,0,0]
输出：1
解释：
点 0 处的水龙头可以灌溉区间 [-3,3]
点 1 处的水龙头可以灌溉区间 [-3,5]
点 2 处的水龙头可以灌溉区间 [1,3]
点 3 处的水龙头可以灌溉区间 [2,4]
点 4 处的水龙头可以灌溉区间 [4,4]
点 5 处的水龙头可以灌溉区间 [5,5]
只需要打开点 1 处的水龙头即可灌溉整个花园 [0,5] 。
示例 2：
输入：n = 3, ranges = [0,0,0,0]
输出：-1
解释：即使打开所有水龙头，你也无法灌溉整个花园。
示例 3：
输入：n = 7, ranges = [1,2,1,0,2,1,0,1]
输出：3
示例 4：
输入：n = 8, ranges = [4,0,0,0,0,0,0,0,4]
输出：2
示例 5：
输入：n = 8, ranges = [4,0,0,0,4,0,0,0,4]
输出：1

提示：
1 <= n <= 10^4
ranges.length == n + 1
0 <= ranges[i] <= 100
*/
func minTaps(n int, ranges []int) int {
	// prev[i]表示以i为右端点的喷射区间最左的左端点
	prev := make([]int, n+1)
	for i := range prev {
		prev[i] = math.MaxInt32
	}
	for i, v := range ranges {
		l, r := max(0, i-v), min(n, i+v)
		prev[r] = min(prev[r], l)
	}

	dp := make([]int, n+1)
	// dp[0] == 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
		for j := prev[i]; j <= i; j++ {
			if dp[j] != math.MaxInt32 {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	if dp[n] == math.MaxInt32 {
		return -1
	}
	return dp[n]
}

/*
贪心
用 breakpoint 表示最近一个被选取的区间的左边界，初始值为 n，表示选取了无效的 [n, n] 区间；

用 furthest 表示遍历到当前位置 i 时，所有 prev(i) 的最小值，
也就是当必须选择一个区间时，应选择的那个区间的左端点。初始值为一个很大的整数；

从 n 开始递减遍历花园中的每一个位置：

对于当前位置 i，用 prev(i) 更新 furthest；
如果当前位置 i 等于左边界 breakpoint，说明遍历到了当前选择的所有区间的交集的左边界，
此时必须要选择一个新的区间，它的左端点是 furthest。如果 furthest >= breakpoint，则无解；
如果 furthest < breakpoint，将 breakpoint 更新为 furthest，并将结果（选择的区间数量）增加 1。
*/
func minTaps1(n int, ranges []int) int {
	// prev[i]表示以i为右端点的喷射区间最左的左端点
	prev := make([]int, n+1)
	for i := range prev {
		prev[i] = math.MaxInt32
	}
	for i, v := range ranges {
		l, r := max(0, i-v), min(n, i+v)
		prev[r] = min(prev[r], l)
	}

	breakPoint, furthhest, result := n, math.MaxInt32, 0
	for i := n; i > 0; i-- {
		furthhest = min(furthhest, prev[i])
		if i == breakPoint {
			if furthhest >= i {
				return -1
			}
			breakPoint = furthhest
			result++
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
