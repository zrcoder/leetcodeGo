/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package rearrange_k_distance_apart

import (
	"container/heap"
	"container/list"
	"sort"
)

/*
先统计每个任务的数量，贪心策略，需要先安排数量大的任务
以n+1个任务为一轮，同一轮中一个任务最多被安排一次。
每一轮中，将当前任务按照剩余次数降序排列，再选择剩余次数最多的n+1个任务一次执行
如果任务的种类 t 少于 n + 1 个，就只选择全部的 t 种任务，其余的时间空闲。

时间复杂度: O(result),给每个任务都安排了时间，因此时间复杂度和最终的答案成正比
空间复杂度: O(26)=O(1)
*/
func leastInterval(tasks []byte, n int) int {
	count := make([]int, 26)
	for _, v := range tasks {
		count[v-'A']++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(count)))
	result := 0
	for count[0] > 0 {
		for i := 0; i <= n && count[0] > 0; i++ {
			result++
			if i < 26 && count[i] > 0 {
				count[i]--
			}
		}
		sort.Sort(sort.Reverse(sort.IntSlice(count)))
	}
	return result
}

/*
继承自上面的方法，在选择每一轮任务时，可用堆代替排序。
一开始，把所有任务的数量加入堆。
每一轮，从堆里选择最多n+1个任务，把它们数量减去1，如果不为0，再重新放回堆中；直到堆为空

时空复杂度与上面一样
*/
func leastInterval1(tasks []byte, n int) int {
	h := prepareHeap(tasks)
	result := 0
	set := list.New()
	for h.Len() > 0 {
		for i := 0; i <= n; i++ {
			if h.Len() == 0 && set.Len() == 0 {
				return result
			}
			result++
			if h.Len() == 0 { // 需要待命只到i==n
				continue
			}
			t := heap.Pop(h).(int)
			if t > 1 {
				set.PushBack(t - 1)
			}
		}
		for set.Len() > 0 {
			heap.Push(h, set.Remove(set.Front()).(int))
		}
	}
	return result
}

func prepareHeap(tasks []byte) *Heap {
	count := make([]int, 26)
	for _, v := range tasks {
		count[v-'A']++
	}
	h := &Heap{}
	for _, v := range count {
		if v > 0 {
			h.Push(v)
		}
	}
	heap.Init(h)
	return h
}

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] > h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

/*
1.统计数量最大的任务的数量max；假设数组 ["A","A","A","B","B","C"]，n = 2，A的频率最高，记为max = 3，
两个A之间必须间隔2个任务，才能满足题意并且是最短时间（两个A的间隔大于2的总时间必然不是最短），
因此执行顺序为： A->X->X->A->X->X->A，这里的X表示除了A以外其他字母，或者是待命，不用关心具体是什么。
max个A，中间有 max - 1个间隔，每个间隔需要搭配n个X，再加上最后一个A，所以总时间为 (max - 1) * (n + 1) + 1
2.要注意可能会出现多个频率相同且都是最高的任务，比如 ["A","A","A","B","B","B","C","C"]，所以最后会剩下一个A和一个B，
因此最后要加上频率最高的不同任务的个数 maxCount
3.公式算出的值可能会比数组的长度小，如["A","A","B","B","D"]，n=1，遗漏了最后的 D（还可以有 E、F等），这种情况要取数组的长度

时间复杂度：O(M)，其中 M 是任务的总数，即tasks数组的长度。
空间复杂度：O(1)。
*/
func leastInterval3(tasks []byte, n int) int {
	// 统计每个任务的数量
	count := [26]int{}
	// 数量最大的任务的数量及个数
	max, maxCount := 0, 0
	for _, v := range tasks {
		c := count[v-'A'] + 1
		count[v-'A'] = c
		if max < c {
			max = c
			maxCount = 1
		} else if max == c {
			maxCount++
		}
	}
	result := (max-1)*(n+1) + maxCount
	if result < len(tasks) {
		result = len(tasks)
	}
	return result
}
