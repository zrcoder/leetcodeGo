/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package design1

type Heap []*TweetInfo

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].Time < h[j].Time }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*TweetInfo))
}
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NewPriorityQueue() *Heap {
	return &Heap{}
}
