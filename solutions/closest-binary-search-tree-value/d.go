/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package closest_binary_search_tree_value

import "math"

/*
270. 最接近的二叉搜索树值
https://leetcode-cn.com/problems/closest-binary-search-tree-value

给定一个不为空的二叉搜索树和一个目标值 target，请在该二叉搜索树中找到最接近目标值 target 的数值。

注意：

给定的目标值 target 是一个浮点数
题目保证在该二叉搜索树中只会存在一个最接近目标值的数
示例：

输入: root = [4,2,5,1,3]，目标值 target = 3.714286

    4
   / \
  2   5
 / \
1   3

输出: 4
*/

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* 即在bst中寻找一个元素target，若不存在，返回最接近的
常规二分搜索
*/
func closestValue(root *TreeNode, target float64) int {
	r := root.Val
	for root != nil {
		if math.Abs(float64(root.Val)-target) < math.Abs(float64(r)-target) {
			r = root.Val
		}
		if target < float64(root.Val) {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return r
}
