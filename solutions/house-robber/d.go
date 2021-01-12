/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package house_robber

/*
198. 打家劫舍
https://leetcode-cn.com/problems/house-robber

你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

示例 1:

输入: [1,2,3,1]
输出: 4
解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 2:

输入: [2,7,9,3,1]
输出: 12
解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
*/
/*
假设f(i)表示有i+1个房子的时候偷到的最大金额(为了和数组索引对应)
那么f(0) = nums[0]
f(1) = max(nums[0], nums[1]
对于一个大于1的i，可以分为两种情况：
偷i房间，收益尽可能大的话就是nums[i]+f(i-2), 不偷i房间，则f(i)=f(i-1);所以f(i) = max(f(i-1), nums[i]+f(i-2))
这就是一个典型的动态规划
时间复杂度O(n)
空间复杂度O(n), 主要为引入dp数组开辟但额外空间
*/
func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[n-1]
}

// 但实际上只需要两个变量，不需要一个dp数组
// 这样空间复杂度为O(1)
func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	// f(i) = max(f(i-1), nums[i]+f(i-2))
	prev, curr := nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		prev, curr = curr, max(curr, nums[i]+prev)
	}
	return curr
}

//另外实际上我们发现，一开始让prev和curr均为0，从头遍历，逻辑同样正确，代码可以简化
func rob(nums []int) int {
	// f(i) = max(f(i-1), nums[i]+f(i-2))
	prev, curr := 0, 0
	for _, v := range nums {
		prev, curr = curr, max(curr, prev+v)
	}
	return curr
}

/*
213. 打家劫舍 II
https://leetcode-cn.com/problems/house-robber-ii

你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。
这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。
同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

示例 1:

输入: [2,3,2]
输出: 3
解释: 你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
示例 2:

输入: [1,2,3,1]
输出: 4
解释: 你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
	 偷窃到的最高金额 = 1 + 3 = 4 。
*/
func robCycle(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

/*
337. 打家劫舍 III
https://leetcode-cn.com/problems/house-robber-iii

在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。
这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。
一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。

计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。

示例 1:

输入: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \
     3   1

输出: 7
解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.

示例 2:

输入: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \
 1   3   1

输出: 9
解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
一个类似的问题是：968. 监控二叉树，比当前这个问题更难些，要考虑的情况多一种

在树上做动态规划
*/
// 以下写法在最后一个用例超时
func robTree1(root *TreeNode) int {
	var dfs func(node *TreeNode, selected bool) int
	dfs = func(node *TreeNode, selected bool) int {
		if node == nil {
			return 0
		}
		lNotSelected := dfs(node.Left, false)
		rNotSelected := dfs(node.Right, false)
		if selected {
			return node.Val + lNotSelected + rNotSelected
		}
		lSelected := dfs(node.Left, true)
		rSelected := dfs(node.Right, true)
		return max(lSelected, lNotSelected) + max(rSelected, rNotSelected)
	}
	return max(dfs(root, true), dfs(root, false))
}

// 优化以上写法，将多个参数改成多个返回值; 代码更精简，且性能比上边好。
// 需要细想想，一开始的写法有比较多的重复计算，但是后边的写法没有重复计算（也因为这样， 无需加备忘录优化）
func robTree(root *TreeNode) int {
	return max(dfs(root))
}

func dfs(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	lSelected, lNotSelected := dfs(node.Left)
	rSelected, rNotSelected := dfs(node.Right)
	selected := node.Val + lNotSelected + rNotSelected
	notSelected := max(lSelected, lNotSelected) + max(rSelected, rNotSelected)
	return selected, notSelected
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
