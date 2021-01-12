# [2. 两数相加](https://leetcode.com/problems/add-two-numbers/)

## 题目
```
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
```

## 分析

```
(2 -> 4 -> 3)是 342

(5 -> 6 -> 4)是 465

(7 -> 0 -> 8)是 807

342 + 465 = 807
```

所以，题目的本意是，把整数换了一种表达方式后，实现其加法。

可以想一想为什么要逆序~~

设计程序时候，需要处理的点有

1). 位上的加法，需要处理进位问题
2). 如何进入下一位运算
3). 按位相加结束后，也还需要处理进位问题

```go
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
```