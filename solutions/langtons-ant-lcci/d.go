/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package langtons_ant_lcci

/*
面试题 16.22. 兰顿蚂蚁
一只蚂蚁坐在由白色和黑色方格构成的无限网格上。开始时，网格全白，蚂蚁面向右侧。每行走一步，蚂蚁执行以下操作。

(1) 如果在白色方格上，则翻转方格的颜色，向右(顺时针)转 90 度，并向前移动一个单位。
(2) 如果在黑色方格上，则翻转方格的颜色，向左(逆时针方向)转 90 度，并向前移动一个单位。

编写程序来模拟蚂蚁执行的前 K 个动作，并返回最终的网格。

网格由数组表示，每个元素是一个字符串，代表网格中的一行，黑色方格由 'X' 表示，白色方格由 '_' 表示，
蚂蚁所在的位置由 'L', 'U', 'R', 'D' 表示，分别表示蚂蚁 左、上、右、下 的朝向。只需要返回能够包含蚂蚁走过的所有方格的最小矩形。

示例 1:

输入: 0
输出: ["R"]
示例 2:

输入: 2
输出:
[
  "_X",
  "LX"
]
示例 3:

输入: 5
输出:
[
  "_U",
  "X_",
  "XX"
]
说明：

K <= 100000
*/
/*
除了题目给的几个用例，还需要明确`K==1`的输出预期是：`["X","D"]`
理解题意后，模拟即可

借助一个哈希表，来装所有的点
假设一开始的点是`（0,0）`，在模拟过程中蚂蚁向左或向上走，点坐标可能变成负值
模拟完需要根据哈希表里的信息确定上下左右边界，从而确定结果矩阵的大小
并需要把所有点移动到第一象限，即坐标全为非负；只需要每个坐标的行`r`减去最小行边界，列`c`减去最小列边界

时间复杂度`O(K)`
*/

type Point struct {
	r, c int
}

func printKMoves(K int) []string {
	// 0, 1, 2, 3 分别代表 右、下、左、上四个方向
	currDir := 0
	// 当前的方向currDir可以通过+1（顺时针）或-1（+3， 逆时针）的方式得到下一个方向, 通过dirs能方便计算出下个坐标
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	currPoint := Point{r: 0, c: 0}
	set := make(map[Point]bool, 0) // 值为颜色:false白色，true黑色
	// 蚂蚁开跑，共K步
	for ; K > 0; K-- {
		// 根据currPoint的情况确定下一步方向并变色
		if set[currPoint] {
			currDir = (currDir + 3) % 4
			set[currPoint] = false
		} else {
			currDir = (currDir + 1) % 4
			set[currPoint] = true
		}
		// 走到下个点
		currPoint = Point{r: currPoint.r + dirs[currDir][0], c: currPoint.c + dirs[currDir][1]}
	}
	set[currPoint] = false // 最后一个格子，要保证加入set；不能翻转了，染成黑色或白色都行，最后被方向字母替换
	// 统计经过的点的边界
	var minR, maxR, minC, maxC int
	for p := range set {
		minR = min(minR, p.r)
		maxR = max(maxR, p.r)
		minC = min(minC, p.c)
		maxC = max(maxC, p.c)
	}
	m, n := maxR-minR+1, maxC-minC+1
	// 创建m*n的结果矩阵，初始化所有格子为白色
	result := make([][]byte, m)
	for i := range result {
		result[i] = make([]byte, n)
		for j := range result[i] {
			result[i][j] = '_'
		}
	}
	// 染黑应该是黑色的格子，最后一个格子需要标明方向
	for point, isBlack := range set {
		r, c := point.r-minR, point.c-minC
		if point.r == currPoint.r && point.c == currPoint.c {
			result[r][c] = []byte{'R', 'D', 'L', 'U'}[currDir]
		} else if isBlack {
			result[r][c] = 'X'
		}
	}
	return parse(result)
}

func parse(grid [][]byte) []string {
	r := make([]string, len(grid))
	for i, v := range grid {
		r[i] = string(v)
	}
	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
