/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package log_system

import "container/list"

/*
在一个简易日志系统中，每条日志有唯一的 ID，以及它的生成时间 timeStamp（ID 和 timeStamp 均为正整数）。
根据给出的函数框架，请实现以下功能（以 C 为例，其他语言参考）：

void Add(int id,int timeStamp)：增加一条新的日志记录，将这条日志记录存到系统中。
int Delete(int id)：在日志系统中尝试删除这个 ID 对应的日志记录。如果该日志 ID 在系统中不存在，返回 -1，否则删除这条日志，并返回 0。
int Query(int startTime,int endTime)：返回日志系统中生成时间大于等于 startTime 且小于等于 endTime 的日志数量。如果没有，返回 0。
注：同一时间可能有多条日志。函数定义以对应语言的右侧实际框架为准。


示例 1：

输入：
["LogSystem","add","add","add","query","delete","delete","query"]
[[],[1,5],[2,5],[3,6],[5,6],[2],[4],[5,6]]
输出：
[null,null,null,null,3,0,-1,2]
解释：
第一个操作是初始化，没有返回值；
前三个 Add 操作加入了 ID 为 1、2、3 的日志，timeStamp 分别为 5、5、6；
第一次 Query 操作查询 timeStamp 范围为[5,6]的日志数量，返回 3；
第一次 Delete 操作删除了 ID 为 2 的日志，删除成功，返回 0；
第二次 Delete 操作试图删除 ID 为 4 的日志，没有该日志，操作失败，返回 -1；
最后一次 Query 操作查询 timeStamp 范围为[5,6]的日志数量，由于该范围中已经有一条日志被删除了，故返回 2。
注：输出中的 null 表示此对应函数无输出（等同于 C 语言的 void 类型）。

限制：
1 <= id <= 10^9
1 <= timeStamp <= 10^9
每次 Add 调用中的 id 是唯一的
Add 最多调用 1000 次，Delete 最多调用 1000 次，Query 最多调用 50 次
*/
type Log struct {
	time, id int
}

type LogSystem struct {
	logs *list.List
	m    map[int]*list.Element
}

func Constructor() LogSystem {
	return LogSystem{logs: list.New(), m: map[int]*list.Element{}}
}

func (ls *LogSystem) add(id int, timeStamp int) {
	item := &Log{id: id, time: timeStamp}
	if ls.logs.Len() == 0 {
		ls.m[id] = ls.logs.PushBack(item)
		return
	}
	e := ls.logs.Front()
	for e != nil {
		if e.Value.(*Log).time >= timeStamp {
			break
		}
		e = e.Next()
	}
	if e != nil {
		ls.m[id] = ls.logs.InsertBefore(item, e)
	} else {
		ls.m[id] = ls.logs.PushBack(item)
	}
}

// 在日志系统中尝试删除这个 ID 对应的日志记录。如果该日志 ID 在系统中不存在，返回 -1，否则删除这条日志，并返回 0。
func (ls *LogSystem) delete(id int) int {
	if e, ok := ls.m[id]; ok {
		_ = ls.logs.Remove(e)
		delete(ls.m, id)
		return 0
	}
	return -1
}

// 返回日志系统中生成时间大于等于 startTime 且小于等于 endTime 的日志数量。如果没有，返回 0。
func (ls *LogSystem) query(startTime int, endTime int) int {
	if startTime > endTime {
		startTime, endTime = endTime, startTime
	}
	e := ls.logs.Front()
	for e != nil && e.Value.(*Log).time < startTime {
		e = e.Next()
	}
	count := 0
	for e != nil && e.Value.(*Log).time <= endTime {
		count++
		e = e.Next()
	}
	return count
}
