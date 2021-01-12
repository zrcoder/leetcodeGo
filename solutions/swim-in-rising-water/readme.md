## [778. 水位上升的泳池中游泳](https://leetcode-cn.com/problems/swim-in-rising-water)
```text
在一个 N * N 的坐标方格 grid 中，每一个方格的值 grid[i][j] 表示在位置 (i,j) 的平台高度。
现在开始下雨了。当时间为 t 时，此时雨水导致水池中任意位置的水位为 t 。
你可以从一个平台游向四周相邻的任意一个平台，但是前提是此时水位必须同时淹没这两个平台。
假定你可以瞬间移动无限距离，也就是默认在方格内部游动是不耗时的。当然，在你游泳的时候你必须待在坐标方格里面。
你从坐标方格的左上平台 (0，0) 出发。最少耗时多久你才能到达坐标方格的右下平台 (N-1, N-1)？

示例 1:
输入: [[0,2],[1,3]]
输出: 3
解释:
时间为0时，你位于坐标方格的位置为 (0, 0)。
此时你不能游向任意方向，因为四个相邻方向平台的高度都大于当前时间为 0 时的水位。
等时间到达 3 时，你才可以游向平台 (1, 1). 因为此时的水位是 3，坐标方格中的平台没有比水位 3 更高的，所以你可以游向坐标方格中的任意位置

示例2:
输入: [[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6]]
输入: 16
解释:
 0  1  2  3  4
24 23 22 21  5
12 13 14 15 16
11 17 18 19 20
10  9  8  7  6
最终的路线：
 0- 1- 2- 3- 4
             5
12-13-14-15-16
11
10- 9- 8- 7- 6

我们必须等到时间为 16，此时才能保证平台 (0, 0) 和 (4, 4) 是连通的

提示:
2 <= N <= 50.
grid[i][j] 位于区间 [0, ..., N*N - 1] 内。
```
**如下解法都是广度优先搜索（BFS），值得一提的是这个策略可以扩展，任意指定起点和终点（不一定是左上角和右下角），行、列数不同也能解决问题**
### 分析
#### 解法一：常规广度优先搜索
* 可以先简化下这个问题：假如给定一个水位h，判断能否游到终点
* 简化问题搞定后，离解决这个问题就不远了
```text
1.对于一个特定的水位h，怎么判断最终能否到达终点呢？用广度优先搜索(BFS)最合适
从起点开始，将相邻且高度不大于h的平台坐标放入集合，然后把这些坐标一一从集合取出，取出后再将它们符合条件的相邻平台坐标放入集合
不断重复这个过程，每次取出后判断是不是终点，是的话就ok了；如果集合已经空了，也没有找到终点，说明无法到达
这个集合用队、栈、list或者set（map）、切片甚至其他容器都可以，对顺序没有要求，当然如果要求出最后能到达终点的路径的话，可以想想怎么解决，这里不展开。

时间复杂度是O(n^2)，所有坐标共n^2个，最坏情况每个坐标都要放入集合又取出
空间复杂度与时间复杂度相同，集合的大小，最坏情况是所有n^2个平台都放入
```
```go
func canReach(h int, grid [][]int) bool {
	const maxN = 50
	n := len(grid)
	// 找相邻位置的一个技巧，减少代码量
	// 遍历dr和dc用原来的横纵坐标加上对应dr、dc里的坐标即得到上下左右相邻位置之一的坐标
	// 也可以把dr、dc合并为一个二维切片（数组）
	dr := []int{1, -1, 0, 0}
	dc := []int{0, 0, 1, -1}
	visited := [maxN][maxN]bool{} // 标识已经访问过的平台，也可以用map；其实应该是n*n大小，但是用n的话代码要多几行，需要遍历初始化每一行
	visited[0][0] = true
	set := list.New()
	set.PushBack([]int{0, 0}) // 用长度为2的切片代表一个点；初始位置放入集合
	for set.Len() > 0 {
		pos := set.Remove(set.Back()).([]int)
		row, column := pos[0], pos[1]
		if row == n-1 && column == n-1 {
			return true
		}
		for i := 0; i < len(dr); i++ {
			nextRow, nextColumn := row+dr[i], column+dc[i]
			if nextRow >= 0 && nextRow < n && nextColumn >= 0 && nextColumn < n &&
				!visited[nextRow][nextColumn] && grid[nextRow][nextColumn] <= h {
				set.PushBack([]int{nextRow, nextColumn})
				visited[nextRow][nextColumn] = true
			}
		}
	}
	return false
}
```
```text
2.假设所有平台最低高度、最高高度分别为min、max，答案就在区间 [min, max] 内
不过有个点要注意，起点的高度如果大于min，得到的结果可能是错的，说白了不允许从高平台跳水跳到低平台，
所以答案的精确区间是[grid[0][0], max]

朴素解法：可以尝试高度从grid[0][0]向max递增，对每个高度，如果发现最终能到达终点，那么当前的高度就是答案
时间复杂度为调用BFS的次数*BFS的复杂度，即O((max-grid[0][0]+1) * n^2)；空间复杂度相同

当然在区间[grid[0][0], max]上用二分法更快，时间复杂度降为： O(log(max-grid[0][0]+1) * n^2)
```
```go
/*
不用二分法的朴素实现，时间复杂度O(n^2*(max-grid[0][0]+1))
leetcode实测会花费376 ms，其他解法的时间在12-28ms内
*/
func swimInWater0(grid [][]int) int {
	start, end := grid[0][0], max(grid)
	for i := start; i < end; i++ {
		if canReach(i, grid) {
			return i
		}
	}
	return end
}
```
```go
// 二分法，用标准库,减少点代码量
func swimInWater2(grid [][]int) int {
	return sort.Search(max(grid)+1, func(i int) bool { // 这里其实有点浪费，在[0,max]的区间里搜所的
		if i < grid[0][0] {
			return false
		}
		return canReach(i, grid)
	})
}
```
用到的max函数即找出grid里最大的元素：
```go
func max(grid [][]int) int {
	result := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid); c++ {
			if grid[r][c] > result {
				result = grid[r][c]
			}
		}
	}
	return result
}
```
## 解法二： 借助小顶堆的广度优先搜索
```text
和解法一本质是一样的
每到一个平台，在相邻平台+以前经过平台的相邻平台中选择高度最小的平台。
题目中两个示例过于特殊，我们看下边的例子：
0 1 4
2 8 7
3 6 5
为方便叙述，例子里特别让平台高度不同，这样可以用平台的高度代表平台
第一步有两个选择 [1,2]，选择平台1，之后新增 [4,8] 两平台能到，现在所有能到的平台为：[2,4,8] （1 已经去过了）
选择高度最低的平台2，之后又多了平台3能到，此时所有能到的平台为：[3,4,8] （2已经去过了）
选择3，多了6可以到，此时所有能到的平台为[4,6,8]（3已经去过了）
选择4，现在能到的平台是[6,7,8] （4已经去过了）
选择第6，平台6能达到终点
观察刚才经过的路线1-2-3-4-6-5，对应水位最高的平台为6，就是答案
建议再结合题目里的示例来理解一遍~

用什么数据结构来存储周围能到的平台呢？因为要迅速找到最小的平台，小顶堆再合适不过
每一次，小顶堆存储周围可以到的平台集合，并选择高度最小的平台，游到该平台后将该平台出堆且将其相邻平台入堆（已经入过堆的就不必了）
以这种方式到达终点，途经最高的平台就是答案
就是借助小顶堆做广度优先搜索

时间复杂度： O(n^2 * log(n^2))=O(n^2 * 2logn)=O(n^2 * logn), 最大经过n^2个平台，每个平台需要O(log(n^2))时间进出堆
空间复杂度：O(n^2)，是堆的最大空间
```
```go
// 平台结构体，方便自定义heap实现
type pos struct {
	height, r, c int // 高度和横纵坐标
}

func swimInWater(grid [][]int) int {
	const maxN = 50
	n := len(grid)
	dr := []int{1, -1, 0, 0}
	dc := []int{0, 0, 1, -1}
	visited := [maxN][maxN]bool{}
	result := 0
	pq := &posHeap{}
	heap.Push(pq, pos{height: grid[0][0], r: 0, c: 0})

	for pq.Len() > 0 {
		info := heap.Pop(pq).(pos) // 游到当前最低的平台上
		if grid[info.r][info.c] > result {
			result = grid[info.r][info.c]
		}
		if info.r == n-1 && info.c == n-1 { // 终点
			break
		}
		for i := 0; i < len(dr); i++ {
			r, c := info.r+dr[i], info.c+dc[i]
			if r >= 0 && r < n && c >= 0 && c < n && !visited[r][c] {
				heap.Push(pq, pos{height: grid[r][c], r: r, c: c})
				visited[r][c] = true
			}
		}
	}
	return result
}

type posHeap []pos

func (h posHeap) Len() int            { return len(h) }
func (h posHeap) Less(i, j int) bool  { return h[i].height < h[j].height }
func (h posHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *posHeap) Push(x interface{}) { *h = append(*h, x.(pos)) }
func (h *posHeap) Pop() interface{} {
	pos := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return pos
}
```
[参考解答](d.go)