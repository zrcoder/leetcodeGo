package uibdoe

import "math"

// 动态规划
// 规划过程中有三种状态：0 - 左侧红叶， 1-中间黄叶， 2-右侧红叶
func minimumOperations(leaves string) int {
	n := len(leaves)
	dp0 := boolToInt(leaves[0] == 'y')
	dp1 := math.MaxInt32
	dp2 := math.MaxInt32
	for i := 1; i < n; i++ {
		isRed := leaves[i] == 'r'
		red, yellow := boolToInt(isRed), boolToInt(!isRed)
		tmp := dp1
		dp0, dp1 = dp0+yellow, min(dp0, dp1)+red
		if i > 1 {
			dp2 = min(tmp, dp2) + yellow
		}
	}
	return dp2
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
