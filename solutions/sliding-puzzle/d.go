/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package sliding_puzzle

import (
	"container/list"
	"strings"
)

/*
773. 滑动谜题 https://leetcode-cn.com/problems/sliding-puzzle/
在一个 2 x 3 的板上（board）有 5 块砖瓦，用数字 1~5 来表示, 以及一块空缺用 0 来表示.

一次移动定义为选择 0 与一个相邻的数字（上下左右）进行交换.

最终当板 board 的结果是 [[1,2,3],[4,5,0]] 谜板被解开。

给出一个谜板的初始状态，返回最少可以通过多少次移动解开谜板，如果不能解开谜板，则返回 -1 。

示例：

输入：board = [[1,2,3],[4,0,5]]
输出：1
解释：交换 0 和 5 ，1 步完成

输入：board = [[1,2,3],[5,4,0]]
输出：-1
解释：没有办法完成谜板

输入：board = [[4,1,2],[5,0,3]]
输出：5
解释：
最少完成谜板的最少移动次数是 5 ，
一种移动路径:
尚未移动: [[4,1,2],[5,0,3]]
移动 1 次: [[4,1,2],[0,5,3]]
移动 2 次: [[0,1,2],[4,5,3]]
移动 3 次: [[1,0,2],[4,5,3]]
移动 4 次: [[1,2,0],[4,5,3]]
移动 5 次: [[1,2,3],[4,5,0]]

输入：board = [[3,2,4],[1,5,0]]
输出：14

提示：
board 是一个如上所述的 2 x 3 的数组.
board[i][j] 是一个 [0, 1, 2, 3, 4, 5] 的排列.
*/
/*
可以把这道题看成一个找出图中最短路径的问题。
每个节点都是棋盘的一个状态，如果两个状态之间可以通过一步操作来完成转换，就用一条边将这两个节点相连。
用 广度优先搜索 来解决最短路径问题。
在广度优先搜索实现中，需要将节点表示成可以哈希的数据结构，同时还需要找到每个节点的邻居节点。


时间复杂度：O(R * C * (R * C)!)，其中 R, CR,C 为棋盘的行数和列数。最多有 O((R * C)!) 种可能的棋盘状态。

空间复杂度：O(R * C * (R * C)!)

作者：LeetCode
链接：https://leetcode-cn.com/problems/sliding-puzzle/solution/hua-dong-mi-ti-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func slidingPuzzle(board [][]int) int {
	queue := list.New()
	target := "123450"
	start := hash(board)
	queue.PushBack(start)
	visited := map[string]int{start: 0} // 记录指定状态是否出现过，并记录到达该状态用了多少步
	/*
		0,1,2,
		3,4,5
	*/
	neighbors := [][]int{
		{1, 3},
		{0, 2, 4},
		{1, 5},
		{0, 4},
		{1, 3, 5},
		{2, 4},
	}
	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).(string)
		if curr == target {
			return visited[target]
		}
		index := 0
		for curr[index] != '0' {
			index++
		}
		for _, v := range neighbors[index] {
			next := swap(curr, index, v)
			if _, ok := visited[next]; ok {
				continue
			}
			visited[next] = visited[curr] + 1
			queue.PushBack(next)
		}
	}
	return -1
}

func hash(s [][]int) string {
	result := strings.Builder{}
	for _, row := range s {
		for _, num := range row {
			result.WriteByte(byte(num + '0'))
		}
	}
	return result.String()
}

func swap(s string, i, j int) string {
	result := []byte(s)
	result[i], result[j] = result[j], result[i]
	return string(result)
}
