/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package campus_bikes

import (
	"container/heap"
	"math"
	"sort"
)

/*
1057. 校园自行车分配 https://leetcode-cn.com/problems/campus-bikes
在由 2D 网格表示的校园里有 n 位工人（worker）和 m 辆自行车（bike），n <= m。所有工人和自行车的位置都用网格上的 2D 坐标表示。

我们需要为每位工人分配一辆自行车。在所有可用的自行车和工人中，
我们选取彼此之间曼哈顿距离最短的工人自行车对  (worker, bike) ，并将其中的自行车分配給工人。
如果有多个 (worker, bike) 对之间的曼哈顿距离相同，那么我们选择工人索引最小的那对。
类似地，如果有多种不同的分配方法，则选择自行车索引最小的一对。
不断重复这一过程，直到所有工人都分配到自行车为止。

给定两点 p1 和 p2 之间的曼哈顿距离为 Manhattan(p1, p2) = |p1.x - p2.x| + |p1.y - p2.y|。

返回长度为 n 的向量 ans，其中 a[i] 是第 i 位工人分配到的自行车的索引（从 0 开始）。
*/

/*
1。计算并记录每个工人和每辆自行车之间的距离，可以用一个m*n大小的数组来记录，之后按照题目要求排序
2。再遍历数组一一确定结果，注意要有两个map（或数组）来记录当前工人或自行车是否已经放入过结果

时间复杂度O(n*m*lg(n*m))空间复杂度O(n*m)
另一个类似的实现是用小顶堆替换数组，实现见d2.go；时空复杂度与这个解法一样，但实际测试花费的时间是直接用数组的约两倍
*/

type item struct {
	workerId int
	bikeId   int
	dist     int
}

func assignBikes(workers [][]int, bikes [][]int) []int {
	n, m := len(workers), len(bikes)
	if n == 0 || n > m {
		return nil
	}
	type item struct {
		workerId int
		bikeId   int
		dist     int
	}
	items := make([]*item, n*m)
	k := 0
	for i, worker := range workers {
		for j, bike := range bikes {
			items[k] = &item{workerId: i, bikeId: j, dist: dist(worker[0], worker[1], bike[0], bike[1])}
			k++
		}
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].dist == items[j].dist {
			if items[i].workerId == items[j].workerId {
				return items[i].bikeId < items[j].bikeId
			}
			return items[i].workerId < items[j].workerId
		}
		return items[i].dist < items[j].dist
	})
	workerUsed := make([]bool, n)
	bikeUsed := make([]bool, m)
	result := make([]int, n)
	for _, item := range items {
		if workerUsed[item.workerId] || bikeUsed[item.bikeId] {
			continue
		}
		result[item.workerId] = item.bikeId
		workerUsed[item.workerId] = true
		bikeUsed[item.bikeId] = true
	}
	return result
}

func dist(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

type Heap []*item

func (h Heap) Len() int      { return len(h) }
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool {
	if h[i].dist == h[j].dist {
		if h[i].workerId == h[j].workerId {
			return h[i].bikeId < h[j].bikeId
		}
		return h[i].workerId < h[j].workerId
	}
	return h[i].dist < h[j].dist
}
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*item))
}
func (h *Heap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func assignBikes1(workers [][]int, bikes [][]int) []int {
	n, m := len(workers), len(bikes)
	if n == 0 || n > m {
		return nil
	}
	var pq Heap = make([]*item, 0, n*m)
	for i, worker := range workers {
		for j, bike := range bikes {
			heap.Push(&pq, &item{workerId: i, bikeId: j, dist: dist(worker[0], worker[1], bike[0], bike[1])})
		}
	}
	workerUsed := make([]bool, n)
	bikeUsed := make([]bool, m)
	result := make([]int, n)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*item)
		if workerUsed[item.workerId] || bikeUsed[item.bikeId] {
			continue
		}
		result[item.workerId] = item.bikeId
		workerUsed[item.workerId] = true
		bikeUsed[item.bikeId] = true
	}
	return result
}
