/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package path_sum

import (
	"bytes"
	"sort"
	"strconv"
)

/*
112. 路径总和 https://leetcode-cn.com/problems/path-sum

给定一个二叉树和一个目标和，判断该树中是否存在根结点到叶子结点的路径，这条路径上所有结点值相加等于目标和。
说明: 叶子结点是指没有子结点的结点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根结点到叶子结点的路径 5->4->11->2。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 常规dfs
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

/* 变体
113. 路径总和 II https://leetcode-cn.com/problems/path-sum-ii

给定一个二叉树和一个目标和，找到所有从根结点到叶子结点路径总和等于给定目标和的路径。
说明: 叶子结点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:
[
   [5,4,11,2],
   [5,8,4,5]
]
*/
/*
用一个切片path记录遍历的路径，到达叶子节点发现path内元素和为sum则将当期path添加到结果里，
注意切片底层是同一个数组，添加到结果时要深拷贝一份
*/
func pathSum(root *TreeNode, sum int) [][]int {
	var result [][]int
	var path []int
	prefixSum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		prefixSum += node.Val
		if node.Left == nil && node.Right == nil && prefixSum == sum {
			tmp := make([]int, len(path))
			_ = copy(tmp, path)
			result = append(result, tmp)
		}
		dfs(node.Left)
		dfs(node.Right)
		path = path[:len(path)-1]
		prefixSum -= node.Val
	}
	dfs(root)
	return result
}

/*变体，类似前缀树Trie的实现，用数组表示树，且树是多叉树，应该怎么解？
假设有k个节点，每个节点从0到k-1编号，编号即为其id
caps数组表示每个节点的值
哈希表relations，是个邻接表，键为节点id，值为节点的孩子节点组成的数组
给定sum，返回每条从根节点（id为0）出发到叶子节点，值相加和为sum的路径组成的集合
路径处理成字符串，前最终结果按照字符串非递增排序
*/
func getPath(caps []int, relations map[int][]int, sum int) []string {
	var result []string
	var path []int
	prefixSum := 0
	var dfs func(nodeId int)

	dfs = func(nodeId int) {
		path = append(path, caps[nodeId])
		prefixSum += caps[nodeId]
		if len(relations[nodeId]) == 0 && prefixSum == sum {
			result = append(result, parsePath(path))
		}
		for _, c := range relations[nodeId] {
			dfs(c)
		}
		path = path[:len(path)-1]
		prefixSum -= caps[nodeId]
	}
	dfs(0)

	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})
	return result
}
func parsePath(path []int) string {
	buf := bytes.NewBuffer(nil)
	for _, v := range path {
		buf.WriteString(strconv.Itoa(v))
		buf.WriteString(" ")
	}
	result := buf.String()
	return result[:len(result)-1]
}

/* 变体 假设不一定要从根节点开始，也不需要走到叶子节点，来查找和为定值的路径呢？
437. 路径总和 III https://leetcode-cn.com/problems/path-sum-iii

给定一个二叉树，它的每个结点都存放着一个整数值。
找出路径和等于给定数值的路径总数。
路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。

示例：
root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

返回 3。和等于 8 的路径有:

1.  5 -> 3
2.  5 -> 2 -> 1
3.  -3 -> 11
*/
// 递归解法，时间复杂度较高，有比较多的重复计算
func pathSumCount(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	result := countPrefix(root, sum)
	result += pathSumCount(root.Left, sum)
	result += pathSumCount(root.Right, sum)
	return result
}

// 返回前缀和为sum的路径个数， 递归版
func countPrefix(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	result := 0
	if root.Val == sum {
		result = 1
	}
	result += countPrefix(root.Left, sum-root.Val)
	result += countPrefix(root.Right, sum-root.Val)
	return result
}

/*
countPrefix的另一个实现，用一个变量prefixSum记录当前路径前缀和，到达叶子节点后会回溯
这个思路引出pathSumCount0的实现
*/
func countPrefix1(root *TreeNode, sum int) int {
	prefixSum := 0
	result := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		prefixSum += node.Val
		if prefixSum == sum {
			result++
		}
		dfs(node.Left)
		dfs(node.Right)
		// 回溯
		prefixSum -= node.Val
	}
	dfs(root)
	return result
}

/*
参考 560. 和为K的子数组问题前缀和技巧
如果某个节点 x 的前缀和等于其某个子孙节点 y 的前缀和减去sum，
即 prefixSum(x) = prefixSum(y)-sum ，说明 x 到 y 这条路径的和是 sum
借助一个哈希表记录每条路径上，每个前缀和出现的次数，减少重复计算
*/
func pathSumCount0(root *TreeNode, sum int) int {
	counts := make(map[int]int, 0) // 记录前缀和，key为前缀和，value为前缀和的个数
	counts[0] = 1                  // 前缀和为0的一条路径，方便边界处理，即节点值就是sum这种情况
	res := 0
	prefixSum := 0

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		prefixSum += node.Val        // 当前节点node的前缀和（即从root到当前节点这条路径的和）
		res += counts[prefixSum-sum] // 如果当前节点之前已经有前缀和为 prefixSum-sum 的节点，说明那些节点到当前节点的和就是sum
		counts[prefixSum]++
		dfs(node.Left)
		dfs(node.Right)
		// 回溯
		counts[prefixSum]--
		prefixSum -= node.Val
	}

	dfs(root)
	return res
}
