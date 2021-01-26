/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package matrix_traversal

import (
	"math"
	"sort"
)

/*
一、最直观的做法，起个名字：方形水波遍历~
从(r0, c0)出发，按照曼哈顿距离向外扩展，可以看到对于曼哈顿距离为定值的点，连接起来是一个旋转了45°的正方形，当然正方形的中心是（r0，c0）
可以将这些正方形形象地称作方形水波~
实际就是从(r0, c0)出发，从内向外层层遍历那些方形水波
对于每个方形水波，有上下左右4个顶点，
每次可以先从上顶点开始顺时针遍历，先遍历到右顶点，再遍历到下顶点，再到左顶点，最后回归上顶点；
当然这道题对顺序没有要求，可以从任意顶点开始，顺、逆时针遍历都行
注意方形水波的某些点可能超出矩阵范围，处理好边界即可

时空复杂度都是O(R*C)。
时间上，会有一些超出矩阵的点去遍历，但总的量级在R*C；
空间复杂度主要在最后的结果数组上，这是没法省的，除去结果数组，方形水波遍历的解法显然是常数级的空间复杂度
*/
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	result := make([][]int, R*C)
	result[0] = []int{r0, c0}
	step := 1
	r, c := r0, c0
	for step < R*C {
		r--          // 从上一次遍历后的方形水波上顶点再上一步到达这次需要遍历的方形水波上顶点
		for r < r0 { // 上顶点 -> 右顶点
			if r >= 0 && c <= C-1 {
				result[step] = []int{r, c}
				step++
			}
			r++
			c++
		}
		for c > c0 { // 右顶点 -> 下顶点
			if r <= R-1 && c <= C-1 {
				result[step] = []int{r, c}
				step++
			}
			r++
			c--
		}
		for r > r0 { // 下顶点 -> 左顶点
			if r <= R-1 && c >= 0 {
				result[step] = []int{r, c}
				step++
			}
			r--
			c--
		}
		for c < c0 { // 左顶点 -> 上顶点 // 这里的判断是c<c0, 保证了不会将上顶点重复放入结果
			if r >= 0 && c >= 0 {
				result[step] = []int{r, c}
				step++
			}
			r--
			c++
		}
	}
	return result
}

/*
二、借助哈希表
1.先用一个哈希表记录每个点到(r0, c0)的曼哈顿距离，以曼哈顿距离为key， 这些点的坐标组成的数组为value
2.再按照曼哈顿距离从小到大遍历哈希表即可
虽然Go里没有可排序的map，但对于这道题，其实比较容易实现上面的第2步;因为可以确定，曼哈顿距离是从0到最大曼哈顿距离的连续序列
距离(r0,c0)最大的曼哈顿距离很好求，就是矩阵4个顶点距离(r0,c0)的最大值；当然也可以在第1步确定这个最大值
时间复杂度O(R*C*2)=O(R*C)，空间复杂度O(R*C*2)=O(R*C)
*/
func allCellsDistOrder1(R int, C int, r0 int, c0 int) [][]int {
	m := make(map[int][][]int, 0)
	maxDist := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			dist := abs(r-r0) + abs(c-c0)
			if dist > maxDist {
				maxDist = dist
			}
			m[dist] = append(m[dist], []int{r, c})
		}
	}

	res := make([][]int, 0, R*C)
	for i := 0; i <= maxDist; i++ {
		for _, v := range m[i] {
			res = append(res, v)
		}
	}
	return res
}

/*
三、先将所有点放入结果，再对结果排序
时间复杂度O(R*C*lg(R*C))， 空间复杂度O(R*C)
*/
func allCellsDistOrder2(R int, C int, r0 int, c0 int) [][]int {
	res := make([][]int, 0, R*C)
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			res = append(res, []int{r, c})
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return abs(res[i][0]-r0)+abs(res[i][1]-c0) < abs(res[j][0]-r0)+abs(res[j][1]-c0)
	})
	return res
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

// 四、BFS
func allCellsDistOrder3(R int, C int, r0 int, c0 int) [][]int {
	res := make([][]int, 0, R*C)
	res = append(res, []int{r0, c0})
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	seen := make([][]bool, R)
	for i := range seen {
		seen[i] = make([]bool, C)
	}
	seen[r0][c0] = true
	k := 0
	for k < cap(res) {
		pos := res[k]
		row, column := pos[0], pos[1]
		for _, d := range dirs {
			r, c := row+d[0], column+d[1]
			if r < 0 || r >= R || c < 0 || c >= C || seen[r][c] {
				continue
			}
			seen[r][c] = true
			res = append(res, []int{r, c})
		}
		k++
	}
	return res
}
