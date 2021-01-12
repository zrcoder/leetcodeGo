# [1135. 最低成本联通所有城市](https://leetcode-cn.com/problems/connecting-cities-with-minimum-cost)
想象一下你是个城市基建规划者，地图上有 N 座城市，它们按以 1 到 N 的次序编号。<br>
给你一些可连接的选项 conections，其中每个选项 conections[i] = [city1, city2, cost] 表示将城市 city1 和城市 city2 连接所要的成本。<br>
（连接是双向的，也就是说城市 city1 和城市 city2 相连也同样意味着城市 city2 和城市 city1 相连）。<br>
返回使得每对城市间都存在将它们连接在一起的连通路径（可能长度为 1 的）最小成本。该最小成本应该是所用全部连接代价的综合。<br>
如果根据已知条件无法完成该项任务，则请你返回 -1。<br>
```
示例 1：
输入：N = 3, conections = [[1,2,5],[1,3,6],[2,3,1]]
输出：6

解释：
选出任意 2 条边都可以连接所有城市，我们从中选取成本最小的 2 条。
```
```
示例 2：
输入：N = 4, conections = [[1,2,3],[3,4,4]]
输出：-1

解释： 
即使连通所有的边，也无法连接所有城市。
```
提示：
```
1 <= N <= 10000
1 <= conections.length <= 10000
1 <= conections[i][0], conections[i][1] <= N
0 <= conections[i][2] <= 10^5
conections[i][0] != conections[i][1]
```
## 最小生成树
把这些城市看成地图上的一个个点，根据connections的信息可以知道每条连线的成本，即边的权值<br>
按照成本从低到高（即按照权值从小到大遍历边），这样做：<br>
假设当前边的两端分别是a，b两座城市，如果a、b是联通的，则不联接a、b；否则联接，当然联接后要将成本计入结果<br>

细致来说a、b的情况如下：
```
i：如果a、b当前都还是孤立的点，即还没有和其他城市连过线——联接
ii：如果a、b有一个是孤立的点，另一个已经和其他城市联接过——联接
iii： 如果a、b之前都和其他城市联接过，分两种情况：
iii-i：a、b已经是联通的——不联接
iii-ii：a、b不联通——联接
```
共N个城市，总共需要联接N-1次就会将所有城市联通；因为用了贪心策略，按照成本从小到大遍历操作，最终的成本是最小的<br>
有个关键的问题是怎么快速判断两个城市是否联通，这里有个有意思的数据结构：[并查集](../../learn/union-find.md)

```go
func minimumCost(n int, connections [][]int) int {
	if len(connections) < n-1 { // 要有每个城市的联接信息，最终才能将所有城市联通，否则总有落单的
		return -1
	}
	sort.Slice(connections, func(i, j int) bool { // 按照成本排序
		return connections[i][2] < connections[j][2]
	})
	unionFind := NewUnionFind(n)
	connected, result, i := 0, 0, 0
	for connected < n-1 {
		connection := connections[i]
		i++
		city1, city2 := connection[0]-1, connection[1]-1
		city1, city2 = unionFind.Find(city1), unionFind.Find(city2)
		if city1 != city2 {
			unionFind.Union(city1, city2)
			connected++
			result += connection[2]
		}
	}
	return result
}
```

[完整实现参考](d.go)