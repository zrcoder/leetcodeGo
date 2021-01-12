/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package maximum

import "container/list"

/*
239. 滑动窗口最大值 https://leetcode-cn.com/problems/sliding-window-maximum
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7


提示：
你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。


进阶：
你能在线性时间复杂度内解决此题吗？
*/
/*
朴素实现：用一个list或切片模拟窗口，或者直接用两个指针确定窗口，遍历确定窗口里的最大值
主要耗时在寻找最大值上，总的时间复杂度是O(n*k)；空间复杂度主要在结果数组：O(N−k+1)
以下为双指针模拟窗口的实现
*/
func maxSlidingWindow0(nums []int, k int) []int {
	if k > len(nums) || len(nums) == 0 {
		return nil
	}
	result := make([]int, len(nums)-k+1)
	left, right := 0, k-1
	result[0] = getMaxInWindow(nums, left, right)
	for i := k; i < len(nums); i++ {
		left++
		right++
		result[i-k+1] = getMaxInWindow(nums, left, right)
	}
	return result
}

func getMaxInWindow(nums []int, left, right int) int {
	r := nums[left]
	for i := left + 1; i <= right; i++ {
		if nums[i] > r {
			r = nums[i]
		}
	}
	return r
}

/*
用一个单调非递增队列来模拟窗口，窗口里的最大值就是第一个元素。这样可以在O(1)复杂度下获知最大值。
不修改原数组的话，额外用一个list或者切片来实现这个单调队列
时间复杂度，O（N），每个元素被处理两次——入队、出队
空间复杂度，输出数组使用O(N−k+1) 空间；队列用切片实现的话使用O（N)空间，list实现的话是O（K）空间；综合为O（N）复杂度
*/
// 使用list实现单调队列，队列里存储元素值（也可以改成存储索引）
func maxSlidingWindow1(nums []int, k int) []int {
	if k > len(nums) || len(nums) == 0 {
		return nil
	}
	result := make([]int, len(nums)-k+1)
	queue := list.New() // 单调队列，并非存储所有窗口里的元素，而是维持单调非递增序列
	for i := 0; i < len(nums); i++ {
		if i >= k {
			dequeue(queue, nums[i-k])
		}
		enqueue(queue, nums[i])
		if i >= k-1 {
			result[i-k+1] = max(queue)
		}
	}
	return result
}

func enqueue(queue *list.List, num int) {
	// 入队前将队列里已有小于当前值的元素移除
	for queue.Len() > 0 && queue.Back().Value.(int) < num {
		_ = queue.Remove(queue.Back())
	}
	queue.PushBack(num)
}

func dequeue(queue *list.List, value int) {
	// 只有当要出队的元素恰好等于队头最大元素时，才出队
	if queue.Len() > 0 && queue.Front().Value.(int) == value {
		_ = queue.Remove(queue.Front())
	}
}

func max(queue *list.List) int {
	return queue.Front().Value.(int)
}

// 使用切片实现单调队列， 队列里存储元素索引（也可以改成存储值）
func maxSlidingWindow(nums []int, k int) []int {
	var result, queue []int // queue是一个单调队列，并非存储所有窗口里的元素，而是维持单调非递增序列； 队列中存储索引
	for i, v := range nums {
		// 出队
		if i >= k && len(queue) > 0 && queue[0] <= i-k {
			queue = queue[1:]
		}
		// 入队
		for len(queue) > 0 && nums[queue[len(queue)-1]] < v { // 入队前将队列里小于当前值的元素移除
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i) // 当前元素入队， 这里入队的是元素的索引
		// 将窗口中最大元素加入结果
		if i >= k-1 {
			result = append(result, nums[queue[0]])
		}
	}
	return result
}
