/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package binary_tree_level_order_traversal

import "container/list"

/*
102. 二叉树的层序遍历 https://leetcode-cn.com/problems/binary-tree-level-order-traversal

给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
有两种通用的遍历树的策略：

深度优先搜索 depth first search（DFS）
在这个策略中，我们采用深度作为优先级，以便从根开始一直到达某个确定的叶子，然后再返回根到达另一个分支。
深度优先搜索策略又可以根据根节点、左孩子和右孩子的相对顺序被细分为先序遍历，中序遍历和后序遍历。

广度优先搜索breadth first search（BFS）
我们按照高度顺序一层一层的访问整棵树，高层次的节点将会比低层次的节点先被访问到。

本问题就是用广度优先搜索遍历二叉树。
*/

// 递归
func levelOrder(root *TreeNode) [][]int {
	var (
		result [][]int
		help   func(node *TreeNode, level int)
	)
	help = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(result) == level {
			result = append(result, make([]int, 0))
		}
		result[level] = append(result[level], node.Val)
		help(node.Left, level+1)
		help(node.Right, level+1)
	}
	help(root, 0)
	return result
}

// 迭代, 使用数组
func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	currentLever := []*TreeNode{root}
	for len(currentLever) > 0 {
		var values []int
		var nextLever []*TreeNode
		for _, node := range currentLever {
			values = append(values, node.Val)
			if node.Left != nil {
				nextLever = append(nextLever, node.Left)
			}
			if node.Right != nil {
				nextLever = append(nextLever, node.Right)
			}
		}
		result = append(result, values)
		currentLever = nextLever
	}
	return result
}

// 迭代，使用队列
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		var values []int
		currentLen := queue.Len()
		for i := 0; i < currentLen; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			values = append(values, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result = append(result, values)
	}
	return result
}

/* 迭代，节点标记法
在出入栈的时候，标记节点，具体为：
标记节点的状态，新节点为false，已使用（在这道题里是指将节点值追加到结果数组）的节点true。
如果遇到未标记的节点，则将其标记为true，然后将其右节点、左节点、自身依次入栈。
注意同时标记节点的层级，根节点标为0；
每次入栈将自节点的层级标为当前节点层级+1
如果遇到的节点标记为true，则使用该节点。
*/
func levelOrder3(root *TreeNode) [][]int {
	type item struct {
		node   *TreeNode
		marked bool
		level  int
	}
	var result [][]int
	stack := []*item{{node: root, level: 0}}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.node == nil {
			continue
		}
		if top.marked {
			if result == nil || len(result) == top.level {
				result = append(result, make([]int, 0))
			}
			result[top.level] = append(result[top.level], top.node.Val)
			continue
		}
		top.marked = true
		stack = append(stack, &item{node: top.node.Right, level: top.level + 1})
		stack = append(stack, &item{node: top.node.Left, level: top.level + 1})
		stack = append(stack, top)
	}
	return result
}
