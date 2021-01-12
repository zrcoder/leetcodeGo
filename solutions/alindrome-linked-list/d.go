/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package alindrome_linked_list

/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/*如果是数组，则很容易判断是否回文，从两端逐渐往中间遍历，每次比较两端是否相等即可
可以将链表元素一一存入数组再判断，不再赘述

链表的话，将后半段反转后和前半段一一对比；最后应该恢复原链表，即将后半段再次反转回来
O(n) 时间复杂度和 O(1) 空间复杂度
*/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverse(firstHalfEnd.Next)
	firstHalfEnd.Next = nil
	// 前半段与反转后的后半段对比
	q := secondHalfStart
	for p := head; q != nil; p, q = p.Next, q.Next {
		if p.Val != q.Val {
			break
		}
	}
	// 后半段再次反转，恢复到原来
	firstHalfEnd.Next = reverse(secondHalfStart)
	return q == nil
}

// 将链表分为两半，返回前一半末尾的节点；这里用了双指针技巧，也可以遍历统计链表长度来做
func endOfFirstHalf(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 反转一个链表，并返回反转后的头节点
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

/*
值得一提的是以下递归解法，即链表的后序遍历
空间复杂度是O(n)：
*/
func isPalindrome0(head *ListNode) bool {
	front := head
	var check func(curr *ListNode) bool
	check = func(curr *ListNode) bool {
		if curr == nil {
			return true
		}
		if !check(curr.Next) {
			return false
		}
		if curr.Val != front.Val {
			return false
		}
		front = front.Next
		return true
	}
	return check(head)
}
