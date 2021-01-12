package nodes

import (
	"container/list"
	"math"
)

/*
847. 访问所有节点的最短路径
https://leetcode-cn.com/problems/shortest-path-visiting-all-nodes

给出 graph 为有 N 个节点（编号为 0, 1, 2, ..., N-1）的无向连通图。
graph.length = N，且只有节点 i 和 j 连通时，j != i 在列表 graph[i] 中恰好出现一次。
返回能够访问所有节点的最短路径的长度。你可以在任一节点开始和停止，也可以多次重访节点，并且可以重用边。

示例 1：
输入：[[1,2,3],[0],[0],[0]]
输出：4
解释：一个可能的路径为 [1,0,2,0,3]

示例 2：
输入：[[1],[0,2,4],[1,3,4],[2],[1,2]]
输出：4
解释：一个可能的路径为 [0,1,4,2,3]


提示：
1 <= graph.length <= 12
0 <= graph[i].length < graph.length
*/

/*
注意到图的节点数很少，最多 12 个，可以用一个 int 变量 state 的二进制表示来代表每个节点是否被访问过
一开始所有节点都没被访问，
state 是 000000...000
可以尝试任意一个节点，for i := range [0, n-1] {...}
假设尝试节点 3， 那么状态变为 000100...000，接下来可以访问节点 3 的邻居 for i := range neibors of node 3  {...}
这样节点+状态即可以作为一个抽象的图，可以用广度优先的策略遍历该图
*/

type item struct {
	node, state int
}

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	total := 1 << n
	queue := list.New()
	dist := genMemo(n, total)

	for node := 0; node < n; node++ {
		state := 1 << node
		queue.PushBack(item{node, state})
		dist[node][state] = 0
	}

	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(item)
		d := dist[cur.node][cur.state]
		if cur.state == total-1 {
			return d
		}
		for _, next := range graph[cur.node] {
			state := cur.state | (1 << next)
			if d+1 < dist[next][state] {
				queue.PushBack(item{next, state})
				dist[next][state] = d + 1
			}
		}
	}
	return 0
}

func genMemo(m, n int) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		for j := range res[i] {
			res[i][j] = math.MaxInt32
		}
	}
	return res
}
