/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package i

import (
	"sort"
)

type Interval struct {
	start, end int
}

type MyCalendar struct {
	calendar []Interval
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (mc *MyCalendar) Book(start int, end int) bool {
	pos := sort.Search(len(mc.calendar), func(i int) bool {
		return mc.calendar[i].start >= start
	})
	if pos < len(mc.calendar) && mc.calendar[pos].start < end ||
		pos-1 >= 0 && mc.calendar[pos-1].end > start {
		return false
	}
	it := Interval{start: start, end: end}
	mc.calendar = append(append(mc.calendar[:pos:pos], it), mc.calendar[pos:]...)
	return true
}
