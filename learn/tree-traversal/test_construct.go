/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package tree_traversal

var testTree = constructTestTree()

func constructTestTree() *TreeNode {
	/* this test tree likes:

	          A
	        / | \
	       B  C  D
	      /|     |
	     E F     G
	    /|\
	   H I J

	*/

	// level 0
	root := &TreeNode{Val: "A"}
	// lever 1
	x, y, z := &TreeNode{Val: "B"}, &TreeNode{Val: "C"}, &TreeNode{Val: "D"}
	root.Children = make([]*TreeNode, 3)
	root.Children[0], root.Children[1], root.Children[2] = x, y, z
	// lever 2
	x, y, z = &TreeNode{Val: "E"}, &TreeNode{Val: "F"}, &TreeNode{Val: "G"}
	first, last := root.Children[0], root.Children[2]
	first.Children = make([]*TreeNode, 2)
	first.Children[0], first.Children[1] = x, y
	last.Children = make([]*TreeNode, 1)
	last.Children[0] = z
	// lever 3
	x, y, z = &TreeNode{Val: "H"}, &TreeNode{Val: "I"}, &TreeNode{Val: "J"}
	first = first.Children[0]
	first.Children = make([]*TreeNode, 3)
	first.Children[0], first.Children[1], first.Children[2] = x, y, z
	return root
}

var levelOrderWant = [][]string{
	{"A"},
	{"B", "C", "D"},
	{"E", "F", "G"},
	{"H", "I", "J"},
}
var simpleLevelOrderWant = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
var preOrderWant = []string{"A", "B", "E", "H", "I", "J", "F", "C", "D", "G"}
