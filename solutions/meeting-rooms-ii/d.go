package meeting_rooms_ii

import (
	"container/heap"
	"sort"
)

/*
 253. 会议室 II https://leetcode-cn.com/problems/meeting-rooms-ii
给定一个会议时间安排的数组，每个会议时间都会包括开始和结束的时间 [[s1,e1],[s2,e2],...] (si < ei)，为避免会议冲突，
同时要考虑充分利用会议室资源，请你计算至少需要多少间会议室，才能满足这些会议安排。

示例 1:
输入: [[0, 30],[5, 10],[15, 20]]
输出: 2

示例 2:
输入: [[7,10],[2,4]]
输出: 1
*/

/*
 1. 朴素实现
 遍历所有会议，当前会议和前边所有会议比较，如果有交叉时间则需要增加会议室; 当然即使所有会议时间不冲突，至少需要一个会议室
 需要考虑这样的用例：[[9,10],[4,9],[4,17]]，用以上的方法求的的会议室是3，但实际上应该是2.
 如果把最后一个会议放到最前边，变成[[4,17],[9,10],[4,9]]，就能得到正确答案，为什么呢？
 想一想容易明白，需要事先把所有会议排序，按照开始时间升序排序即可
 时间复杂度O(n^2), 空间复杂度O(1)
*/
func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := 0
	for i := range intervals {
		result = max(result, cal(intervals, i))
	}
	return result + 1
}

func cal(intervals [][]int, i int) int {
	result := 0
	r := intervals[i]
	for _, v := range intervals[:i] {
		if max(v[0], r[0]) < min(v[1], r[1]) { // 会议r和v有冲突
			result++
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
 2. 在朴素实现基础上，加一个小顶堆来帮忙：
 和朴素实现一样，首先依然需要把所有会议按照开始时间排序，然后开始一一处理每个会议
 维护一个小顶堆，里边存放每个会议的结束时间。
 堆顶是最早的结束时间，遍历时，把当前会议的结束时间入堆，入堆前要注意开始时间如果晚于堆顶的时间，则堆顶元素出堆
 遍历处理完所有会议后，堆内元素的个数即需要的会议室数

 时间复杂度降为O(nlgn)，最坏情况所有会议都入堆；空间复杂度升为O(n)
*/
type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func minMeetingRooms1(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	h := &IntHeap{}
	for _, v := range intervals {
		if h.Len() > 0 && (*h)[0] <= v[0] {
			_ = heap.Pop(h)
		}
		heap.Push(h, v[1])
	}
	return h.Len()
}

/*
 3. 有序化

 理解成上下车问题会比较容易解决，不用在意是谁上车还是下车，只需要注意什么时候上下车就可以。 以第一个示例来说：
 [[0, 30],[5, 10],[15, 20]]
 ↑    ↑    ↓     ↑      ↓             ↓
 0----5----10----15-----20-----------30-->
 这样可以把上车和下车的时间分成两组，通过两个指针滑动的方式，判断同时在车上的最大数就可以了。

 时间复杂度O(nlgn)， 空间复杂度O(n)
*/
func minMeetingRooms2(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	starts := make([]int, n)
	ends := make([]int, n)
	for i, v := range intervals {
		starts[i] = v[0]
		ends[i] = v[1]
	}
	sort.Ints(starts)
	sort.Ints(ends)
	result := 0
	for s, e := 0, 0; s < n; s++ {
		if starts[s] >= ends[e] {
			result--
			e++
		}
		result++
	}
	return result
}
