/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package max_queue

/*
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：

输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]
示例 2：

输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]


限制：

1 <= push_back,pop_front,max_value的总操作数 <= 10000
1 <= value <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
和最小栈的问题类似， 但也有重要区别
进队和出队没法做到O(1), 可以用二分法做到O(lgn)， n代表max数组的长度，在最坏的情况下，是队列的长度
*/
type MaxQueue struct {
	queue []int
	maxes []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (q *MaxQueue) Max_value() int {
	if len(q.maxes) == 0 {
		return -1
	}
	return q.maxes[0]
}

func (q *MaxQueue) Push_back(value int) {
	q.queue = append(q.queue, value)
	i := search(q.maxes, value)
	q.maxes = append(q.maxes[:i], value)
}

func search(nums []int, value int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > value {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func (q *MaxQueue) Pop_front() int {
	if len(q.queue) == 0 {
		return -1
	}
	v := q.queue[0]
	q.queue = q.queue[1:]
	if v == q.maxes[0] {
		q.maxes = q.maxes[1:]
	}
	return v
}
