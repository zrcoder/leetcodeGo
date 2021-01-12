/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package verify_postorder_sequence_inbinary_search_tree

import (
	"container/list"
	"math"
)

/*
面试题33. 二叉搜索树的后序遍历序列  https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。
假设输入的数组的任意两个数字都互不相同。

参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
示例 2：

输入: [1,3,2,6,5]
输出: true


提示：

数组长度 <= 1000
*/

/*
思路同判断一个序列是不是BST的先序遍历结果
有递归和借助单调栈的两种解法
*/
// 递归解法
func verifyPostorder(postorder []int) bool {
	n := len(postorder)
	if n == 0 {
		return true
	}
	root := postorder[n-1]
	i := 0
	for ; i < n-1; i++ {
		if postorder[i] > root {
			break
		}
	}
	// i为右子树根所在位置
	for j := i; j < n-1; j++ {
		if postorder[j] < root {
			return false
		}
	}
	// 截止目前，检查了左子树所有元素小于root，右子树所有元素大于root
	if i > 0 && !verifyPostorder(postorder[:i]) {
		return false
	}
	if i < n-1 && !verifyPostorder(postorder[i:n-1]) {
		return false
	}
	return true
}

/* 单调栈解法
后序遍历的逆序是什么？root -> right -> left

维护一个单调递增栈；从后往前遍历给定的数组：
1.每次检查元素是否不大于上界，如果大于则不是bst的后序遍历序列
2.如果一直递增意味着在右子树中，将这些元素一一入栈；
3.如果突然不再递增，说明进入了某个左子树中；这时候需要把对应的右子树元素全部出栈
同时记录新的上界值，即对应根节点的值，也就是比当前遍历的元素大的栈中最后一个元素

时间复杂度O(n), 每个元素进出栈共2次；空间复杂度O（n）
*/
func verifyPostorder1(postorder []int) bool {
	high := math.MaxInt32
	stack := list.New() // 也可用切片实现栈；如果允许修改原数组，甚至可以用原数组来做这个栈
	for i := len(postorder) - 1; i >= 0; i-- {
		if postorder[i] > high {
			return false
		}
		for stack.Len() > 0 && stack.Back().Value.(int) > postorder[i] {
			high = stack.Remove(stack.Back()).(int)
		}
		stack.PushBack(postorder[i])
	}
	return true
}
