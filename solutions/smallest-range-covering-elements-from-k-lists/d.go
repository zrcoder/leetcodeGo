package smallest_range_covering_elements_from_k_lists

import (
	"container/heap"
	"math"
)

/*
解法一： 借助堆的贪心解法
问题可以转化为，从 k 个列表中各取一个数，使得这 k 个数中的最大值与最小值的差最小。

每个列表都是有序的，可以用贪心策略来解决。以题目中给的例子为例，每次在数字前边加个“.”表示数字被选中：

- 1.选择每个列表的首元素

.4  10  15  24  26
.0  9   12  20
.5  18  22  30

max：5， min：0， abs：5-0=5

- 2.移除最小元素，选中其下一个

.4  10  15  24  26
0   .9  12  20
.5  18  22  30

max：9， min：4， abs：9-4=5

- 3.继续如上操作

4   .10 15  24  26
0   .9  12  20
.5  18  22  30

max：10， min：5， abs：10-5=5

- 4.继续如上操作

4   .10 15  24  26
0   .9  12  20
5   .18 22  30

max：18， min：9， abs：18-9=9

- 5.继续如上操作

4   .10 15  24  26
0   9   .12 20
5   .18 22  30

max：18， min：10， abs：18-10=8

- 6.继续如上操作

4   10  .15 24  26
0   9   .12 20
5   .18 22  30

max：18， min：12， abs：18-12=6

- 7.继续如上操作

4   10  .15 24  26
0   9   12  .20
5   .18 22  30

max：20， min：15， abs：20-15=5

- 8.继续如上操作

4   10  15  .24  26
0   9   12  .20
5   .18 22  30

max：24， min：18， abs：24-18=6

- 9.继续如上操作

4   10  15  .24  26
0   9   12  .20
5   18  .22 30

max：24， min：20， abs：24-20=4

最小值所在列表元素已经尝试完，整个流程结束
以上所有步骤中，abs最小的min和max即为所求

如果每次遍历k次找到k个元素的最小值和最大值，时间复杂度会比较大
可以借助一个大小为k的小顶堆h来记录每次尝试的k个数字，一个变量 curMax 记录k个数字里最大的元素
每次计算 curMax 和堆顶元素的差得到 abs，之后堆顶元素修改成其所在列表的下一个元素
直到某一个列表元素尝试完

假设所有数字共n个，则时间复杂度为(n*lgk)
空间复杂度O(k)，堆的大小
*/

// 记录每个列表当前尝试元素的索引
var pos []int

func smallestRange(nums [][]int) []int {
	size := len(nums)
	pos = make([]int, size)

	h := &Heap{s: make([]Item, 0, size)}
	curMax := math.MinInt32
	// 先处理每个列表首元素
	for i, v := range nums {
		heap.Push(h, Item{val: v[0], listIndex: i})
		curMax = max(curMax, v[0])
	}
	left, right, minAbs := math.MinInt32, math.MaxInt32, math.MaxInt32
	for {
		peek := h.s[0] // 堆顶
		abs := curMax - peek.val
		if abs < minAbs {
			minAbs = abs
			left = peek.val
			right = curMax
		}
		index := peek.listIndex
		curPos := pos[index] + 1
		if curPos == len(nums[index]) {
			break
		}
		pos[index] = curPos
		curNum := nums[index][curPos]
		h.s[0].val = curNum
		heap.Fix(h, 0)
		curMax = max(curMax, curNum)
	}
	return []int{left, right}
}

type Item struct {
	val, listIndex int // 记录每个数字及其所在列表的索引
}

type Heap struct {
	s []Item
}

func (h *Heap) Len() int           { return len(h.s) }
func (h *Heap) Swap(i, j int)      { h.s[i], h.s[j] = h.s[j], h.s[i] }
func (h *Heap) Less(i, j int) bool { return h.s[i].val < h.s[j].val }
func (h *Heap) Push(x interface{}) { h.s = append(h.s, x.(Item)) }
func (h *Heap) Pop() interface{} {
	x := h.s[len(h.s)-1]
	h.s = h.s[:len(h.s)-1]
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
