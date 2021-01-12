package sum_of_distances_in_tree

/*
834. 树中距离之和
https://leetcode-cn.com/problems/sum-of-distances-in-tree

给定一个无向、连通的树。树中有 N 个标记为 0...N-1 的节点以及 N-1 条边 。

第 i 条边连接节点 edges[i][0] 和 edges[i][1] 。

返回一个表示节点 i 与其他所有节点距离之和的列表 ans。

示例 1:

输入: N = 6, edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]
输出: [8,12,6,10,10,10]
解释:
如下为给定的树的示意图：
  0
 / \
1   2
   /|\
  3 4 5

我们可以计算出 dist(0,1) + dist(0,2) + dist(0,3) + dist(0,4) + dist(0,5)
也就是 1 + 1 + 2 + 2 + 2 = 8。 因此，answer[0] = 8，以此类推。
说明: 1 <= N <= 10000
*/
/*
参考题解：
https://leetcode-cn.com/problems/sum-of-distances-in-tree/solution/shou-hua-tu-jie-shu-zhong-ju-chi-zhi-he-shu-xing-d

时间复杂度：O(N)，其中 N 是树中的节点个数。只需要遍历整棵树两次即可得到答案，其中每个节点被访问两次。
空间复杂度：O(N)。
*/
func sumOfDistancesInTree(n int, edges [][]int) []int {
	// 为方便迅速得到某个节点的相邻节点，将输入处理成邻接表
	graph := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	// 当前记录的是每个节点到它所在子树的节点的距离和
	dp := make([]int, n)
	// 记录每个节点作为根节点的子树的节点总数
	sz := make([]int, n)
	var dfs func(node, parent int)
	dfs = func(node, parent int) {
		sz[node] = 1 // 节点自身个数需计算
		for _, v := range graph[node] {
			if v == parent {
				continue
			}
			dfs(v, node)
			sz[node] += sz[v]
			dp[node] += dp[v] + sz[v]
		}
	}
	dfs(0, -1)

	// 做换根操作，之后的 dp[u] 表示节点 `u` 到其他节点的距离和。
	var dfs1 func(node, parent int)
	dfs1 = func(node, parent int) {
		for _, v := range graph[node] {
			if v == parent {
				continue
			}
			dp[v] = dp[node] - sz[v] + (n - sz[v])
			dfs1(v, node)
		}
	}
	dfs1(0, -1) // 第一个参数要和上边 dfs 调用时一致
	return dp
}
