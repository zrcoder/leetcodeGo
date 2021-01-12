/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package iii

import (
	"container/list"
)

type point struct {
	pos  int // 该点在数轴上的位置。
	deep int // 该点的深度——即被多少条线段包含。
}

type MyCalendarThree struct {
	points *list.List
	k      int
}

func Constructor() MyCalendarThree {
	mc := MyCalendarThree{points: list.New()}
	// 结合list特点, 方便后续处理，先预置两个点，无限小点和和无限大点
	// 注意输入的start和end在范围[0, 10^9]内
	mc.points.PushBack(&point{pos: -1, deep: 0})
	mc.points.PushBack(&point{pos: 1e9 + 1, deep: 0})
	return mc
}

func (mc *MyCalendarThree) Book(start int, end int) int {
	var startNode, endNode *list.Element
	// 插入起始点，如果已经存在则不插入
	for e := mc.points.Front(); e.Next() != nil; e = e.Next() {
		p := e.Value.(*point)
		if start == p.pos {
			startNode = e
			break
		}
		// 插入点，注意其深度暂时和其前驱点深度一致
		nextP := e.Next().Value.(*point)
		if start > p.pos && start < nextP.pos {
			p := &point{pos: start, deep: p.deep}
			startNode = mc.points.InsertAfter(p, e)
			break
		}
	}
	// 插入结束点，如果已经存在则不插入
	for e := mc.points.Back(); e.Prev() != nil; e = e.Prev() {
		p := e.Value.(*point)
		if end == p.pos {
			endNode = e
			break
		}
		prevP := e.Prev().Value.(*point)
		if end < p.pos && end > prevP.pos {
			p := &point{pos: end, deep: prevP.deep}
			endNode = mc.points.InsertBefore(p, e)
			break
		}
	}
	// 对于起始和结束点之间的所有点，深度都加一。
	for e := startNode; e != endNode && e != nil; e = e.Next() {
		p := e.Value.(*point)
		p.deep++
		mc.k = max(mc.k, p.deep)
	}
	return mc.k
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
