/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package tree_traversal

type TreeNode struct {
	Children []*TreeNode
	Val      string // 简单起见，假设节点中的值都是字符串
}
