/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */
package coin_path

import "math"

/*
给定一个数组 A（下标从 1 开始）包含 N 个整数：A1，A2，……，AN 和一个整数 B。
你可以从数组 A 中的任何一个位置（下标为 i）跳到下标 i+1，i+2，……，i+B 的任意一个可以跳到的位置上。
如果你在下标为 i 的位置上，你需要支付 Ai 个金币。如果 Ai 是 -1，意味着下标为 i 的位置是不可以跳到的。

现在，你希望花费最少的金币从数组 A 的 1 位置跳到 N 位置，你需要输出花费最少的路径，依次输出所有经过的下标（从 1 到 N）。
如果有多种花费最少的方案，输出字典顺序最小的路径。
如果无法到达 N 位置，请返回一个空数组。

样例 1 :
输入: [1,2,4,-1,2], 2
输出: [1,3,5]

样例 2 :
输入: [1,2,4,-1,2], 1
输出: []

注释 :
路径 Pa1，Pa2，……，Pan 是字典序小于 Pb1，Pb2，……，Pbm 的，
当且仅当第一个 Pai 和 Pbi 不同的 i 满足 Pai < Pbi，如果不存在这样的 i 那么满足 n < m。
A1 >= 0。 A2, ..., AN （如果存在） 的范围是 [-1, 100]。
A 数组的长度范围 [1, 1000].
B 的范围 [1, 100].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/coin-path
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
自底向上动态规划
从位置i跳到末尾的花费只与i后面的元素相关
创建长度为n的dp和next数组：dp记录从当前到终点的最小花费；next记录下一跳的位置，方便最终得到路径

正向考虑，比较难的是有多个路径时，取字典序最小的
逆向思考：从后往前计算dp和next
在位置i处，下一跳可能的位置是从i+1到i+B中任意一个。假设下一跳位置为j时，从j跳到终点的花费最小，此时更新next和dp
next[i] = j, dp[i] = A[i] + dp[j]
根据数组 next 找出最小花费的路径方案之一，且是字典序最小的方案！

时间复杂度：O(nB)
空间复杂度：O(n)
*/
func cheapestJump(A []int, B int) []int {
	n := len(A)
	if n == 0 || A[n-1] == -1 {
		return nil
	}
	dp, next := make([]int, n), make([]int, n) // 分别统计从当前位置跳到终点的最小花费，及使得最终花费最小的下一跳的位置
	dp[n-1], next[n-1] = A[n-1], -1            // -1表示下一跳无路可走
	for i := n - 2; i >= 0; i-- {
		dp[i], next[i] = math.MaxInt64, -1 // 先假定无路可走
		if A[i] == -1 {                    // 真的无路可走
			continue
		}
		// 可能有路，在区间[i+1, i+B]中挑一个花费最小的下一跳的位置
		end := min(i+B, n-1)            // i+B可能越界
		for j := i + 1; j <= end; j++ { // 正向遍历；如果求字典序最大的方案，这里可以逆向遍历
			if dp[j] != math.MaxInt64 && A[i]+dp[j] < dp[i] {
				dp[i], next[i] = A[i]+dp[j], j
			}
		}
	}
	if dp[0] == math.MaxInt64 {
		return nil
	}
	var result []int
	for i := 0; i != -1; i = next[i] {
		result = append(result, i+1)
	}
	return result
}

/*
带备忘的回溯，或理解为自顶向下动态规划
在位置 i处，可以跳跃到[i+1,i+B]区间里任意一个位置上。
假设跳到了 j 处，那么从 i 到终点的花费为 A[i] + jump(j)，如果此花费小于当前的最小花费，则更新 next[i]=j
为了减少递归，需要增加备忘录存储已经计算出的结果
时间复杂度：O(nB)
空间复杂度：O(n)
*/
func cheapestJump1(A []int, B int) []int {
	n := len(A)
	if n == 0 || A[n-1] == -1 {
		return nil
	}
	memo, next := make([]int, n), make([]int, n) // 分别统计从当前位置跳到终点的最小花费，及使得最终花费最小的下一跳的位置
	for i := range next {
		next[i] = -1 // 全部置为-1，方便最终求路径时判断
	}
	var jump func(pos int) int
	jump = func(pos int) int {
		if memo[pos] > 0 {
			return memo[pos]
		}
		if pos == n-1 {
			return A[pos]
		}
		minCost := math.MaxInt64
		end := min(pos+B, n-1)
		for i := pos + 1; i <= end; i++ { // 正向遍历；如果求字典序最大的方案，这里可以逆向遍历
			if A[i] < 0 {
				continue
			}
			cost := A[pos] + jump(i)
			if cost < minCost {
				minCost = cost
				next[pos] = i
			}
		}
		memo[pos] = minCost
		return minCost
	}
	jump(0)

	var result []int
	for i := 0; i != -1; i = next[i] {
		result = append(result, i+1)
		if i == n-1 {
			return result
		}
	}
	return nil
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
