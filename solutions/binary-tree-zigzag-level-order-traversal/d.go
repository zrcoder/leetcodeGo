package binary_tree_zigzag_level_order_traversal

/*
103. 二叉树的锯齿形层次遍历 https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	level := []*TreeNode{root}
	toRight := true
	for len(level) > 0 {
		result = append(result, travasLevel(level, toRight))
		level = genNext(level)
		toRight = !toRight
	}
	return result
}

func travasLevel(level []*TreeNode, toRight bool) []int {
	result := make([]int, 0, len(level))
	if toRight {
		for _, n := range level {
			result = append(result, n.Val)
		}
	} else {
		for i := len(level) - 1; i >= 0; i-- {
			result = append(result, level[i].Val)
		}
	}
	return result
}

func genNext(level []*TreeNode) []*TreeNode {
	var next []*TreeNode
	for _, v := range level {
		if v.Left != nil {
			next = append(next, v.Left)
		}
		if v.Right != nil {
			next = append(next, v.Right)
		}
	}
	return next
}
