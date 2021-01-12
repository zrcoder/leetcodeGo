/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package sort_list

import "sort"

/*
148. 排序链表 https://leetcode-cn.com/problems/sort-list/

在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
归并排序：自顶向下，使用递归
时间复杂度O(nlogn), 空间复杂度O(logn)，为栈空间 —— 空间复杂度不满足题意
*/
func sortList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	right := divide(head)
	return mergeSortedLists(sortList(head), sortList(right))
}

func divide(head *ListNode) *ListNode {
	/*
		快慢指针，最终慢指针到达中点
		我们希望最终切断中点前一个节点和中点，用prev来记录慢指针的前一个节点

		还有个不需要prev指针的做法，需要一开始将fast指向head.Next；
		首先divide的调用保证了head != nil && head.Next != nil, 可以放心这么写
		其次这么做以后，最终slow指向的是中点的前一个节点
	*/
	fast, slow, prev := head, head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		fast = fast.Next.Next
		slow = slow.Next
	}
	prev.Next = nil // 切断
	return slow
}
func mergeSortedLists(first, second *ListNode) *ListNode {
	dummy := new(ListNode)
	for p := dummy; first != nil || second != nil; p = p.Next {
		if first != nil && second != nil && first.Val <= second.Val || second == nil {
			p.Next = first
			first = first.Next
		} else {
			p.Next = second
			second = second.Next
		}
	}
	p := dummy.Next
	dummy.Next = nil
	return p
}

/*
自底向上归并排序
时间复杂度O(nlogn)， 空间复杂度O(1)
*/
func sortList(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	length := getLen(head)
	for step := 1; step < length; step *= 2 {
		p, head := dummy, dummy.Next
		for head != nil {
			first := head
			second := cut(first, step)
			head = cut(second, step)
			p.Next, p = merge(first, second)
		}
	}
	head, dummy.Next = dummy.Next, nil
	return head
}

func getLen(head *ListNode) int {
	length := 0
	for head != nil {
		head = head.Next
		length++
	}
	return length
}

func cut(head *ListNode, n int) *ListNode {
	for n > 1 && head != nil {
		head = head.Next
		n--
	}
	if n == 1 && head != nil {
		p := head.Next
		head.Next = nil
		return p
	}
	return nil
}

func merge(first *ListNode, second *ListNode) (*ListNode, *ListNode) {
	dummy := new(ListNode)
	p := dummy
	for ; first != nil || second != nil; p = p.Next {
		if (first != nil && second != nil && first.Val < second.Val) || second == nil {
			p.Next = first
			first = first.Next
		} else {
			p.Next = second
			second = second.Next
		}
	}
	first, dummy.Next = dummy.Next, nil
	return first, p
}

/*
快排，自顶向下，使用递归
选择一个标准值，将比它大的放在一个链表中，比它小的放在一个链表中，和它一样大的，放在另一个链表中。
然后针对小的和大的链表，继续排序。最终将三个链表按照小、相等、大进行连接。

最坏时间复杂度O(n^2), 平均时间复杂度O(nlogn), ——但因常数因子以及最坏情况出现的概率较小，实际比归并排序快
空间复杂度O(1)
*/
func sortList2(head *ListNode) *ListNode {
	return quickSort(head)
}
func quickSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	lowDummy, midDummy, highDummy := new(ListNode), new(ListNode), new(ListNode)
	low, mid, high := lowDummy, midDummy, highDummy
	val := head.Val
	for p := head; p != nil; p = p.Next {
		if p.Val < val {
			low.Next = p
			low = low.Next
		} else if p.Val > val {
			high.Next = p
			high = high.Next
		} else {
			mid.Next = p
			mid = mid.Next
		}
	}
	low.Next, mid.Next, high.Next = nil, nil, nil
	lowDummy.Next = quickSort(lowDummy.Next)
	highDummy.Next = quickSort(highDummy.Next)
	low = lowDummy
	for low.Next != nil {
		low = low.Next
	}
	low.Next = midDummy.Next
	mid.Next = highDummy.Next
	low = lowDummy.Next
	lowDummy.Next, midDummy.Next, highDummy.Next = nil, nil, nil
	return low
}

/*
如果不限定空间复杂度为常数级，可以这么玩：
用一个数组，装入链表所有节点，然后用标准库对数组排序即可
时间复杂度是O(nlogn), 空间复杂度O(n)
*/
func sortList9(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var arr []*ListNode
	for ; head != nil; head = head.Next {
		arr = append(arr, head)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Val < arr[j].Val
	})
	for i, v := range arr {
		if i == len(arr)-1 {
			v.Next = nil
		} else {
			v.Next = arr[i+1]
		}
	}
	return arr[0]
}
