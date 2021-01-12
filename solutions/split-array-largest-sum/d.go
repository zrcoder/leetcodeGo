package sum

import (
	"math"
)

/*
410. 分割数组的最大值
链接：https://leetcode-cn.com/problems/split-array-largest-sum


给定一个非负整数数组和一个整数 m，你需要将这个数组分成 m 个非空的连续子数组。
设计一个算法使得这 m 个子数组各自和的最大值最小。

注意:
数组长度 n 满足以下条件:

1 ≤ n ≤ 1000
1 ≤ m ≤ max(50, n)
示例:

输入:
nums = [7,2,5,10,8]
m = 2

输出:
18

解释:
一共有四种方法将nums分割为2个子数组。
其中最好的方式是将其分为[7,2,5] 和 [10,8]，
因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。
*/

// dp 时间复杂度 O(n^2*m), 空间复杂度O(n*m)
func splitArray1(nums []int, m int) int {
	n := len(nums)
	// 前缀和数组，方便后边求nums中任意连续子数组的和
	prefix := make([]int, n+1)
	for i, v := range nums {
		prefix[i+1] = prefix[i] + v
	}
	// dp[i][j] 表示前i个数字划分成j组得到的结果
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		end := max(m, i)
		for j := 1; j <= end; j++ { // i个数字划分为j组
			for k := 0; k < i; k++ {
				dp[i][j] = max(dp[i][j], max(dp[k][j-1], prefix[i]-prefix[k]))
			}

		}
	}
	return dp[n][m]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 二分法 时间复杂度O(n*log(sum-max)), 其中sum和max分别表示nums所有元素和及最大元素
func splitArray(nums []int, m int) int {
	n := len(nums)
	sum, maxNum := 0, 0
	for _, v := range nums {
		sum += v
		if v > maxNum {
			maxNum = v
		}
	}
	if m == 1 {
		return sum
	}
	if m == n {
		return maxNum
	}
	return search(maxNum, sum, m, nums)
}

func search(lo, hi, m int, nums []int) int {
	for lo < hi {
		mid := lo + (hi-lo)/2
		if count(nums, mid) <= m {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func count(nums []int, x int) int {
	// 从头遍历，连续子数组和不超过x划分为一组，计算最终划分了几组
	subSum, count := 0, 1
	for _, v := range nums {
		if subSum+v > x {
			subSum = v
			count++
		} else {
			subSum += v
		}
	}
	return count
}
