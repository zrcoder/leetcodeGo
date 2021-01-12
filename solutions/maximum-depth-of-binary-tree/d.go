/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package maximum_depth_of_binary_tree

import "math"

/*
104. 二叉树的最大深度 https://leetcode-cn.com/problems/maximum-depth-of-binary-tree

给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// 时间复杂度O(n): 每个节点遍历一次 —— n为节点数量
// 空间复杂度，当树完全不平衡，退化为链表，最坏为O(n)；当树平衡时，为O(lgn)
func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(maxDepth(node.Left), maxDepth(node.Right))
}

/*
变体： 如果求二叉树的最大直径呢？

543. 二叉树的直径 https://leetcode-cn.com/problems/diameter-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

直径所覆盖的箭头上总节点数 = 左子树最大深度 + 右子树最大深度 + 1（node本身）
直径需要-1，故直径 = leftDepth + rightDepth
*/
func diameterOfBinaryTree(node *TreeNode) int {
	if node == nil {
		return 0
	}
	sum := maxDepth(node.Left) + maxDepth(node.Right)
	sum = max(diameterOfBinaryTree(node.Left), sum)
	sum = max(diameterOfBinaryTree(node.Right), sum)
	return sum
}

// 可以进一步优化，在求最大深度的过程中更新结果
func diameterOfBinaryTree1(root *TreeNode) int {
	result := 0
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int { // 计算node的最大深度
		if node == nil {
			return 0
		}
		leftDepth, rightDepth := depth(node.Left), depth(node.Right)
		result = max(result, leftDepth+rightDepth)
		return 1 + max(leftDepth, rightDepth)
	}
	_ = depth(root)
	return result
}

/*
变体：
687. 最长同值路径 https://leetcode-cn.com/problems/longest-univalue-path

给定一个二叉树，找到最长的路径，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。

注意：两个节点之间的路径长度由它们之间的边数表示。

示例 1:

输入:

              5
             / \
            4   5
           / \   \
          1   1   5
输出:
2
示例 2:

输入:

              1
             / \
            4   5
           / \   \
          4   4   5
输出:
2

注意: 给定的二叉树不超过10000个结点。 树的高度不超过1000。
*/

/*
对于一个节点node， 以其为根节点的路径就是分别向左下和右下延伸形成的箭头

设计一个递归函数，返回当前节点的最大同值路径
如果该节点的值与左右子树的值都不相等，则其最大同值路径为0
如果该节点的值与左右子树的值都相等，则其最大同值路径是其左子树最大同值路径+1 + 右子树最大同值路径+1
如果该节点的值等于其左子树的值但不等于右子树的值，则最长同值路径为左子树的最长同值路径+1
如果该节点的值等于其右子树的值，则最长同值路径为右子树的最长同值路径+1

我们用一个全局变量记录这个最大值，不断更新
*/
func longestUnivaluePath(root *TreeNode) int {
	result := 0

	var calculate func(node *TreeNode) int
	calculate = func(node *TreeNode) int { // 返回包含 node 的最大同值路径
		if node == nil {
			return 0
		}
		left, right := calculate(node.Left), calculate(node.Right)
		if node.Left != nil && node.Left.Val == node.Val {
			left++
		} else {
			left = 0
		}
		if node.Right != nil && node.Right.Val == node.Val {
			right++
		} else {
			right = 0
		}
		result = max(result, left+right)
		return max(left, right)
	}

	_ = calculate(root)
	return result
}

/*
变体
124. 二叉树中的最大路径和 https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
给定一个非空二叉树，返回其最大路径和。

本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。

示例 1:

输入: [1,2,3]

       1
      / \
     2   3

输出: 6
示例 2:

输入: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

输出: 42
*/
func maxPathSum(root *TreeNode) int {
	result := math.MinInt32
	var help func(node *TreeNode) int
	help = func(node *TreeNode) int { // 返回包含node的最大路径和
		if node == nil {
			return 0
		}
		left, right := help(node.Left), help(node.Right)
		left, right = max(0, left), max(0, right)
		result = max(result, node.Val+left+right)
		return node.Val + max(left, right)
	}
	_ = help(root)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
