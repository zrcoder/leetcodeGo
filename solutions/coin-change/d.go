package change

/*
322. 零钱兑换  https://leetcode-cn.com/problems/coin-change/

给定不同面额的硬币 coins 和一个总金额 amount。
编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

示例 1:

输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
示例 2:

输入: coins = [2], amount = 3
输出: -1

说明:
你可以认为每种硬币的数量是无限的。
*/
/*
不求max而求min的背包问题, 参考[用背包问题思想来理解硬币找零系列问题]
(https://leetcode-cn.com/problems/coin-change/solution/yong-bei-bao-wen-ti-si-xiang-lai-li-jie-ying-bi-zh/)
动态规划，用背包的思路来解决这个问题
*/
func coinChange(coins []int, amount int) int {
	maxAmount := amount + 1
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = maxAmount
	}
	dp[0] = 0
	for _, v := range coins {
		for j := v; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-v]+1)
		}
	}
	if dp[amount] == maxAmount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
518. 零钱兑换 II https://leetcode-cn.com/problems/coin-change-2/
给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。

示例 1:
输入: amount = 5, coins = [1, 2, 5]
输出: 4
解释: 有四种方式可以凑成总金额:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1

示例 2:
输入: amount = 3, coins = [2]
输出: 0
解释: 只用面额2的硬币不能凑成总金额3。

示例 3:
输入: amount = 10, coins = [10]
输出: 1

注意:
你可以假设：

0 <= amount (总金额) <= 5000
1 <= coin (硬币面额) <= 5000
硬币种类不超过 500 种
结果符合 32 位符号整数
*/
/*
动态规划，用背包的思路来解决这个问题
*/
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, v := range coins {
		for j := v; j <= amount; j++ {
			dp[j] += dp[j-v]
		}
	}
	return dp[amount]
}
