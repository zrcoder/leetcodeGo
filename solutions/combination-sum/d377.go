package combination_sum

// 递归回溯，超时
func combinationSum4Timeout(nums []int, target int) int {
	var res int
	var backtrack func(t int)
	backtrack = func(t int) {
		if t == 0 {
			res++
		}
		if t < 0 {
			return
		}
		for _, v := range nums {
			backtrack(t - v)
		}
	}
	backtrack(target)
	return res
}

// 加上备忘录来优化
func combinationSum4Memo(nums []int, target int) int {
	memo := make(map[int]int, 0)
	var backtrack func(t int) int
	backtrack = func(t int) int {
		if t == 0 {
			return 1
		}
		if t < 0 {
			return -1
		}
		if v, ok := memo[t]; ok {
			return v
		}
		res := 0
		for _, v := range nums {
			if backtrack(t-v) != -1 {
				res += backtrack(t - v)
			}
		}
		memo[t] = res
		return res
	}
	res := backtrack(target)
	if res == -1 {
		return 0
	}
	return res
}

// 根据以上记忆化解法，可以得到动态规划解法
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for t := 1; t <= target; t++ {
		for _, v := range nums {
			if v <= t {
				dp[t] += dp[t-v]
			}
		}
	}
	return dp[target]
}
