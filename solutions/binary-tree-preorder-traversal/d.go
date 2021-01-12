/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package binary_tree_preorder_traversal

/*
144. 二叉树的前序遍历
https://leetcode-cn.com/problems/binary-tree-preorder-traversal

给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]

进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归实现，时空复杂度O(n), n为节点总数
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return res
}

// 迭代实现，时空复杂度O(n), n为节点总数
func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}

/* 迭代，节点标记法
在出入栈的时候，标记节点，具体为：
标记节点的状态，新节点为false，已使用（在这道题里是指将节点值追加到结果数组）的节点true。
如果遇到未标记的节点，则将其标记为true，然后将其右节点、左节点、自身依次入栈；注意到顺序与遍历次序正好相反。
如果遇到的节点标记为true，则使用该节点。

这个方法在前序、中序、后续遍历里的实现代码总体逻辑一致，只是入栈的顺序稍微调整即可
*/
func preorderTraversal2(root *TreeNode) []int {
	var res []int
	stack := []*TreeNode{root}
	marked := make(map[*TreeNode]bool, 0)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		if marked[node] {
			res = append(res, node.Val)
			continue
		}
		marked[node] = true
		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
		stack = append(stack, node)
	}
	return res
}

/*
morris 迭代实现
优秀的是空间复杂度
详见 94. 二叉树的中序遍历 问题 morris 迭代实现解法说明
*/
var res []int

func preorderTraversalMorris(root *TreeNode) []int {
	res = nil
	cur := root
	var node *TreeNode
	for cur != nil {
		if cur.Left == nil {
			res = append(res, cur.Val)
			cur = cur.Right
			continue
		}
		// 找 cur 的前驱
		node = cur.Left
		for node.Right != nil && node.Right != cur {
			node = node.Right
		}
		if node.Right == nil { // 还没线索化，建立线索
			res = append(res, cur.Val)
			node.Right = cur
			cur = cur.Left
		} else { // 已经线索化，访问节点并删除线索以恢复树的结构
			node.Right = nil
			cur = cur.Right
		}
	}
	return res
}
