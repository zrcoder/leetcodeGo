package main

import (
	"container/heap"
	"fmt"
)

func main() {
	arr := make([]int, 0, 10)
	fmt.Println(arr, len(arr), cap(arr))

	for i := 0; i < 1025; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr, len(arr), cap(arr))
}

type Cmp func(int, int) bool

type Heap struct {
	slice   []int
	cmp     Cmp
	delMemo map[int]int
	size    int
}

func (h *Heap) Len() int           { return len(h.slice) }
func (h *Heap) Less(i, j int) bool { return h.cmp(h.slice[i], h.slice[j]) }
func (h *Heap) Swap(i, j int)      { h.slice[i], h.slice[j] = h.slice[j], h.slice[i] }
func (h *Heap) Push(v interface{}) { h.slice = append(h.slice, v.(int)) }
func (h *Heap) Pop() interface{}   { a := h.slice; v := a[len(a)-1]; h.slice = a[:len(a)-1]; return v }
func (h *Heap) push(v int)         { h.size++; heap.Push(h, v) }
func (h *Heap) pop() int {
	h.size--
	res := heap.Pop(h).(int)
	h.prune()
	return res
}
func (h *Heap) prune() {
	for h.Len() > 0 {
		num := h.slice[0]
		d, has := h.delMemo[num]
		if !has {
			break
		}
		if d > 1 {
			h.delMemo[num]--
		} else {
			delete(h.delMemo, num)
		}
		heap.Pop(h)
	}
}

var small, large *Heap

func medianSlidingWindow(nums []int, k int) []float64 {
	small = &Heap{delMemo: map[int]int{}} // 大根堆，维护较小的一半元素
	small.cmp = func(a int, b int) bool {
		return a > b
	}
	large = &Heap{delMemo: map[int]int{}} // 小根堆，维护较大的一半元素
	large.cmp = func(a int, b int) bool {
		return a < b
	}
	makeBalance := func() {
		// 调整 small 和 large 中的元素个数，使得二者的元素个数满足要求
		if small.size > large.size+1 { // small 比 large 元素多 2 个
			large.push(small.pop())
		} else if small.size < large.size { // large 比 small 元素多 1 个
			small.push(large.pop())
		}
	}
	insert := func(num int) {
		if small.Len() == 0 || num <= small.slice[0] {
			small.push(num)
		} else {
			large.push(num)
		}
		makeBalance()
	}
	erase := func(num int) {
		if num <= small.slice[0] {
			small.delMemo[num]++
			small.size--
			if num == small.slice[0] {
				small.prune()
			}
		} else {
			large.delMemo[num]++
			large.size--
			if num == large.slice[0] {
				large.prune()
			}
		}
		makeBalance()
	}
	getMedian := func() float64 {
		if k&1 > 0 {
			return float64(small.slice[0])
		}
		return float64(small.slice[0]+large.slice[0]) / 2
	}

	for _, num := range nums[:k] {
		insert(num)
	}
	n := len(nums)
	res := make([]float64, 0, n-k+1)
	res = append(res, getMedian())
	for i := k; i < n; i++ {
		insert(nums[i])
		erase(nums[i-k])
		res = append(res, getMedian())
	}
	return res
}
