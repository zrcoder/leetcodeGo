/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package path_with_maximum_minimum_value

import (
	"container/heap"
)

/*
1102. 得分最高的路径 https://leetcode-cn.com/problems/path-with-maximum-minimum-value
给你一个 R 行 C 列的整数矩阵 A。矩阵上的路径从 [0,0] 开始，在 [R-1,C-1] 结束。
路径沿四个基本方向（上、下、左、右）展开，从一个已访问单元格移动到任一相邻的未访问单元格。
路径的得分是该路径上的 最小 值。例如，路径 8 →  4 →  5 →  9 的值为 4 。
找出所有路径中得分 最高 的那条路径，返回其 得分。

示例 1：
输入：[[5,4,5],[1,2,6],[7,4,6]]
输出：4
解释：
得分最高的路径用黄色突出显示。

示例 2：
输入：[[2,2,1,2,2,2],[1,2,2,2,1,2]]
输出：2

示例 3：
输入：[[3,4,6,3,4],[0,2,1,1,7],[8,8,3,2,7],[3,2,4,9,8],[4,1,2,0,0],[4,6,5,4,3]]
输出：3
提示：
1 <= R, C <= 100
0 <= A[i][j] <= 10^9
*/
/*
和[778] 水位上升的泳池中游泳这个问题非常类似
第一个想法，借助大顶堆的广度优先搜索，每次选值较大的格子，到达终点后，经过的格子里的最小值即为所得
*/
type Info struct {
	r, c, v int // 行、列、值
}
type Heap []Info

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].v > h[j].v }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(Info)) }
func (h *Heap) Pop() interface{} {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

func maximumMinimumPath(A [][]int) int {
	m, n := len(A), len(A[0])
	pq := &Heap{}
	const maxSize = 100
	visited := [maxSize][maxSize]bool{}
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	heap.Push(pq, Info{r: 0, c: 0, v: A[0][0]})
	result := A[0][0]
	visited[0][0] = true
	for pq.Len() > 0 {
		info := heap.Pop(pq).(Info)
		if info.v < result {
			result = info.v
		}
		if info.r == m-1 && info.c == n-1 {
			return result
		}
		for _, d := range dirs {
			r, c := info.r+d[0], info.c+d[1]
			if r >= 0 && r < m && c >= 0 && c < n && !visited[r][c] {
				heap.Push(pq, Info{r: r, c: c, v: A[r][c]})
				visited[r][c] = true
			}
		}
	}
	return 0
}

/*
题目明确约束数组里的元素非负，同时可以确定最终的结果不会大于min(起点值, 终点值),为方便叙述，称为limited
即结果在[0, limited]区间里，
可以从limited开始递减，看看是否在从起点到终点的某条路径上（用dfs更方便），路径上的所有值都不小于尝试的值
用二分法将更快
*/
func maximumMinimumPath1(A [][]int) int {
	m, n := len(A), len(A[0])
	const maxSize = 100
	var visited [maxSize][maxSize]bool
	var isValid func(r, c, lo int) bool
	isValid = func(r, c, lo int) bool { // dfs
		if r < 0 || r >= m || c < 0 || c >= n || A[r][c] < lo || visited[r][c] {
			return false
		}
		if r == m-1 && c == n-1 {
			return true
		}
		visited[r][c] = true
		return isValid(r+1, c, lo) || isValid(r-1, c, lo) || isValid(r, c+1, lo) || isValid(r, c-1, lo)
	}
	limited := A[0][0]
	if A[m-1][n-1] < limited {
		limited = A[m-1][n-1]
	}
	return searchFromHiToLow(limited+1, func(i int) bool {
		visited = [maxSize][maxSize]bool{}
		return isValid(0, 0, i)
	})
}

func searchFromHiToLow(n int, f func(i int) bool) int {
	lo, hi := 0, n
	for lo < hi {
		mid := lo + (hi-lo)/2
		if f(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo - 1
}
