/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package exam_room

import "container/list"

type ExamRoom struct {
	seated *list.List // 装坐着同学的位置
	last   int        // 最后一个座位， 就是总座位数减一
}

func Constructor(N int) ExamRoom {
	return ExamRoom{
		seated: list.New(),
		last:   N - 1,
	}
}

func (room *ExamRoom) Seat() int {
	if room.seated.Len() == 0 { // 还没有人入座，选座位0
		room.seated.PushFront(0)
		return 0
	}
	prevSeated := room.seated.Front().Value.(int)
	targetVal := 0                           // 需要插入的座位
	maxDist := prevSeated                    // 入座后距离最近的人的最大距离，当前是从位置0到第一个坐了人的位置的距离
	targetNextElement := room.seated.Front() // 需要插入的点的后一个元素。方便找到后直接插入
	for e := room.seated.Front().Next(); e != nil; e = e.Next() {
		currSeated := e.Value.(int)
		distant := (currSeated - prevSeated) / 2 // 两点之间的最远距离
		if distant > maxDist {
			maxDist = distant
			targetNextElement = e
			targetVal = prevSeated + distant
		}
		prevSeated = currSeated
	}
	if room.last-prevSeated > maxDist { // 尾部特殊判断
		room.seated.PushBack(room.last)
		return room.last
	}
	room.seated.InsertBefore(targetVal, targetNextElement)
	return targetVal
}

func (room *ExamRoom) Leave(p int) {
	for e := room.seated.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == p {
			room.seated.Remove(e)
			return
		}
	}
	return
}
