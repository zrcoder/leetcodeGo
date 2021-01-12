/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package next_greater_node_in_linked_list

import "container/list"

/*
1019. 链表中的下一个更大结点 https://leetcode-cn.com/problems/next-greater-node-in-linked-list

给出一个以头结点 head 作为第一个结点的链表。链表中的结点分别编号为：node_1, node_2, node_3, ... 。
每个结点都可能有下一个更大值（next larger value）：对于 node_i，如果其 next_larger(node_i) 是 node_j.val，
那么就有 j > i 且  node_j.val > node_i.val，而 j 是可能的选项中最小的那个。如果不存在这样的 j，那么下一个更大值为 0 。
返回整数答案数组 answer，其中 answer[i] = next_larger(node_{i+1}) 。
注意在下面的示例中，诸如 [2,1,5] 这样的输入（不是输出）是链表的序列化表示，其头结点的值为 2，第二个结点值为 1，第三个节点值为 5 。


示例 1：
输入：[2,1,5]
输出：[5,5,0]

示例 2：
输入：[2,7,4,3,5]
输出：[7,0,5,5,0]

示例 3：
输入：[1,7,5,1,9,2,5,1]
输出：[7,9,9,9,0,5,0,0]


提示：
对于链表中的每个节点，1 <= node.val <= 10^9
给定列表的长度在 [0, 10000] 范围内
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
朴素实现
时间复杂度O(n^2),空间复杂度O(n)
*/
func nextLargerNodes1(head *ListNode) []int {
	var r []int
	for head != nil {
		r = append(r, nextLargerVal(head))
		head = head.Next
	}
	return r
}

func nextLargerVal(n *ListNode) int {
	for p := n.Next; p != nil; p = p.Next {
		if p.Val > n.Val {
			return p.Val
		}
	}
	return 0
}

/*
借助单调栈
关于单调栈，参考如下讲解：
https://labuladong.gitbook.io/algo/shu-ju-jie-gou-xi-lie/dan-tiao-zhan

时间复杂度降为O(n),每个节点索引进栈出栈各一次
空间复杂度O(n), 主要为辅助数组nums、栈及结果数组所占的空间
*/
func nextLargerNodes(head *ListNode) []int {
	var nums []int
	for p := head; p != nil; p = p.Next {
		nums = append(nums, p.Val)
	}
	result := make([]int, len(nums))
	stack := list.New() // 单调递减栈；存储nums索引
	for i, v := range nums {
		for stack.Len() > 0 && nums[stack.Back().Value.(int)] < v {
			result[stack.Remove(stack.Back()).(int)] = v
		}
		stack.PushBack(i)
	}
	return result
}
