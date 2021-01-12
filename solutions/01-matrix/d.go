package _1_matrix

import "container/list"

/*
542. 01 矩阵
给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。

两个相邻元素间的距离为 1 。

示例 1:
输入:

0 0 0
0 1 0
0 0 0
输出:

0 0 0
0 1 0
0 0 0
示例 2:
输入:

0 0 0
0 1 0
1 1 1
输出:

0 0 0
0 1 0
1 2 1
注意:

给定矩阵的元素个数不超过 10000。
给定矩阵中至少有一个元素是 0。
矩阵中的元素只在四个方向上相邻: 上、下、左、右。
*/
/*
假设整个矩阵就1个0，应该怎么做？
从元素0那个位置开始向四周不断扩散，修改四周位置的值，第一层为1，第二层为2，。。。
也就是bfs
如果矩阵中有若干个0，和只有一个0类似，还是可以用bfs
bfs中用队列，保证一层一层遍历
另有个小技巧，可以在第一次遍历矩阵的时候将非0元素标记为-1，方便后边bfs时判断某些位置是否已经访问过，省去visited数组
*/
func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}
	q := list.New()
	calZeros(matrix, q)
	bfs(matrix, q)
	return matrix
}

func calZeros(matrix [][]int, q *list.List) {
	for r := range matrix {
		for c, v := range matrix[r] {
			if v == 0 {
				q.PushBack([]int{r, c})
			} else {
				matrix[r][c] = -1 // 在后续bfs过程中方便判断某位置是否访问过
			}
		}
	}
}

func bfs(matrix [][]int, q *list.List) {
	m, n := len(matrix), len(matrix[0])
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for q.Len() > 0 {
		pos := q.Remove(q.Front()).([]int)
		for _, d := range dirs {
			r, c := pos[0]+d[0], pos[1]+d[1]
			if isInRange(r, c, m, n) && matrix[r][c] == -1 {
				matrix[r][c] = matrix[pos[0]][pos[1]] + 1
				q.PushBack([]int{r, c})
			}
		}
	}
}

func isInRange(r, c, m, n int) bool {
	return r >= 0 && r < m && c >= 0 && c < n
}
