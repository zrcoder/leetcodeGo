/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package construct_binary_tree_from_preorder_and_inorder_traversal

/*
105. 从前序与中序遍历序列构造二叉树 https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
如果给定后序序列和中序序列，和这道题的解法是非常类似的
首先，可以迅速确定根节点的值，即为先序序列的第一个（或后序序列的最后一个）
这个值在inorder的位置，把inorder划分成左右两部分，同样也把preorder（或postorder）划分成了两部分
问题可以递归解决
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := search(inorder, root.Val)
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
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
