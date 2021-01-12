/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	carry := 0 // must be 0 or 1
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		current.Val = sum % 10
		if l1 != nil || l2 != nil || carry != 0 {
			current.Next = &ListNode{}
			current = current.Next
		}
	}
	return result
}
