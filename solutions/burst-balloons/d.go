package balloons

/*
312. 戳气球 https://leetcode-cn.com/problems/burst-balloons/


有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。

现在要求你戳破所有的气球。每当你戳破一个气球 i 时，你可以获得 nums[left] * nums[i] * nums[right] 个硬币。

这里的 left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。

求所能获得硬币的最大数量。

说明:

你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100

示例:

输入: [3,1,5,8]
输出: 167

解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
	 coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
*/
/*
动态规划
定义dp(left, right)表示开区间(left, right)的结果，即不戳破left和right处气球，只戳破中间的气球所得的最多硬币数
那么dp(left, ritht) = max(memo(left, right), memo(left, k) + memo(k, right) + nums[left]*nums[k]*nums[right])
初始状态，left和right相同或相差1，中间没有气球可戳，所以得到的金币数为0
为方便处理，可以事先给nums前后各添加一个1

时间复杂度O(n^3)
空间复杂度O(n^2)
*/
func maxCoins(nums []int) int {
	nums = prepareNums(nums)
	memo := genDpMemo(len(nums))
	return dp(nums, memo)
}

func prepareNums(nums []int) []int {
	n := len(nums) + 2
	help := make([]int, n)
	help[0], help[n-1] = 1, 1
	_ = copy(help[1:], nums)
	return help
}

func genDpMemo(n int) [][]int {
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	return memo
}

func dp(nums []int, memo [][]int) int {
	n := len(nums)
	for left := n - 3; left >= 0; left-- {
		for right := left + 2; right < n; right++ {
			cal(left, right, nums, memo)
		}
	}
	return memo[0][n-1]
}

func cal(left, right int, nums []int, memo [][]int) {
	for k := left + 1; k < right; k++ {
		coins := memo[left][k] + memo[k][right] + nums[left]*nums[k]*nums[right]
		memo[left][right] = max(memo[left][right], coins)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
