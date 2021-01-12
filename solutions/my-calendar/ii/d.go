/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package ii

import (
	"math"
)

type interval struct {
	start, end int
}

type MyCalendarTwo struct {
	calendar, overlap []interval // 分别表示已经添加的所有日程和已有日程重复的时间段组成的列表
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (mc *MyCalendarTwo) Book(start int, end int) bool {
	for _, val := range mc.overlap {
		if start < val.end && end > val.start {
			return false
		}
	}
	for _, val := range mc.calendar {
		if start < val.end && end > val.start {
			it := interval{start: max(start, val.start), end: min(end, val.end)}
			mc.overlap = append(mc.overlap, it)
		}
	}
	it := interval{start: start, end: end}
	mc.calendar = append(mc.calendar, it)
	return true
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
