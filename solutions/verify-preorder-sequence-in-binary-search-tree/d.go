/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package verify_preorder_sequence_in_binary_search_tree

import (
	"math"
)

/*
255. 验证前序遍历序列二叉搜索树 https://leetcode-cn.com/problems/verify-preorder-sequence-in-binary-search-tree
给定一个整数数组，你需要验证它是否是一个二叉搜索树正确的先序遍历序列。

你可以假定该序列中的数都是不相同的。

参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：
输入: [5,2,6,1,3]
输出: false

示例 2：
输入: [5,2,1,3,6]
输出: true
*/

// 递归解法；如果题目要求判断是否是后序遍历序列，可以根据这个递归方式简单修改解决
func verifyPreorder0(preorder []int) bool {
	if len(preorder) == 0 {
		return true
	}
	i := 1
	root := preorder[0]
	for ; i < len(preorder); i++ {
		if preorder[i] > root {
			break
		}
	}
	// i 为左右子树分界点，精确说是右子树根节点所在位置
	for j := i; j < len(preorder); j++ {
		if preorder[j] < root {
			return false
		}
	}
	// 截止目前，保证了i之前的元素（不包括最开始的root）均小于root， i之后的元素包括i处元素均大于root
	if !verifyPreorder0(preorder[1:i]) {
		return false
	}
	if i < len(preorder) && !verifyPreorder0(preorder[i:]) {
		return false
	}
	return true
}

/*
维护一个单调递减栈，遍历序列：
1.检查元素是否不小于下界，如果小于，说明不是bst的前序遍历序列
2.如果一直递减，说明在左子树中，将这些元素一一入栈；
3.如果突然不再递减，说明到达了某个右子树，则将其对应的左子树及其父节点全部出栈（即将比当前元素小的元素都出栈）
并将对应的根节点（即最后一个小于当前元素的元素）记为当前的下限

时间复杂度O(n), 每个元素进出栈共2次；空间复杂度O（n）

如果题目要求判断是否是后序遍历序列， 可以从后向前遍历数组，相当于根->右->左的遍历方式
*/
func verifyPreorder(preorder []int) bool {
	var stack []int // 也可以用list实现栈
	low := math.MinInt32
	for _, v := range preorder {
		if v < low {
			return false
		}
		for len(stack) > 0 && stack[len(stack)-1] < v {
			low = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
	return true
}

// 如果允许修改原数组，用原数组模拟单调递减栈， 空间复杂度降为常数级
func verifyPreorder1(preorder []int) bool {
	min := math.MinInt32
	i := -1
	for _, v := range preorder {
		if v < min {
			return false
		}
		for i >= 0 && preorder[i] < v {
			min = preorder[i]
			i--
		}
		i++
		preorder[i] = v
	}
	return true
}
