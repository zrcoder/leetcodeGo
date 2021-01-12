/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package merge_two_sorted_lists

/*
21. 合并两个有序链表 https://leetcode-cn.com/problems/merge-two-sorted-lists

将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 遍历
// 时间复杂度O(m+n), 空间复杂度O(1)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	for p := dummy; l1 != nil || l2 != nil; p = p.Next {
		if l1 != nil && l2 != nil && l1.Val < l2.Val || l2 == nil {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
	}
	l1, dummy.Next = dummy.Next, nil
	return l1
}

// 递归的解法
//时间复杂度与空间复杂度都是O(m+n)
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l2.Next, l1)
	return l2
}
