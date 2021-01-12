/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package connecting_cities_with_minimum_cost

import "sort"

func minimumCost(n int, connections [][]int) int {
	if len(connections) < n-1 { // 要有每个城市的联接信息，最终才能将所有城市联通，否则总有落单的
		return -1
	}
	sort.Slice(connections, func(i, j int) bool { // 按照成本排序
		return connections[i][2] < connections[j][2]
	})
	unionFind := NewUnionFind(n)
	connected, result, i := 0, 0, 0
	for connected < n-1 {
		connection := connections[i]
		i++
		city1, city2 := connection[0]-1, connection[1]-1
		city1, city2 = unionFind.Find(city1), unionFind.Find(city2)
		if city1 != city2 {
			unionFind.Union(city1, city2)
			connected++
			result += connection[2]
		}
	}
	return result
}

type UnionFind []int

func NewUnionFind(n int) UnionFind {
	unionFind := make([]int, n)
	for i := range unionFind {
		unionFind[i] = i
	}
	return unionFind
}
func (uf UnionFind) Find(x int) int {
	for uf[x] != x {
		x, uf[x] = uf[x], uf[uf[x]]
	}
	return x
}
func (uf UnionFind) Union(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	uf[rootX] = rootY // 可以按秩合并，即高度较小的根插入高度较大的根下面，进一步减少整个Union、Find操作的复杂度
}
