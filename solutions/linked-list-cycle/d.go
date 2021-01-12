/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package linked_list_cycle

/*
给定一个链表，判断链表中是否有环。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：
输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：false
解释：链表中没有环。

进阶：
你能用 O(1)（即，常量）内存解决此问题吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-cycle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 借助一个set解决；空间、时间复杂度都是O(n)
func hasCycle1(head *ListNode) bool {
	set := make(map[*ListNode]struct{}, 0)
	for p := head; p != nil; p = p.Next {
		if _, found := set[p]; found {
			return true
		}
		set[p] = struct{}{}
	}
	return false
}

// 双指针，空间复杂度O(1); 时间复杂度O(n)
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
