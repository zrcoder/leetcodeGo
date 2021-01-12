/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package shortest_distance_to_target_color

import (
	"math"
	"sort"
)

/*
1182. 与目标颜色间的最短距离
https://leetcode-cn.com/problems/shortest-distance-to-target-color

给你一个数组 colors，里面有  1、2、 3 三种颜色。
我们需要在 colors 上进行一些查询操作 queries，其中每个待查项都由两个整数 i 和 c 组成。
现在请你帮忙设计一个算法，查找从索引 i 到具有目标颜色 c 的元素之间的最短距离。
如果不存在解决方案，请返回 -1。

示例 1：
输入：colors = [1,1,2,1,3,2,2,3,3], queries = [[1,3],[2,2],[6,1]]
输出：[3,0,3]
解释：
距离索引 1 最近的颜色 3 位于索引 4（距离为 3）。
距离索引 2 最近的颜色 2 就是它自己（距离为 0）。
距离索引 6 最近的颜色 1 位于索引 3（距离为 3）。

示例 2：
输入：colors = [1,2], queries = [[0,3]]
输出：[-1]
解释：colors 中没有颜色 3。

提示：
1 <= colors.length <= 5*10^4
1 <= colors[i] <= 3
1 <= queries.length <= 5*10^4
queries[i].length == 2
0 <= queries[i][0] < colors.length
1 <= queries[i][1] <= 3
*/

/*
朴素实现，超时
*/
func shortestDistanceColor1(colors []int, queries [][]int) []int {
	r := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		query := queries[i]
		r[i] = minDist1(colors, query[0], query[1])
	}
	return r
}

func minDist1(colors []int, index int, color int) int {
	if colors[index] == color {
		return 0
	}
	for i, j := index-1, index+1; i >= 0 || j < len(colors); i, j = i-1, j+1 {
		if i >= 0 && color == colors[i] {
			return index - i
		}
		if j < len(colors) && color == colors[j] {
			return j - index
		}
	}
	return -1
}

/*
使用一个哈希表colorMap，对于颜色c(1<=c<=3), colorMap[c]按升序保存colors里颜色为c的元素索引——这里颜色有限，也可以用三个切片代替哈希表
对于要搜索的索引和颜色，在colorMap[dstColor]里用二分搜索与srcIndex最接近的索引，计算处与srcIndex的距离即可
*/
func shortestDistanceColor(colors []int, queries [][]int) []int {
	colorMap := make(map[int][]int, 3)
	for index, color := range colors {
		colorMap[color] = append(colorMap[color], index)
	}
	r := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		r[i] = minDist(queries[i], colorMap)
	}
	return r
}

func minDist(query []int, m map[int][]int) int {
	dstIndex, dstColor := query[0], query[1]
	indexes := m[dstColor]
	if len(indexes) == 0 {
		return -1
	}
	i := sort.SearchInts(indexes, dstIndex)
	if i == len(indexes) {
		return dstIndex - indexes[i-1]
	}
	if i == 0 {
		return indexes[0] - dstIndex
	}
	return min(abs(indexes[i]-dstIndex), abs(indexes[i-1]-dstIndex))
}
func minDist2(query []int, m map[int][]int) int {
	dstIndex, dstColor := query[0], query[1]
	indexes := m[dstColor]
	if len(indexes) == 0 {
		return -1
	}
	left, right := 0, len(indexes)
	for left < right {
		mid := left + (right-left)/2
		if indexes[mid] == dstIndex {
			return 0
		}
		if indexes[mid] < dstIndex {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left == len(indexes) {
		return dstIndex - indexes[left-1]
	}
	if left == 0 {
		return indexes[0] - dstIndex
	}
	return min(abs(indexes[left]-dstIndex), abs(indexes[left-1]-dstIndex))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
