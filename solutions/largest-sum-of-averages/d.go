/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package largest_sum_of_averages

import "math"

/*
813. 最大平均值和的分组 https://leetcode-cn.com/problems/largest-sum-of-averages/

我们将给定的数组 A 分成 K 个相邻的非空子数组 ，我们的分数由每个子数组内的平均值的总和构成。
计算我们所能得到的最大分数是多少。
注意我们必须使用 A 数组中的每一个数进行分组，并且分数不一定需要是整数。

示例:
输入:
A = [9,1,2,3,9]
K = 3
输出: 20
解释:
A 的最优分组是[9], [1, 2, 3], [9]. 得到的分数是 9 + (1 + 2 + 3) / 3 + 9 = 20.
我们也可以把 A 分成[9, 1], [2], [3, 9].
这样的分组得到的分数为 5 + 2 + 6 = 13, 但不是最大值.

说明:
1 <= A.length <= 100.
1 <= A[i] <= 10000.
1 <= K <= A.length.
答案误差在 10^-6 内被视为是正确的。
*/

/*
动态规划
如过k==1，或者数组长度为1问题就会变简单；
这样可以用k和数组长度两个状态，从（1,1）递推得到结果；即动态规划
设dp(i, k)表示将前i个元素划分成k个分组得到的最大分数
则对于j(k-1 <= j <= i-1)
dp(i, k) = max(dp(i, k), dp(j, k-1) + average(A[j+1:i+1]))
初始状态：
dp(i, 1) = average(A[:i])
为迅速计算average(A[j+1:i+1])，可以事先计算出A的前缀和数组

时空复杂度都是O(n*K)，其中n为数组大小
*/
func largestSumOfAverages(A []int, K int) float64 {
	n := len(A)
	if n == 0 {
		return 0
	}
	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, K+1)
	}
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + A[i-1]
		dp[i][1] = float64(prefixSum[i]) / float64(i)
	}
	for i := 1; i <= n; i++ {
		for k := 2; k <= K; k++ {
			calDp(i, k, dp, prefixSum)
		}
	}
	return dp[n][K]
}

func calDp(i, k int, dp [][]float64, prefixSum []int) {
	for j := k - 1; j < i; j++ {
		dp[i][k] = math.Max(dp[i][k], dp[j][k-1]+average(prefixSum, i, j))
	}
}

func average(prefixNum []int, i, j int) float64 {
	return float64(prefixNum[i]-prefixNum[j]) / float64(i-j)
}
