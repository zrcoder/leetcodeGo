package escape_a_large_maze

import (
	"container/list"
	"strconv"
)

/*
1036. 逃离大迷宫
https://leetcode-cn.com/problems/escape-a-large-maze

在一个 10^6 x 10^6 的网格中，每个网格块的坐标为 (x, y)，其中 0 <= x, y < 10^6。
我们从源方格 source 开始出发，意图赶往目标方格 target。
每次移动，我们都可以走到网格中在四个方向上相邻的方格，只要该方格不在给出的封锁列表 blocked 上。
只有在可以通过一系列的移动到达目标方格时才返回 true。否则，返回 false。

示例 1：
输入：blocked = [[0,1],[1,0]], source = [0,0], target = [0,2]
输出：false
解释：
从源方格无法到达目标方格，因为我们无法在网格中移动。

示例 2：
输入：blocked = [], source = [0,0], target = [999999,999999]
输出：true
解释：
因为没有方格被封锁，所以一定可以到达目标方格。

提示：
0 <= blocked.length <= 200
blocked[i].length == 2
0 <= blocked[i][j] < 10^6
source.length == target.length == 2
0 <= source[i][j], target[i][j] < 10^6
source != target
*/

/*
比较有意思的BFS，有几个点要考虑：
1. 首先，记录某个坐标是否访问过不能用一个 10^6 * 10^6 的二维数组，这实在太大
换用哈希表，只存访问过的坐标，简单起见，可以将横、纵坐标转化为string来存储
2. 其次，哈希表也不能一直加元素，可能超内存也可能超时。
在bfs过程中，如果队列已经包含的方格数目超过给定的阻塞方格数，可能是沿着阻塞方格的边界添加了些普通方格，
添加的普通方格数超过了给定阻塞方个数，貌似可以确定起点终点是联通的，直接返回true吧
3. 最后，要非常注意，2最后返回true其实不完全正确，比如这样的用例： `终点`上下左右四个点都是阻塞方格，但是其他任意格子都是非阻塞方格的情况，结论就错了
为解决这种情况，可以将起点和终点互换，再尝试一次bfs，如果两次bfs都返回true，才表明起点终点是联通的
*/
func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	if len(blocked) == 0 {
		return true
	}
	blocks := make(map[string]bool, len(blocked))
	for _, v := range blocked {
		if isSame(v, source) || isSame(v, target) {
			return false
		}
		blocks[pointToString(v)] = true
	}
	return bfs(source, target, blocks) && bfs(target, source, blocks)
}

func bfs(source, target []int, blocks map[string]bool) bool {
	const maxSize = 1e6
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	visited := map[string]bool{pointToString(source): true}
	queue := list.New()
	queue.PushBack(source)
	for queue.Len() > 0 {
		// 尝试过的普通方格数超过了给定阻塞方个数，可以直接返回true
		if queue.Len() > len(blocks) {
			return true
		}
		point := queue.Remove(queue.Front()).([]int)
		for _, v := range dirs {
			r, c := point[0]+v[0], point[1]+v[1]
			if r < 0 || r >= maxSize || c < 0 || c >= maxSize {
				continue
			}
			nextPoint := []int{r, c}
			key := pointToString(nextPoint)
			if visited[key] || blocks[key] {
				continue
			}
			if isSame(nextPoint, target) {
				return true
			}
			visited[key] = true
			queue.PushBack(nextPoint)
		}
	}
	return false
}

func isSame(p1, p2 []int) bool {
	return p1[0] == p2[0] && p1[1] == p2[1]
}

func pointToString(point []int) string {
	return strconv.Itoa(point[0]) + "," + strconv.Itoa(point[1])
}
