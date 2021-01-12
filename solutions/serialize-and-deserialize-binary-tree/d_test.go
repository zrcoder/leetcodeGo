/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package serialize_and_deserialize_binary_tree

import (
	"fmt"
)

func makeTestTree() *TreeNode {
	/* this tree is:
	  1
	 / \
	2   3
	   / \
	  4   5
	*/
	t := &TreeNode{Val: 1}
	l := &TreeNode{Val: 2}
	r := &TreeNode{Val: 3}
	t.Left = l
	t.Right = r
	l = &TreeNode{Val: 4}
	r = &TreeNode{Val: 5}
	t.Right.Left = l
	t.Right.Right = r
	return t
}

func Example() {
	t := makeTestTree()
	printTreePreorder(t)
	fmt.Println()
	c := Constructor()
	s := c.serialize(t)
	fmt.Println(s)
	t = c.deserialize(s)
	printTreePreorder(t)
	fmt.Println()

	// output:
	//
	// 1,2,#,#,3,4,#,#,5,#,#,
	// 1,2,#,#,3,4,#,#,5,#,#,
	// 1,2,#,#,3,4,#,#,5,#,#,
}

func printTreePreorder(root *TreeNode) {
	if root == nil {
		fmt.Print("#,")
		return
	}
	fmt.Print(root.Val, ",")
	printTreePreorder(root.Left)
	printTreePreorder(root.Right)
}
