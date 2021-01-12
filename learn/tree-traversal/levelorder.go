/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package tree_traversal

import "container/list"

func simpleLevelOrder(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var result []string
	currLever := []*TreeNode{root}
	for len(currLever) > 0 {
		var nextLever []*TreeNode
		for _, v := range currLever {
			result = append(result, v.Val)
			for _, c := range v.Children {
				nextLever = append(nextLever, c)
			}
		}
		currLever = nextLever
	}
	return result
}

// 上面是基础层次遍历，下面稍微精确控制了下，需要返回一个数组，包含每一层的结点

// 迭代版。借助一个queue实现，或者用数组也可以
func levelOrder(root *TreeNode) [][]string {
	if root == nil {
		return nil
	}
	var result [][]string
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		var values []string
		total := queue.Len()
		for i := 0; i < total; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			values = append(values, node.Val)
			if len(node.Children) == 0 {
				continue
			}
			for _, v := range node.Children {
				queue.PushBack(v)
			}
		}
		result = append(result, values)
	}
	return result
}

// 递归版
func levelOrder1(root *TreeNode) [][]string {
	var result [][]string
	var bfs func(node *TreeNode, level int)
	bfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(result) == level {
			result = append(result, make([]string, 0))
		}
		result[level] = append(result[level], node.Val)
		for _, v := range node.Children {
			bfs(v, level+1)
		}
	}
	bfs(root, 0)
	return result
}
