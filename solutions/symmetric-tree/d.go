/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package symmetric_tree

/*
101. 对称二叉树 https://leetcode-cn.com/problems/symmetric-tree
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
说明:

如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
*/

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

// 递归式
func isMirror(t1, t2 *TreeNode) bool {
	switch {
	case t1 == nil && t2 == nil:
		return true
	case t1 == nil || t2 == nil:
		return false
	case t1.Val != t2.Val:
		return false
	default:
		return isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
	}
}

// 迭代式
func isMirror1(t1, t2 *TreeNode) bool {
	var list []*TreeNode
	list = append(list, t1, t2)
	for len(list) > 0 {
		/*
			取出list里的最后两个，并缩短list
			这里也可以取出前边两个，但是在缩短list的时候（list = list[2:])可能会导致底层数组更容易扩容
			也可以用真正的list而不是切片
		*/
		n := len(list)
		t1, t2 = list[n-2], list[n-1]
		list = list[:n-2]
		switch {
		case t1 == nil && t2 == nil:
			continue
		case t1 == nil || t2 == nil:
			return false
		case t1.Val != t2.Val:
			return false
		default:
			list = append(list, t1.Left, t2.Right, t1.Right, t2.Left)
		}
	}
	return true
}
