/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package lexicographically_smallest_equivalent_string

func smallestEquivalentString(A string, B string, S string) string {
	if len(A) == 0 || len(A) != len(B) {
		return ""
	}
	const maxLetters = 26
	const firstLetter = 'a'
	uf := MakeUnionFind(maxLetters)
	for i := range A {
		uf.Union(int(A[i]-firstLetter), int(B[i]-firstLetter))
	}
	r := []byte(S)
	for i, v := range r {
		r[i] = byte(uf.Find(int(v-firstLetter))) + firstLetter
	}
	return string(r)
}

type UnionFind []int

func MakeUnionFind(n int) UnionFind {
	r := make([]int, n)
	for i := range r {
		r[i] = i
	}
	return r
}

func (uf UnionFind) Union(v1, v2 int) {
	c1 := uf.Find(v1)
	c2 := uf.Find(v2)
	c := min(c1, c2)
	uf[c1], uf[c2] = c, c // 不要按秩合并，让较小的节点做父节点
	uf[v1], uf[v2] = c, c
}

func (uf UnionFind) Find(v int) int {
	for v != uf[v] {
		v, uf[v] = uf[v], uf[uf[v]] // 路径压缩
	}
	return v
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
