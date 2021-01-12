/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */
package push_box

/*
推箱子」是一款风靡全球的益智小游戏，玩家需要将箱子推到仓库中的目标位置。

游戏地图用大小为 n * m 的网格 grid 表示，其中每个元素可以是墙、地板或者是箱子。

现在你将作为玩家参与游戏，按规则将箱子 'B' 移动到目标位置 'T' ：

玩家用字符 'S' 表示，只要他在地板上，就可以在网格中向上、下、左、右四个方向移动。
地板用字符 '.' 表示，意味着可以自由行走。
墙用字符 '#' 表示，意味着障碍物，不能通行。
箱子仅有一个，用字符 'B' 表示。相应地，网格上有一个目标位置 'T'。
玩家需要站在箱子旁边，然后沿着箱子的方向进行移动，此时箱子会被移动到相邻的地板单元格。记作一次「推动」。
玩家无法越过箱子。
返回将箱子推到目标位置的最小 推动 次数，如果无法做到，请返回 -1。

示例 1：
输入：grid = [["#","#","#","#","#","#"],
             ["#","T","#","#","#","#"],
             ["#",".",".","B",".","#"],
             ["#",".","#","#",".","#"],
             ["#",".",".",".","S","#"],
             ["#","#","#","#","#","#"]]
输出：3
解释：我们只需要返回推箱子的次数。

示例 2：
输入：grid = [["#","#","#","#","#","#"],
             ["#","T","#","#","#","#"],
             ["#",".",".","B",".","#"],
             ["#","#","#","#",".","#"],
             ["#",".",".",".","S","#"],
             ["#","#","#","#","#","#"]]
输出：-1

示例 3：
输入：grid = [["#","#","#","#","#","#"],
             ["#","T",".",".","#","#"],
             ["#",".","#","B",".","#"],
             ["#",".",".",".",".","#"],
             ["#",".",".",".","S","#"],
             ["#","#","#","#","#","#"]]
输出：5
解释：向下、向左、向左、向上再向上。

示例 4：
输入：grid = [["#","#","#","#","#","#","#"],
             ["#","S","#",".","B","T","#"],
             ["#","#","#","#","#","#","#"]]
输出：-1

提示：
1 <= grid.length <= 20
1 <= grid[i].length <= 20
grid 仅包含字符 '.', '#',  'S' , 'T', 以及 'B'。
grid 中 'S', 'B' 和 'T' 各只能出现一个。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-moves-to-move-a-box-to-their-target-location
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
1.先简化下问题：如果没有箱子，人最短移动多少次到达目的地？
这样就是一个比较典型的最短路径问题，可以用广度优先遍历来解决
可以参考下边canMoveTo函数的实现
2.对于这个问题，还需要判断人能不能到达箱子后边去推它，可以用广度优先遍历，也可用深度优先遍历
用于广度优先遍历的队列里应该同时存放人和箱子的坐标
*/
func minPushBox(grid [][]byte) int {
	box, person, boxTarget := getInitialPositions(grid)
	result := 0
	const maxSize = 20
	visited := [maxSize][maxSize][maxSize][maxSize]bool{}
	queue := [][]int{{box[0], box[1], person[0], person[1]}}
	visited[box[0]][box[1]][person[0]][person[1]] = true
	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	for len(queue) > 0 {
		length := len(queue)
		for k := 0; k < length; k++ {
			current := queue[0]
			queue = queue[1:]
			if current[0] == boxTarget[0] && current[1] == boxTarget[1] {
				return result
			}
			box := []int{current[0], current[1]}
			person := []int{current[2], current[3]}
			for _, d := range dirs {
				boxTarget := []int{current[0] + d[0], current[1] + d[1]}    // 箱子将要到达的位置
				personTarget := []int{current[0] - d[0], current[1] - d[1]} // 人将要到达的位置，与箱子要去的位置对称
				if isValid(boxTarget, grid) &&
					!visited[boxTarget[0]][boxTarget[1]][current[0]][current[1]] &&
					isValid(personTarget, grid) &&
					canMoveTo(personTarget, box, person, grid) {
					queue = append(queue, []int{boxTarget[0], boxTarget[1], current[0], current[1]}) // person 到达当前box位置
					visited[boxTarget[0]][boxTarget[1]][current[0]][current[1]] = true
				}
			}
		}
		result++
	}
	return -1
}

// 获取初始状态箱子、人及目的地的位置
func getInitialPositions(grid [][]byte) ([]int, []int, []int) {
	var box, person, target []int
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			switch grid[r][c] {
			case 'B':
				box = []int{r, c}
			case 'S':
				person = []int{r, c}
			case 'T':
				target = []int{r, c}
			}
			if box != nil && person != nil && target != nil {
				return box, person, target
			}
		}
	}
	return box, person, target
}

// 是否在矩阵之内，且不是墙
func isValid(pos []int, grid [][]byte) bool {
	r, c := pos[0], pos[1]
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && grid[r][c] != '#'
}

// person能否到达位置target，其中position是box的上下左右邻居位置里的一个
func canMoveTo(target, box, person []int, grid [][]byte) bool {
	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	queue := [][]int{person}
	const maxSize = 20
	visited := [maxSize][maxSize]bool{}
	visited[person[0]][person[1]] = true
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current[0] == target[0] && current[1] == target[1] {
			return true
		}
		for _, d := range dirs {
			r, c := current[0]+d[0], current[1]+d[1]
			if isValid([]int{r, c}, grid) &&
				!visited[r][c] &&
				(r != box[0] || c != box[1]) {
				queue = append(queue, []int{r, c})
				visited[r][c] = true
			}
		}
	}
	return false
}
