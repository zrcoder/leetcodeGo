package redundant_connection

import "reflect"

/*
684. 冗余连接
https://leetcode-cn.com/problems/redundant-connection

在本问题中, 树指的是一个连通且无环的无向图。

输入一个图，该图由一个有着N个节点 (节点值不重复1, 2, ..., N) 的树及一条附加的边构成。
附加的边的两个顶点包含在1到N中间，这条附加的边不属于树中已存在的边。

结果图是一个以边组成的二维数组。
每一个边的元素是一对[u, v] ，满足 u < v，表示连接顶点u 和v的无向图的边。

返回一条可以删去的边，使得结果图是一个有着N个节点的树。
如果有多个答案，则返回二维数组中最后出现的边。
答案边 [u, v] 应满足相同的格式 u < v。

示例 1：
输入: [[1,2], [1,3], [2,3]]
输出: [2,3]
解释: 给定的无向图为:
  1
 / \
2 - 3

示例 2：
输入: [[1,2], [2,3], [3,4], [1,4], [1,5]]
输出: [1,4]
解释: 给定的无向图为:
5 - 1 - 2
    |   |
    4 - 3

注意:
输入的二维数组大小在 3 到 1000。
二维数组中的整数在1到N之间，其中N是输入数组的大小。
更新(2017-09-26):
我们已经重新检查了问题描述及测试用例，明确图是无向图。
对于有向图详见冗余连接II。对于造成任何不便，我们深感歉意。
*/
// Union-Find
func findRedundantConnection(edges [][]int) []int {
	uf := make([]int, len(edges)+1)
	for i := range uf {
		uf[i] = i
	}
	for _, e := range edges {
		if !union(uf, e[0], e[1]) {
			return e
		}
	}
	return nil
}

// DFS
func findRedundantConnectionDfs(edges [][]int) []int {
	n := len(edges)
	graph := make([][]int, n+1)
	for _, edge := range edges {
		src, dest := edge[0], edge[1]
		if len(graph[src]) > 0 && len(graph[dest]) > 0 && isConnected(graph, src, dest, make([]bool, n+1)) {
			return edge
		}
		graph[src] = append(graph[src], dest)
		graph[dest] = append(graph[dest], src)
	}
	return nil
}

func isConnected(graph [][]int, src, dest int, seen []bool) bool {
	if seen[src] {
		return false
	}
	if src == dest {
		return true
	}
	seen[src] = true
	for _, nei := range graph[src] {
		if isConnected(graph, nei, dest, seen) {
			return true
		}
	}
	return false
}

/*
685. 冗余连接 II
https://leetcode-cn.com/problems/redundant-connection-ii

在本问题中，有根树指满足以下条件的有向图。
该树只有一个根节点，所有其他节点都是该根节点的后继。每一个节点只有一个父节点，除了根节点没有父节点。

输入一个有向图，该图由一个有着N个节点 (节点值不重复1, 2, ..., N) 的树及一条附加的边构成。
附加的边的两个顶点包含在1到N中间，这条附加的边不属于树中已存在的边。

结果图是一个以边组成的二维数组。
每一个边 的元素是一对 [u, v]，用以表示有向图中连接顶点 u 和顶点 v 的边，其中 u 是 v 的一个父节点。

返回一条能删除的边，使得剩下的图是有N个节点的有根树。若有多个答案，返回最后出现在给定二维数组的答案。

示例  1:

输入: [[1,2], [1,3], [2,3]]
输出: [2,3]
解释: 给定的有向图如下:
  1
 / \
v   v
2-->3

示例 2:

输入: [[1,2], [2,3], [3,4], [4,1], [1,5]]
输出: [4,1]
解释: 给定的有向图如下:
5 <- 1 -> 2
     ^    |
     |    v
     4 <- 3

注意:

二维数组大小的在3到1000范围内。
二维数组中的每个整数在1到N之间，其中 N 是二维数组的大小。
*/

/*
如果所有节点最多只有一个父节点，那么解法和上面问题（684）完全一样；即只需找到使图出现环的那条边。
否则，先找到有两个父节点的节点 X，结果就是指向 X 的这两条边里的一条，具体是哪一条？
看两个例子：

1 -> 2
|   ^
v /
3
删除 1->2 或 3->2 都可以，根据题意，删除在 edges 数组最后出现的即可

2 -> 1 <- 3
^   /
|  V
 4
如果删除 3->1， 剩余的图就会有一个环； 所以只能删除 2->1
*/
func findRedundantDirectedConnection(edges [][]int) []int {
	// 如果某个节点有两条入边，这两条入边之一就是答案
	candidates := [2][]int{}
	parent := make([]int, len(edges)+1)
	for _, e := range edges {
		if parent[e[1]] != 0 {
			candidates = [2][]int{{parent[e[1]], e[1]}, {e[0], e[1]}}
			break
		}
		parent[e[1]] = e[0]
	}
	// 没有有两条入边的节点，只需要找到使图成环的边
	if candidates[0] == nil {
		return detectCycle(edges, nil)
	}
	// 在两条候选边，需要确定返回哪一条
	if detectCycle(edges, candidates[1]) == nil { // // 忽略第二条候选边没有在图里发现环
		return candidates[1]
	}
	return candidates[0] // 忽略第二条候选边在图里发现了环
}

func detectCycle(edges [][]int, ignore []int) []int {
	uf := make([]int, len(edges)+1)
	for i := range uf {
		uf[i] = i
	}
	for _, e := range edges {
		if reflect.DeepEqual(e, ignore) {
			continue
		}
		if !union(uf, e[0], e[1]) {
			return e
		}
	}
	return nil
}

func union(uf []int, x, y int) bool {
	xRoot, yRoot := find(uf, x), find(uf, y)
	if xRoot == yRoot {
		return false
	}
	uf[xRoot] = uf[yRoot]
	return true
}

func find(uf []int, x int) int {
	for x != uf[x] {
		x, uf[x] = uf[x], uf[uf[x]]
	}
	return x
}
