/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package intheap

import "container/heap"

type Heap struct {
	slice []int
	cmp   func(i, j int) bool
}

func (h *Heap) InitWithCmp(cmp func(i, j int) bool) {
	h.cmp = cmp
}

func (h Heap) Len() int            { return len(h.slice) }
func (h Heap) Less(i, j int) bool  { return h.cmp(i, j) }
func (h Heap) Swap(i, j int)       { h.slice[i], h.slice[j] = h.slice[j], h.slice[i] }
func (h *Heap) Peek() int          { return h.slice[0] }
func (h *Heap) Push(x interface{}) { h.slice = append(h.slice, x.(int)) }
func (h *Heap) Pop() interface{} {
	x := h.slice[len(h.slice)-1]
	h.slice = (h.slice)[:len(h.slice)-1]
	return x
}
func (h *Heap) Get(i int) int {
	return h.slice[i]
}
func (h *Heap) Remove(v int) {
	for i, n := range h.slice {
		if v == n {
			_ = heap.Remove(h, i)
			return
		}
	}
}
