package set_intersection_size_at_least_two

import (
	"container/list"
	"sort"
)

/*
757.  设置交集大小至少为2 https://leetcode-cn.com/problems/set-intersection-size-at-least-two

一个整数区间 [a, b]  ( a < b ) 代表着从 a 到 b 的所有连续整数，包括 a 和 b。

给你一组整数区间intervals，请找到一个最小的集合 S，
使得 S 里的元素与区间intervals中的每一个整数区间都至少有2个元素相交。

输出这个最小集合S的大小。

示例 1:

输入: intervals = [[1, 3], [1, 4], [2, 5], [3, 5]]
输出: 3
解释:
考虑集合 S = {2, 3, 4}. S与intervals中的四个区间都有至少2个相交的元素。
且这是S最小的情况，故我们输出3。
示例 2:

输入: intervals = [[1, 2], [2, 3], [2, 4], [4, 5]]
输出: 5
解释:
最小的集合S = {1, 2, 3, 4, 5}.
注意:

intervals 的长度范围为[1, 3000]。
intervals[i] 长度为 2，分别代表左、右边界。
intervals[i][j] 的值是 [0, 10^8]范围内的整数。
*/

/*
贪心策略：
先将所有区间按照起点降序排序，如果起点相同则终点升序排列——或者反过来，起点升序，起点相同时终点降序
这样排序的好处是在遍历的过程中只需要关注集合中最小的两个数字

遍历时根据当前区间[start, end]和交集中两个最小数字min1, min2的关系，分情况讨论：
1)、min1、min2都不在区间内，
集合中应该加入start和start+1两个数字，同时更新min1、min2为这两个数字

2)、min1、min2都在区间内，不需更新

3)、只有min1在区间内，
如果min1==start, 集合加入start+1，同时min2更新为start+1
否则，min2更新为min1， min1更新为start， 集合加入start

时间复杂度O(n*lgn + n) = O(nlgn), 主要为排序的复杂度，排序后只是一次遍历
空间复杂度为set集合的复杂度，set元素不会超过 2*n， 所以空间复杂度为O(n)
*/
func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] > intervals[j][0]
	})
	return help(intervals)
}

func help(intervals [][]int) int {
	min1, min2 := -1, -1
	set := list.New() // 用map或slice也行
	for _, v := range intervals {
		switch {
		case !isInRange(min1, v) && !isInRange(min2, v):
			set.PushBack(v[0])
			set.PushBack(v[1])
			min1, min2 = v[0], v[0]+1
		case isInRange(min1, v) && !isInRange(min2, v):
			if v[0] == min1 {
				set.PushBack(v[0] + 1)
				min2 = min1 + 1
			} else {
				set.PushBack(v[0])
				min1, min2 = v[0], min1
			}
		}
	}
	return set.Len()
}

// 这个题目求最终集合元素的个数，可以用一个int变量统计，不必开辟一个真正的集合去装元素，空间复杂度降为O(1)
func help1(intervals [][]int) int {
	min1, min2 := -1, -1
	result := 0
	for _, v := range intervals {
		switch {
		case !isInRange(min1, v) && !isInRange(min2, v):
			result += 2 // 集合中加入 v[0]和v[0]+1
			min1, min2 = v[0], v[0]+1
		case isInRange(min1, v) && !isInRange(min2, v):
			result++
			if v[0] == min1 {
				// 集合中加入v[0]+1
				min2 = min1 + 1
			} else {
				// 集合中加入v[0]
				min1, min2 = v[0], min1
			}
		}
	}
	return result
}

func isInRange(m int, interval []int) bool {
	return interval[0] <= m && m <= interval[1]
}
