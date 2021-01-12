/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package course_schedule_iii

import (
	"container/heap"
	"math"
	"sort"
)

/*
这里有 n 门不同的在线课程，他们按从 1 到 n 编号。每一门课程有一定的持续上课时间（课程时间）t 以及关闭时间第 d 天。
一门课要持续学习 t 天直到第 d 天时要完成，你将会从第 1 天开始。

给出 n 个在线课程用 (t, d) 对表示。你的任务是找出最多可以修几门课。

示例：

输入: [[100, 200], [200, 1300], [1000, 1250], [2000, 3200]]
输出: 3
解释:
这里一共有 4 门课程, 但是你最多可以修 3 门:
首先, 修第一门课时, 它要耗费 100 天，你会在第 100 天完成, 在第 101 天准备下门课。
第二, 修第三门课时, 它会耗费 1000 天，所以你将在第 1100 天的时候完成它, 以及在第 1101 天开始准备下门课程。
第三, 修第二门课时, 它会耗时 200 天，所以你将会在第 1300 天时完成它。
第四门课现在不能修，因为你将会在第 3300 天完成它，这已经超出了关闭日期。

提示:
整数 1 <= d, t, n <= 10,000 。
你不能同时修两门课程。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/course-schedule-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/* 对于两门课(t1, d1) 和 （t2, d2)， 若d1 <= d2, 则先修前者总是最优的
分析如下：假设在这两门课之前已经因为别的课花费了days天，则
先修第一门，需要满足： days + t1 <= d1 ...(1) 且 days + t1 + t2 <= d2 ...(2)
先修第二门，需要满足： days + t2 <= d2 ...(3) 且 days + t1 + t2 <= d1 ...(4)
如果(4)成立，因d1 <= d2, 则(2)必成立，且(1) 也成立；也就是说选修第二门课的情况下必能选修第一门课
反之则不成立，参考这个例子：days = 0， (t1, d1) = (2, 3), (t2, d2) = (2, 100)

基于以上结论，那么我们可以对所有课程按照结束时间递增排序，然后一一放入结果集合中
我们用一个变量 days 记录结果集中花费的总时间，对于第i门课，
如果 days + ti <= di，则将这门课加入结果集；
否则，看下结果集中耗时最长的课，是否可以替换为第i门课；显然仅判断耗时最长的课和第i门课谁更耗时即可决定
最后的结果即在结果中

为了迅速获知结果集中耗时最长的课，用优先队列即大顶堆即可

时间复杂度O(nlgn)——n为课程数目
空间复杂度O(n)
*/
type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *Heap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] <= courses[j][1]
	})
	pq := &Heap{}
	days := 0
	for _, c := range courses {
		if days+c[0] <= c[1] {
			heap.Push(pq, c[0])
			days += c[0]
		} else if pq.Len() > 0 {
			poped := heap.Pop(pq).(int)
			if poped > c[0] {
				heap.Push(pq, c[0])
				days = days - poped + c[0]
			} else {
				heap.Push(pq, poped)
			}
		}
	}
	return pq.Len()
}

/*
动态规划

主要思路： 记录课程数使用的最短时间
首先根据结束时间排序，如果结束时间相同，课时短的在前，这样计算下一个课的时候就不用考虑上一个课的结束时间
举例： 有3个课程 [9, 10], [3, 12], [7, 17]
初始化所有课数最短时间为maxInt,根据当前结束时间和总课长判断是否可以增加
[max, max, max]
[9, 10]: 只可以上一节课，即 [9, max, max]
[3, 12]: 可以上一节课，也可以上两节课，上两节课的时候为12; 因为3 < 9， 所以上一个节的时间最小为3 [3, 12, max]
[7, 17]: [3, 10, max]

时间复杂度O(n^2)
空间复杂度O(n)

作者：resara
链接：https://leetcode-cn.com/problems/course-schedule-iii/solution/golang-tan-xin-suan-fa-by-resara/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func scheduleCourse1(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		if courses[i][1] == courses[j][1] {
			return courses[i][0] < courses[j][0]
		}
		return courses[i][1] < courses[j][1]
	})
	max := 0
	dp := make([]int, len(courses))
	for i := 0; i < len(courses); i++ {
		t, d := courses[i][0], courses[i][1]
		dp[i] = math.MaxInt32
		for j := i; j >= 0; j-- { // 注意：这里需要从最后一节课开始计算，防止修改前面的对后面的造成影响
			tmp := t
			if j > 0 {
				tmp += dp[j-1]
			}
			if tmp <= d && tmp < dp[j] { // 如果这节课可以在结束时间上完，并且时长小于之前的，优化
				dp[j] = tmp
				if max < j+1 {
					max = j + 1
				}
			}
		}
	}
	return max
}
