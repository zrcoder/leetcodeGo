/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package recover_binary_search_tree

/*
99. 恢复二叉搜索树
https://leetcode-cn.com/problems/recover-binary-search-tree

二叉搜索树中的两个节点被错误地交换。

请在不改变其结构的情况下，恢复这棵树。

示例 1:
输入: [1,3,null,null,2]
   1
  /
 3
  \
   2
输出: [3,1,null,null,2]
   3
  /
 1
  \
   2

示例 2:
输入: [3,1,4,null,null,2]
  3
 / \
1   4
   /
  2
输出: [2,1,4,null,null,3]
  2
 / \
1   4
   /
  3
进阶:
使用 O(n) 空间复杂度的解法很容易实现。
你能想出一个只使用常数空间的解决方案吗？
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
如果一个已排序数组的两个元素被交换了，要怎么恢复？

看如下序列：
1 2 3 4 5 6 7 8 9
情况一：如果相邻的元素（如3和4）被交换了，就会变为
1 2 4 3 5 6 7 8 9
情况二：相隔一个元素的两个元素（如3和5）被交换，就会变为
1 2 5 4 3 6 7 8 9
情况三：如果不相邻的元素（如3和8）被交换了，就会变为
1 2 8 4 5 6 7 3 9

对于情况一，遍历时可以找到一个逆序序列 4，3
对于情况二，遍历时可以找到两个逆序序列5,4和4，3
对于情况三，遍历时可以找到两个逆序序列8,4 和7，3

那么可以遍历数组，如果找到一个逆序序列，交换逆序序列的两个元素即可；
如果找到两个逆序序列，交换第一个逆序序列前一个元素和第二个逆序序列第二个元素即可

// nums是一个升序序列，但是其中两个元素被交换
// findTwoSwapped返回已经交换的两个元素的索引
func findTwoSwapped(nums []int) (int, int) {
	first, second := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if first == -1 {
				first = nums[i]
				second = nums[i+1]
			} else {
				second = nums[i+1]
				break
			}
		}
	}
	return first, second
}

BST的中序遍历就是一个单调递增序列，可以在中序遍历时使用以上数组中的方法来解决
不改变树的结构，通过交换两个节点的Val来达到交换节点的目的

时间复杂度：最好的情况下是O(1)；最坏的情况下是交换节点之一是最右边的叶节点时，此时是O(N)。
空间复杂度：最大是O(H) 来维持递归调用堆栈的大小，其中 H是树的高度。
*/
func recoverTree(root *TreeNode) {
	var prev, first, second *TreeNode
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)

		if prev != nil && prev.Val > node.Val {
			if first == nil { // 找到了第一个逆序对
				first = prev
				second = node
			} else { // 找到了第二个逆序对
				second = node
				return
			}
		}
		prev = node

		inorder(node.Right)
	}

	inorder(root)
	if first == nil {
		return
	}
	first.Val, second.Val = second.Val, first.Val
}
