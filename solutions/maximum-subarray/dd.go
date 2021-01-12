package subarray

func kConcatenationMaxSum(arr []int, k int) int {
	n := len(arr)
	total := min(2, k) * n
	if total == 0 {
		return 0
	}

	dp, res, sum := 0, arr[0], 0
	for i := 0; i < total; i++ {
		v := arr[i%n]
		dp = (max(dp, 0) + v) % 1000000007
		res = max(res, dp)
		if i < len(arr) {
			sum = (sum + v) % 1000000007
		}
	}
	if k <= 2 {
		return res
	}
	// 题目没有描述清楚，实际在结果为负时，需要返回 0
	// return (max(sum, 0)*(k-2) + res) % 1000000007
	return max(0, (max(sum, 0)*(k-2)+res)%1000000007)
}

func kConcatenationMaxSum1(arr []int, k int) int {
	n := len(arr)
	total := min(2, k) * n
	if total == 0 {
		return 0
	}

	dp, res, sum := 0, arr[0], 0
	for i := 0; i < total; i++ {
		v := arr[i%n]
		dp = (dp + v) % 1000000007
		res = max(res, dp)
		if dp < 0 {
			dp = 0
		}
		if i < len(arr) {
			sum = (sum + v) % 1000000007
		}
	}
	if k <= 2 {
		return res
	}
	// 题目没有描述清楚，实际在结果为负时，需要返回 0
	// return (max(sum, 0)*(k-2) + res) % 1000000007
	return max(0, (max(sum, 0)*(k-2)+res)%1000000007)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
