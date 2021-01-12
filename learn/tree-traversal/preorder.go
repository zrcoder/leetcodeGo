/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package tree_traversal

import "container/list"

// 递归版
func preorder(root *TreeNode) []string {
	var result []string
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		for _, v := range node.Children {
			dfs(v)
		}
	}
	dfs(root)
	return result
}

// 迭代版，借助一个栈实现
func preorder1(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var result []string
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		result = append(result, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack.PushBack(node.Children[i])
		}
	}
	return result
}
