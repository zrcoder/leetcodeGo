/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package construct_binary_tree_from_inorder_and_postorder_traversal

/*
106. 从中序与后序遍历序列构造二叉树 https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
*/

/*
如果给定先序序列和中序序列，和这道题的解法是非常类似的
首先，可以迅速确定根节点的值，即为后序序列的最后一个（或先序序列的第一个）
这个值在inorder的位置，把inorder划分成左右两部分，同样也把postorder（或preorder）划分成了两部分
问题可以递归解决
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	i := search(inorder, root.Val)
	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
	return root
}

func search(arr []int, val int) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}
