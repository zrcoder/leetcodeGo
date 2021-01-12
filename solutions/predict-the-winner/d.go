package predict_the_winner

/*
解法一：
时间复杂度 O(2^n), 空间复杂度 O(n)
*/
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

/*
解法二：
复杂度同上
*/
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

/*
解法一 + 备忘录
时间复杂度会降低到 O(n^2)， 空间复杂度 O(n^2)
代码略
*/
/*
解法三
解法二 + 备忘录
时间复杂度 O(n^2)， 空间复杂度 O(n^2)
*/
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

/*
解法四
动态规划，根据解法二、三，不难想出动态规划的解法，且可以优化dp数组的空间，二维降低到一维
时间复杂度会降低到 O(n^2)， 空间复杂度 O(n)

二维 dp 的代码略
*/
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
