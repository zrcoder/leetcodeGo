/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package optimize_water_distribution_in_a_village

import "sort"

func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
	for i, cost := range wells {
		house := i + 1
		pipes = append(pipes, []int{0, house, cost})
	}
	sort.Slice(pipes, func(i, j int) bool {
		return pipes[i][2] < pipes[j][2]
	})
	connected, i, result := 0, 0, 0
	uf := NewUnionFind(n + 1)
	for connected < n {
		pipe := pipes[i]
		i++
		house1, house2, cost := pipe[0], pipe[1], pipe[2]
		house1, house2 = uf.Find(house1), uf.Find(house2)
		if house1 != house2 {
			uf.Union(house1, house2)
			result += cost
			connected++
		}
	}
	return result
}

type UnionFind []int

func NewUnionFind(n int) UnionFind {
	uf := make([]int, n)
	for i := range uf {
		uf[i] = i
	}
	return uf
}
func (uf UnionFind) Find(x int) int {
	root := x
	for root != uf[root] {
		root = uf[root]
	}
	for x != root {
		x, uf[x] = uf[x], root
	}
	return root
}
func (uf UnionFind) Union(x, y int) {
	uf[x] = y
}
