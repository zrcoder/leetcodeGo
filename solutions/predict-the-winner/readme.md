# [486. 预测赢家](https://leetcode-cn.com/problems/predict-the-winner)
```
给定一个表示分数的非负整数数组。 玩家 1 从数组任意一端拿取一个分数，随后玩家 2 继续从剩余数组任意一端拿取分数，然后玩家 1 拿，…… 。
每次一个玩家只能拿取一个分数，分数被拿取之后不再可取。直到没有剩余分数可取时游戏结束。最终获得分数总和最多的玩家获胜。

给定一个表示分数的数组，预测玩家1是否会成为赢家。你可以假设每个玩家的玩法都会使他的分数最大化。


示例 1：
输入：[1, 5, 2]
输出：False
解释：一开始，玩家1可以从1和2中进行选择。
如果他选择 2（或者 1 ），那么玩家 2 可以从 1（或者 2 ）和 5 中进行选择。如果玩家 2 选择了 5 ，那么玩家 1 则只剩下 1（或者 2 ）可选。
所以，玩家 1 的最终分数为 1 + 2 = 3，而玩家 2 为 5 。
因此，玩家 1 永远不会成为赢家，返回 False 。

示例 2：
输入：[1, 5, 233, 7]
输出：True
解释：玩家 1 一开始选择 1 。然后玩家 2 必须从 5 和 7 中进行选择。无论玩家 2 选择了哪个，玩家 1 都可以选择 233 。
     最终，玩家 1（234 分）比玩家 2（12 分）获得更多的分数，所以返回 True，表示玩家 1 可以成为赢家。

提示：

1 <= 给定的数组长度 <= 20.
数组里所有分数都为非负数且不会大于 10000000 。
如果最终两个玩家的分数相等，那么玩家 1 仍为赢家。
```
## 解法一: 模拟递归

有两个玩家，先手玩家1和后手玩家2，需要分别计算他们的总分；

也可以相对先手玩家计算两个玩家的分数总和，但注意后手玩家的分数按照负数计算，这样最终只要看总分是不是 >= 0 即可

每个玩家每次能从左端或右端拿一个分数，之后的分数还是连续序列，

可以记 `f(left, right，isFirstPlayer)` 表示在 `[left, right]` 区间相对玩家1能得到的总分

如果轮到玩家1，即 isFirstPlayer 为 true，那么 

`f(left, right，true) = max(nums[left] + f(left+1, right, false), nums[right] + f(left, right-1, false))`

如果轮到的是玩家2， 那么

`f(left, right, false) = min(-nums[left] + f(left+1, right, true), -nums[right] + f(left, right-1, true))`

时间复杂度 `O(2^n)`,每个元素都有两种尝试； 空间复杂度 `O(n)`，是递归栈的大小
```go
func PredictTheWinner0(nums []int) bool {
	var f func(left, right int, isFirstPlayer bool) int
	f = func(left, right int, isFirstPlayer bool) int {
		if left == right {
			if isFirstPlayer {
				return nums[left]
			}
			return -nums[left]
		}
		if isFirstPlayer {
			return max(nums[left]+f(left+1, right, false), nums[right]+f(left, right-1, false))
		}
		return min(-nums[left]+f(left+1, right, true), -nums[right]+f(left, right-1, true))
	}
	return f(0, len(nums)-1, true) >= 0
}
```
## 解法二：方法一参数简化
也可以去除 isFirstPlayer 参数

对于当前玩家，所能得到的最大分数，应该是将上面的 +f() 改成 -f() 就行

复杂度同上
```go
func PredictTheWinner1(nums []int) bool {
	var f func(left, right int) int
	f = func(left, right int) int {
		if left == right {
			return nums[left]
		}
		return max(nums[left]-f(left+1, right), nums[right]-f(left, right-1))
	}
	return f(0, len(nums)-1) >= 0
}
```

## 解法一 + 备忘录

时间复杂度会降低到 `O(n^2)`， 空间复杂度 `O(n^2)`
代码略

## 解法三：解法二 + 备忘录
时间复杂度 `O(n^2)`， 空间复杂度 `O(n^2)`
```go
func PredictTheWinner2(nums []int) bool {
	n := len(nums)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	var f func(left, right int) int
	f = func(left, right int) int {
		if left == right {
			memo[left][right] = nums[left]
			return nums[left]
		}
		if memo[left][right] > 0 {
			return memo[left][right]
		}
		memo[left][right] = max(nums[left]-f(left+1, right), nums[right]-f(left, right-1))
		return memo[left][right]
	}
	return f(0, len(nums)-1) >= 0
}
```

## 解法四:动态规划
根据解法二、三，不难想出动态规划的解法，且可以优化 dp 数组的空间，二维降低到一维
时间复杂度会降低到 `O(n^2)`， 空间复杂度 `O(n)`

二维 dp 的代码略, 以下为优化到一维的解法
```go
func PredictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([]int, n)
	_ = copy(dp, nums) // dp[i] = nums[i]
	for left := n - 2; left >= 0; left-- {
		for right := left + 1; right < n; right++ {
			dp[right] = max(nums[left]-dp[right], nums[right]-dp[right-1])
		}
	}
	return dp[n-1] >= 0
}
```
[参考具体实现](d.go)