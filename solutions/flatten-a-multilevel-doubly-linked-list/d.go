/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package flatten_a_multilevel_doubly_linked_list

/*
430. 扁平化多级双向链表 https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list/

多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。
这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。

给你位于列表第一级的头节点，请你扁平化列表，使所有结点出现在单级双链表中。

示例 1：
输入：head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
输出：[1,2,3,7,8,11,12,9,10,4,5,6]
解释：

输入的多级列表如下图所示：
扁平化后的链表如下图：

示例 2：
输入：head = [1,2,null,3]
输出：[1,3,2]
解释：
输入的多级列表如下图所示：
  1---2---NULL
  |
  3---NULL

示例 3：
输入：head = []
输出：[]

如何表示测试用例中的多级链表？

以 示例 1 为例：

 1---2---3---4---5---6--NULL
         |
         7---8---9---10--NULL
             |
             11--12--NULL
序列化其中的每一级之后：

[1,2,3,4,5,6,null]
[7,8,9,10,null]
[11,12,null]
为了将每一级都序列化到一起，我们需要每一级中添加值为 null 的元素，以表示没有节点连接到上一级的上级节点。

[1,2,3,4,5,6,null]
[null,null,7,8,9,10,null]
[null,11,12,null]
合并所有序列化结果，并去除末尾的 null 。

[1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]

提示：

节点数目不超过 1000
1 <= Node.val <= 10^5
*/

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

/*
遍历链表，遇到child就将其插入当前节点和下一个节点之间
假设节点总数为n，时间复杂度为O(n), 空间复杂度O(1)
*/
func flatten(root *Node) *Node {
	for p := root; p != nil; p = p.Next {
		if p.Child == nil {
			continue
		}
		next := p.Next
		p.Next, p.Child.Prev = p.Child, p
		q := p.Child
		for q.Next != nil {
			q = q.Next
		}
		q.Next = next
		if next != nil {
			next.Prev = q
		}
		p.Child = nil
	}
	return root
}

/*
递归解法
*/
func flatten1(root *Node) *Node {
	if root == nil {
		return nil
	}
	dummy := &Node{Next: root}
	_ = dfs(dummy, root)
	dummy.Next.Prev = nil
	root, dummy.Next = dummy.Next, nil
	return root
}

// 返回扁平化后的尾节点
func dfs(prev, curr *Node) *Node {
	if curr == nil {
		return prev
	}
	prev.Next, curr.Prev = curr, prev
	next := curr.Next
	tail := dfs(curr, curr.Child)
	curr.Child = nil
	return dfs(tail, next)
}
