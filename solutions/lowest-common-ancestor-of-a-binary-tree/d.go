/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package lowest_common_ancestor_of_a_binary_tree

import "container/list"

/*
236. 二叉树的最近公共祖先 https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree

给定一个二叉树, 找到该树中两个指定结点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，
满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个结点也可以是它自己的祖先）。”

例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]

示例 1:
输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 结点 5 和结点 1 的最近公共祖先是结点 3。

示例 2:
输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 结点 5 和结点 4 的最近公共祖先是结点 5。因为根据定义最近公共祖先结点可以为结点本身。

说明:
所有结点的值都是唯一的。
p、q 为不同结点且均存在于给定的二叉树中。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
递归
时空复杂度均为O(n), n为树中所有结点个数
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

/*
类似并查集的思想
在遍历树的过程中，用一个哈希表记录每个结点的父指针，如果p、q都遍历了，可以不再遍历其他的结点
另用一个哈希表，存储p的所有祖先；最后在其中找q的祖先，找到的第一个就是结果
遍历可以用递归也可以用栈迭代
时空复杂度均为O(n)
*/
// 并查集思路迭代版
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	stack := list.New()
	stack.PushBack(root)
	parent := map[*TreeNode]*TreeNode{}
	for stack.Len() > 0 && (parent[p] == nil || parent[q] == nil) {
		node := stack.Remove(stack.Back()).(*TreeNode)
		if node.Right != nil {
			parent[node.Right] = node
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			parent[node.Left] = node
			stack.PushBack(node.Left)
		}
	}
	ancestors := map[*TreeNode]bool{}
	for p != nil {
		ancestors[p] = true
		p = parent[p]
	}
	for !ancestors[q] {
		q = parent[q]
	}
	return q
}

// 并查集思路递归版
func lowestCommonAncestor11(root, p, q *TreeNode) *TreeNode {
	parent := map[*TreeNode]*TreeNode{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || parent[p] != nil && parent[q] != nil {
			return
		}
		if node.Right != nil {
			parent[node.Right] = node
		}
		if node.Left != nil {
			parent[node.Left] = node
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	ancestors := map[*TreeNode]bool{}
	for p != nil {
		ancestors[p] = true
		p = parent[p]
	}
	for !ancestors[q] {
		q = parent[q]
	}
	return q
}
