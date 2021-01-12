## [752. 打开转盘锁](https://leetcode-cn.com/problems/open-the-lock)
```text
你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。
每个拨轮可以自由旋转：例如把 '9' 变为  '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1。

示例 1:
输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
输出：6
解释：
可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
因为当拨动到 "0102" 时这个锁就会被锁定。

示例 2:
输入: deadends = ["8888"], target = "0009"
输出：1
解释：
把最后一位反向旋转一次即可 "0000" -> "0009"。

示例 3:
输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
输出：-1
解释：
无法旋转到目标数字且不被锁定。
示例 4:

输入: deadends = ["0000"], target = "8888"
输出：-1

提示：

死亡列表 deadends 的长度范围为 [1, 500]。
目标数字 target 不会在 deadends 之中。
每个 deadends 和 target 中的字符串的数字会在 10,000 个可能的情况 '0000' 到 '9999' 中产生。
```
这是个好问题，思路是用广度优先搜索来找到最短路径，有一些细节好好挖掘能形成编程技巧。<br>
如果对BFS不熟，可以先参考这个题解：[[778] 在水位上升的泳池中游泳](../swim-in-rising-water/readme.md)
```text
解法一：常规广度优先搜索
用一个集合存储当前访问状态列表
从 0000 开始搜索，对于每一个状态，可以扩展到最多 8 个状态，即将每一位增加 1 或减少 1，
将这些状态中没有搜索过并且不在 deadends 中的状态全部加入到队列中，并继续搜索，直到到达目标状态。

注意 0000 本身有可能也在 deadends 中。
为方便计数，可以在visited里维护到达每个状态的步数
存储当前访问状态列表的集合可以任选，切片、list、map等都可以，对顺序没有要求
```
```go
func openLock(deadends []string, target string) int {
	// 预处理，边界情况及时返回
	const initial = "0000"
	if target == initial {
		return 0
	}
	isDead := make(map[string]bool, 0) // 将deadends处理成哈希表，方便迅速查找某个状态是否在里边
	for _, v := range deadends {
		isDead[v] = true
	}
	if isDead[initial] {
		return -1
	}
	// 广度优先搜索
	// BFS 准备
	queue := list.New()
	queue.PushBack(initial)
	visited := map[string]int{initial: 0} // 已访问过的状态集合，并记录到达该状态花的步数
	visitNextStatus := func(origin string) { // 获取origin的下一个状态（在这里有8种）并将合适的状态入队
		b := []byte(origin)
		for i, c := range b {
			for d := -1; d <= 1; d += 2 { // 1 或 -1， b的每一位要加1或减1，以得到下一个状态
				b[i] = byte((int(c-'0')+d+10)%10) + '0'
				next := string(b)
				b[i] = c // 恢复原字符串
				if _, ok := visited[next]; ok || isDead[next] {
					continue
				}
				visited[next] = visited[origin] + 1
				queue.PushBack(next)
			}
		}
	}
	// BFS主体逻辑
	for queue.Len() > 0 {
		s := queue.Remove(queue.Front()).(string)
		if s == target {
			return visited[s]
		}
		visitNextStatus(s)
	}
	return -1
}
```
```text
解法二：优化后的广度优先搜索
首先一开始可以把deadends里的状态放入visited，这样不用另开辟isDead的空间；实际测试这个优化对性能提升并不明显
看了leetcode题解，有这样一个优化：
可以分别从初始状态和目标状态向中间状态搜索，如果中间有状态相同，即完成了BFS
实际测试发现，维持这两个搜索使用的集合大小相当，会显著优化时间花费
注意，为了能迅速获知两个集合是不是有相同元素，这里的集合用哈希表再合适不过
```
```go
func openLock(deadends []string, target string) int {
	const initial = "0000"
	if target == initial {
		return 0
	}
	visited := make(map[string]bool)
	for _, v := range deadends {
		visited[v] = true
	}
	if visited[initial] {
		return -1
	}

	start := map[string]bool{initial: true}
	end := map[string]bool{target: true}
	return bfs(start, end, visited, 0) // count代表步数，从0开始
}

// 模拟双向搜索。 count为步数
func bfs(start, end, visited map[string]bool, count int) int {
	if len(start) > len(end) { // 从已经遍历少的那一端开始， 维持两端搜索的数量相当，能明显优化搜索步数
		return bfs(end, start, visited, count)
	}
	if len(start) <= 0 {
		return -1
	}

	nextStatus := make(map[string]bool) //存储start端下一步需要处理的状态
	for s := range start {
		if _, ok := end[s]; ok { // end队列也有，说明从初始到目标状态的一个通路形成了
			return count
		}
		visited[s] = true
		b := []byte(s)
		for i, c := range b {
			for d := -1; d <= 1; d += 2 { // 1 或 -1， b的每一位要加1或减1，以得到下一个状态
				b[i] = byte((int(c-'0')+d+10)%10) + '0'
				next := string(b)
				b[i] = c // 复原状态
				if visited[next] {
					continue
				}
				nextStatus[next] = true
			}
		}
	}
	count++
	return bfs(nextStatus, end, visited, count)
}
```