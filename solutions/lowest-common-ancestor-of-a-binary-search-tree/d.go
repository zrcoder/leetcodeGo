/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package lowest_common_ancestor_of_a_binary_search_tree

/*
235. 二叉搜索树的最近公共祖先 https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree
给定一个二叉搜索树, 找到该树中两个指定结点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，
满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个结点也可以是它自己的祖先）。”

例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]

示例 1:
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6
解释: 结点 2 和结点 8 的最近公共祖先是 6。

示例 2:
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 结点 2 和结点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先结点可以为结点本身。

说明:
所有结点的值都是唯一的。
p、q 为不同结点且均存在于给定的二叉搜索树中。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 根据bst性质递归；时空复杂度都是O(N)
func lowestCommonAncestor0(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}

//  可以处理成尾递归；空间复杂度降为O(1)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	lo, hi := p.Val, q.Val
	if lo > hi {
		lo, hi = hi, lo
	}
	if root.Val > lo && root.Val < hi {
		return root
	}
	if lo > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return lowestCommonAncestor(root.Left, p, q)
}

// 根据bst性质迭代;时间复杂度O(n),空间复杂度O(1)
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		switch {
		case root.Val > p.Val && root.Val > q.Val:
			root = root.Left
		case root.Val < p.Val && root.Val < q.Val:
			root = root.Right
		default:
			return root
		}
	}
	return nil
}

// 不借助bst性质，对一个普遍二叉树递归;时空复杂度都是O(N)
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
