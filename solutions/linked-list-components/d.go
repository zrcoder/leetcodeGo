/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package linked_list_components

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
给定一个链表（链表结点包含一个整型值）的头结点 head。

同时给定列表 G，该列表是上述链表中整型值的一个子集。

返回列表 G 中组件的个数，这里对组件的定义为：链表中一段最长连续结点的值（该值必须在列表 G 中）构成的集合。

示例 1：

输入:
head: 0->1->2->3
G = [0, 1, 3]
输出: 2
解释:
链表中,0 和 1 是相连接的，且 G 中不包含 2，所以 [0, 1] 是 G 的一个组件，同理 [3] 也是一个组件，故返回 2。
示例 2：

输入:
head: 0->1->2->3->4
G = [0, 3, 1, 4]
输出: 2
解释:
链表中，0 和 1 是相连接的，3 和 4 是相连接的，所以 [0, 1] 和 [3, 4] 是两个组件，故返回 2。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-components
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
遍历链表，通过查看链表节点的值是否在G中来统计
用一个临时变量表示上一个节点的值是否在G中，这样就能正确统计
判断元素是否在数组G中比较耗时，可以将G里的元素放入哈希表来优化

假设链表和列表的长度分别为m，n，则时间复杂度是O(m+n)，空间复杂度是O(n)
*/
func numComponents(head *ListNode, G []int) int {
	if head == nil || len(G) == 0 {
		return 0
	}
	set := make(map[int]struct{}, len(G))
	for _, v := range G {
		set[v] = struct{}{}
	}
	count := 0
	isPrevNodeInG := false
	for head != nil {
		_, isCurrentNodeInG := set[head.Val]
		if isCurrentNodeInG && !isPrevNodeInG {
			count++
		}
		isPrevNodeInG = isCurrentNodeInG
		head = head.Next
	}
	return count
}

// 也可以在遍历链表的时候，同时看当前节点和下一个节点是否在G中来统计
func numComponents1(head *ListNode, G []int) int {
	if len(G) == 0 || head == nil {
		return 0
	}
	set := make(map[int]struct{}, len(G))
	for _, v := range G {
		set[v] = struct{}{}
	}
	count := 0
	for p := head; p != nil; p = p.Next {
		_, isCurrentNodeInG := set[p.Val]
		isNextNodeInG := false
		if p.Next != nil {
			_, isNextNodeInG = set[p.Next.Val]
		}
		if isCurrentNodeInG && !isNextNodeInG {
			count++
		}
	}
	return count
}
