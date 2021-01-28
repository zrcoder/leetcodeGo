/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package merge_k_sorted_lists

/*
23. 合并K个排序链表 https://leetcode-cn.com/problems/merge-k-sorted-lists/
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
可以两两合并
首先实现将两个链表合并的merge函数
1.朴素实现就是
	var r *ListNode
	for _, v := range lists {
		r = merge(r, v)
	}
	return r
实际测试这样花费的时间是220ms
每次合并有一些相同的节点都会参与比较，比如lists[0]里的节点，参与了每一次合并

2.修改为比较均衡的合并：即相邻的链表先两两合并，再将合并后的链表两两合并，不断重复，直到所有链表合并完成
时间复杂度O(nlogk)，其中k是链表数目， n是每次合并的两个链表的节点总数
空间复杂度O(1)
实际测试修改后的耗时在10ms之内
*/
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	for interval := 1; interval < n; interval *= 2 {
		for i := 0; i+interval < n; i += interval * 2 {
			lists[i] = merge(lists[i], lists[i+interval])
		}
	}
	return lists[0]
}

func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for end := len(lists) - 1; end > 0; {
		for from := 0; from < end; from++ {
			lists[from] = merge(lists[from], lists[end])
			end--
		}
	}
	return lists[0]
}

func merge(l1, l2 *ListNode) *ListNode {
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
