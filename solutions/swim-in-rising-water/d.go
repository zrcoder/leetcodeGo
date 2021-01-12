/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package swim_in_rising_water

import (
	"container/heap"
	"container/list"
	"sort"

	s "github.com/zrcoder/leetcodeGo/util/sort"
)

/*
不用二分法的朴素实现，时间复杂度O(n^2*(max-grid[0][0]+1))
leetcode实测会花费376 ms，其他解法的时间在12-28ms内
*/
func swimInWater0(grid [][]int) int {
	start, end := grid[0][0], max(grid)
	for i := start; i < end; i++ {
		if canReach(i, grid) {
			return i
		}
	}
	return end
}

/*
二分法
*/
func swimInWater1(grid [][]int) int {
	return s.Search(grid[0][0], max(grid)+1, func(i int) bool {
		return canReach(i, grid)
	})
}

// 二分法也可用标准库
func swimInWater2(grid [][]int) int {
	return sort.Search(max(grid)+1, func(i int) bool { // 这里其实有点浪费，在[0,max]的区间里搜所的
		if i < grid[0][0] {
			return false
		}
		return canReach(i, grid)
	})
}

func max(grid [][]int) int {
	result := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid); c++ {
			if grid[r][c] > result {
				result = grid[r][c]
			}
		}
	}
	return result
}

func canReach(t int, grid [][]int) bool {
	const maxN = 50
	n := len(grid)
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	visited := [maxN][maxN]bool{}
	visited[0][0] = true
	queue := list.New()
	queue.PushBack([]int{0, 0})
	for queue.Len() > 0 {
		pos := queue.Remove(queue.Front()).([]int)
		r, c := pos[0], pos[1]
		if r == n-1 && c == n-1 {
			return true
		}
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < n && nc >= 0 && nc < n &&
				!visited[nr][nc] && grid[nr][nc] <= t {
				queue.PushBack([]int{nr, nc})
				visited[nr][nc] = true
			}
		}
	}
	return false
}

func canReach1(t int, grid [][]int) bool {
	const maxN = 50
	n := len(grid)
	visited := [maxN][maxN]bool{}
	var dfs func(r, c int) bool
	dfs = func(r, c int) bool {
		if r < 0 || c < 0 || r >= n || c >= n ||
			visited[r][c] || grid[r][c] > t {
			return false
		}
		if r == n-1 && c == n-1 {
			return true
		}
		visited[r][c] = true
		return dfs(r+1, c) || dfs(r-1, c) || dfs(r, c+1) || dfs(r, c-1)
	}
	return dfs(0, 0)
}

/*
解法二： 借助小顶堆的广度优先搜索
*/

// 平台结构体，方便自定义heap实现
type pos struct {
	height, r, c int // 高度和横纵坐标
}

func swimInWater(grid [][]int) int {
	const maxN = 50
	n := len(grid)
	dr := []int{1, -1, 0, 0}
	dc := []int{0, 0, 1, -1}
	visited := [maxN][maxN]bool{}
	result := 0
	pq := &posHeap{}
	heap.Push(pq, pos{height: grid[0][0], r: 0, c: 0})

	for pq.Len() > 0 {
		info := heap.Pop(pq).(pos) // 游到当前最低的平台上
		if grid[info.r][info.c] > result {
			result = grid[info.r][info.c]
		}
		if info.r == n-1 && info.c == n-1 { // 终点
			break
		}
		for i := 0; i < len(dr); i++ {
			r, c := info.r+dr[i], info.c+dc[i]
			if r >= 0 && r < n && c >= 0 && c < n && !visited[r][c] {
				heap.Push(pq, pos{height: grid[r][c], r: r, c: c})
				visited[r][c] = true
			}
		}
	}
	return result
}

type posHeap []pos

func (h posHeap) Len() int            { return len(h) }
func (h posHeap) Less(i, j int) bool  { return h[i].height < h[j].height }
func (h posHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *posHeap) Push(x interface{}) { *h = append(*h, x.(pos)) }
func (h *posHeap) Pop() interface{} {
	pos := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return pos
}
