package stream

import "container/heap"

/*
295. 数据流的中位数
https://leetcode-cn.com/problems/find-median-from-data-stream

中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。

例如，

[2,3,4] 的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。
示例：

addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3)
findMedian() -> 2
进阶:

如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
通过次数19,864提交次数42,041
*/
/*
朴素实现
维护一个有序的切片，这样可以在常数级找到中位数，但是插入需要O(n) 的复杂度

type MedianFinder struct {
    s []int
}

func Constructor() MedianFinder {
    return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int)  {
    mf.s = append(mf.s, num)
    i := len(mf.s)-2
    for ; i >= 0 && mf.s[i] > num ; i-- {
        mf.s[i+1] = mf.s[i]
    }
    mf.s[i+1] = num
}

func (mf *MedianFinder) FindMedian() float64 {
    n := len(mf.s)
    if n % 2 == 1 {
        return float64(mf.s[n/2])
    }
    return float64(mf.s[n/2-1]+mf.s[n/2])*0.5
}
*/
/*
用堆来改进：
一个大顶堆保存数组中较小的一半元素
一个小顶堆保存数组中较大的一半元素
可保持大顶堆的元素和小顶堆相等或比小顶堆多一个
这样查找中位数还是常数级的复杂度，添加元素的复杂度将为 O(lgn)
*/
type MedianFinder struct {
	small, large *Heap
}

func Constructor() MedianFinder {
	res := MedianFinder{&Heap{}, &Heap{}}
	res.small.cmp = func(i, j int) bool {
		return res.small.s[i] > res.small.s[j]
	}
	res.large.cmp = func(i, j int) bool {
		return res.large.s[i] < res.large.s[j]
	}
	return res
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.small.Len() == mf.large.Len() {
		mf.small.push(num)
	} else {
		mf.large.push(num)
	}
	if mf.large.Len() > 0 && mf.small.s[0] > mf.large.s[0] {
		mf.large.push(mf.small.pop())
		mf.small.push(mf.large.pop())
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.small.Len() > mf.large.Len() {
		return float64(mf.small.s[0])
	}
	if mf.large.Len() == 0 && mf.small.Len() == 0 {
		return 0
	}
	return float64(mf.small.s[0]+mf.large.s[0]) * 0.5
}

type Heap struct {
	s   []int
	cmp func(i, j int) bool
}

func (h *Heap) Len() int           { return len(h.s) }
func (h *Heap) Less(i, j int) bool { return h.cmp(i, j) }
func (h *Heap) Swap(i, j int)      { h.s[i], h.s[j] = h.s[j], h.s[i] }
func (h *Heap) Push(x interface{}) { h.s = append(h.s, x.(int)) }
func (h *Heap) Pop() interface{} {
	x := h.s[h.Len()-1]
	h.s = h.s[:h.Len()-1]
	return x
}
func (h *Heap) push(x int) { heap.Push(h, x) }
func (h *Heap) pop() int   { return heap.Pop(h).(int) }
