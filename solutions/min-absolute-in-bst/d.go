/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package min_absolute_in_bst

import "math"

/*
530. 二叉搜索树的最小绝对差 https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst

给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。

示例：

输入：

   1
    \
     3
    /
   2

输出：
1

解释：
最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。


提示：

树中至少有 2 个节点。
本题与 783 https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/ 相同
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
如果有一个整数数组，怎么获取任意两个元素之间的最小差值呢？如果这个数组是排好序的呢？
如果无序，需要两层循环计算所有两两元素差值来找最小差值
如果是排好序的，那么最小的差值只会出现在相邻元素之间，一层循环就够了

BST中序遍历即得到一个排序好的列表；问题划归为求一个已经排序数组中相邻元素的最小差值
当然不用真用一个数组存储，在遍历过程中即可求每个差值
*/
func getMinimumDifference(root *TreeNode) int {
	result := math.MaxInt32
	var prev *TreeNode
	var inorder func(*TreeNode)
	inorder = func(t *TreeNode) {
		if t == nil {
			return
		}
		inorder(t.Left)
		if prev != nil && t.Val-prev.Val < result {
			result = t.Val - prev.Val
		}
		prev = t
		inorder(t.Right)
	}
	inorder(root)
	return result
}
