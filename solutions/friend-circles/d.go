package circles

import "container/list"

/*
547. 朋友圈 https://leetcode-cn.com/problems/friend-circles/
班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。
如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是 C 的朋友。所谓的朋友圈，是指所有朋友的集合。

给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。
如果M[i][j] = 1，表示已知第 i 个和 j 个学生互为朋友关系，否则为不知道。
你必须输出所有学生中的已知的朋友圈总数。

示例 1:

输入:
[[1,1,0],
 [1,1,0],
 [0,0,1]]
输出: 2
说明：已知学生0和学生1互为朋友，他们在一个朋友圈。
第2个学生自己在一个朋友圈。所以返回2。
示例 2:

输入:
[[1,1,0],
 [1,1,1],
 [0,1,1]]
输出: 1
说明：已知学生0和学生1互为朋友，学生1和学生2互为朋友，所以学生0和学生2也是朋友，所以他们三个在一个朋友圈，返回1。
注意：

N 在[1,200]的范围内。
对于所有学生，有M[i][i] = 1。
如果有M[i][j] = 1，则有M[j][i] = 1。
*/
// bfs 时间复杂度O(n^2)， 空间复杂度O(n)
func findCircleNum(M [][]int) int {
	n := len(M)
	seen := make([]bool, n)
	q := list.New()
	count := 0
	for i := 0; i < n; i++ {
		if seen[i] {
			continue
		}
		q.PushBack(i)
		for q.Len() > 0 {
			k := q.Remove(q.Front()).(int)
			seen[k] = true
			for j := range M[k] {
				if M[k][j] == 0 || seen[j] {
					continue
				}
				q.PushBack(j)
			}
		}
		count++
	}
	return count
}

// union find, 时空复杂度同上
func findCircleNum1(M [][]int) int {
	n := len(M)
	uf := NewUnionFind(n)
	for i := range M {
		for j := range M[i] {
			if M[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}
	set := make(map[int]bool, 0)
	for _, v := range uf {
		set[uf.Find(v)] = true
	}
	return len(set)
}

type UF []int

func NewUnionFind(n int) UF {
	r := make([]int, n)
	for i := range r {
		r[i] = i
	}
	return r
}

func (uf UF) Union(i, j int) {
	uf[uf.Find(i)] = uf[uf.Find(j)]
}

func (uf UF) Find(i int) int {
	for i != uf[i] {
		i, uf[i] = uf[i], uf[uf[i]]
	}
	return i
}
