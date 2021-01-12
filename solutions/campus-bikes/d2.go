/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package campus_bikes

import "container/heap"

type item struct {
	workerId int
	bikeId   int
	dist     int
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
