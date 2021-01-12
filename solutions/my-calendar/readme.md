## 我的日程安排表
来源：[力扣（LeetCode）](https://leetcode-cn.com)
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
### [729 我的日程安排表 I](https://leetcode-cn.com/problems/my-calendar-i)
```text
实现一个 MyCalendar 类来存放你的日程安排。
如果要添加的时间内没有其他安排，则可以存储这个新的日程安排。

MyCalendar 有一个 book(int start, int end)方法。
它意味着在 start 到 end 时间内增加一个日程安排，
注意，这里的时间是半开区间，即 [start, end),
实数 x 的范围为，  start <= x < end。

当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生重复预订。

每次调用 MyCalendar.book方法时，如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true。
否则，返回 false 并且不要将该日程安排添加到日历中。

请按照以下步骤调用 MyCalendar 类:
MyCalendar cal = new MyCalendar();
MyCalendar.book(start, end)

示例 1:
MyCalendar();
MyCalendar.book(10, 20); // returns true
MyCalendar.book(15, 25); // returns false
MyCalendar.book(20, 30); // returns true
解释: 
第一个日程安排可以添加到日历中.  第二个日程安排不能添加到日历中，因为时间 15 已经被第一个日程安排预定了。
第三个日程安排可以添加到日历中，因为第一个日程安排并不包含时间 20 。
说明:
每个测试用例，调用 MyCalendar.book 函数最多不超过 100次。
调用函数 MyCalendar.book(start, end)时， start 和 end 的取值范围为 [0, 10^9]。
```
共三个解法，从朴素实现开始逐步优化：
```text
1.朴素实现：用一个集合来存储不断添加的日程；每次添加需要确定已有日程是否和要添加的日程有重叠
集合可以用list或切片；查找需要遍历集合，插入可以简单在末尾追加
假设已有日程有n个，时间复杂度为O(n),空间复杂度为集合的大小，O(n)
```
```go
type Interval struct {
    start, end int
}

type MyCalendar struct {
    calendar *list.List
}

func Constructor() MyCalendar {
    return MyCalendar{calendar: list.New()}
}

func (mc *MyCalendar) Book(start int, end int) bool {
    for e := mc.calendar.Front(); e != nil; e = e.Next() {
        interval := e.Value.(Interval)
        if max(start, interval.start) < min(end, interval.end) {
            return false
        }
    }
    mc.calendar.PushBack(Interval{start:start, end:end})
    return true
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
```
```text
2.二分法优化朴素实现：
可以维持集合里的日程有序(可依据每个日程的开始时间排序)，每次插入新日程就可以用二分法迅速找到要插入的位置
用list或堆（优先队列）的话，查找用不了二分法，这里可用切片，查找要插入的位置时间复杂度为O(lgn),
插入的话，因为要把待插入位置及其后边元素一一后移，时间复杂度是O(n)
综合最坏情况时间复杂度、空间复杂度与朴素实现相同，为O(n)；但是在普遍情况下要比朴素实现快
```
```go
type Interval struct {
	start, end int
}

type MyCalendar struct {
	calendar []Interval
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (mc *MyCalendar) Book(start int, end int) bool {
	pos := sort.Search(len(mc.calendar), func(i int) bool {
		return mc.calendar[i].start >= start
	})
	if pos < len(mc.calendar) && mc.calendar[pos].start < end ||
		pos-1 >= 0 && mc.calendar[pos-1].end > start {
		return false
	}
    it := Interval{start:start, end:end}
    mc.calendar = append(append(mc.calendar[:pos:pos], it), mc.calendar[pos:]...)
	return true
}
```
```text
3.进一步优化插入时间复杂度：
在插入新日程的时候，如果集合用切片，则有一一向后移动元素的复杂度在，
有没有一个数据结构能在常数时间插入元素，还能保证元素有序，
且能在常数时间查询任意索引元素从而能用二分法做查询呢？
堆、list和切片都不行
二叉搜索树BST可堪此任，不过如果插入的顺序比较特别，bst会退化成一个链表
实际需要一个能维持平衡的搜索树，比如红黑树，像Java有TreeMap可用，
Go标准库并没有这样的数据结构，手动实现起来有些复杂，暂不尝试
仅朴素BST来一把
时间复杂度最坏O(n^2), 平均情况O(nlgn)；空间复杂度O(n)
```
```go
type Node struct {
	 left, right *Node
	 start, end int
}

func (n *Node) insert(node *Node) bool {
	if node.start >= n.end {
		if n.right == nil {
			n.right = node
			return true
		}
		return n.right.insert(node)
	}
	if node.end <= n.start {
		if n.left == nil {
			n.left = node
			return true
		}
		return n.left.insert(node)
	}
	return false
}

type MyCalendar struct {
	root *Node
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (mc *MyCalendar) Book(start int, end int) bool {
	node := &Node{start:start, end:end}
	if mc.root == nil {
		mc.root = node
		return true
	}
	return mc.root.insert(node)
}
```
```text
实际测试，朴素BST的实现比朴素实现好一点，但是不比切片的实现好
有缘再手动实现下红黑树吧~~
```
[参考实现](i/d.go)
### [731. 我的日程安排表 II](https://leetcode-cn.com/problems/my-calendar-ii)
与上面的问题类似，但是这次允许日程有所重叠：
```text
当三个日程安排有一些时间上的交叉时（例如三个日程安排都在同一时间内），就会产生三重预订。
每次调用 MyCalendar.book方法时，如果可以将日程安排成功添加到日历中而不会导致三重预订，返回 true。
否则，返回 false 并且不要将该日程安排添加到日历中。

请按照以下步骤调用MyCalendar 类:
MyCalendar cal = new MyCalendar(); 
MyCalendar.book(start, end)

示例：

MyCalendar();
MyCalendar.book(10, 20); // returns true
MyCalendar.book(50, 60); // returns true
MyCalendar.book(10, 40); // returns true
MyCalendar.book(5, 15); // returns false
MyCalendar.book(5, 10); // returns true
MyCalendar.book(25, 55); // returns true
解释： 
前两个日程安排可以添加至日历中。 第三个日程安排会导致双重预订，但可以添加至日历中。
第四个日程安排活动（5,15）不能添加至日历中，因为它会导致三重预订。
第五个日程安排（5,10）可以添加至日历中，因为它未使用已经双重预订的时间10。
第六个日程安排（25,55）可以添加至日历中，因为时间 [25,40] 将和第三个日程安排双重预订；
时间 [40,50] 将单独预订，时间 [50,55）将和第二个日程安排双重预订。

提示：
每个测试用例，调用 MyCalendar.book 函数最多不超过 1000次。
调用函数 MyCalendar.book(start, end)时， start 和 end 的取值范围为 [0, 10^9]。
```
```text
解决方法与上面的问题类似，除了calendar，可以另外增加一个集合存储双重预定重合的部分，不妨称其为overlap
新插入日程时，先遍历overlap看看有没有和overlap重叠，有则返回false，
否则插入calendar且注意与calendar有重叠需要同步往overlap里追加下

时空复杂度都是O(n)
```
```go
type interval struct {
	start, end int
}

type MyCalendarTwo struct {
	calendar, overlap []interval // 分别表示已经添加的所有日程和已有日程重复的时间段组成的列表
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (mc *MyCalendarTwo) Book(start int, end int) bool {
	for _, val := range mc.overlap {
		if start < val.end && end > val.start {
			return false
		}
	}
	for _, val := range mc.calendar {
		if start < val.end && end > val.start {
			it := interval{start: max(start, val.start), end: min(end, val.end)}
			mc.overlap = append(mc.overlap, it)
		}
	}
	it := interval{start: start, end: end}
	mc.calendar = append(mc.calendar, it)
	return true
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
```
```text
尝试用二分法优化朴素实现
要稍微复杂点，寻找要插入的位置用二分法；
但是在统计新日程和已有日程的重叠部分的时候，还是要从头遍历calendar，因为有可能一个区间跨越了多个区间
时空复杂度都是O(n)
实际测试并没有明显优化，还是需要一个类似红黑树的数据结构啊~
```
```go
type interval struct {
	start, end int
}

type MyCalendarTwo struct {
	// 分别表示已经添加的所有日程和已有日程重复的时间段组成的列表
	calendar, overlap []interval
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (mc *MyCalendarTwo) Book(start int, end int) bool {
	if len(mc.overlap) > 0 {
		// 二分搜索出新日程在overlap里的位置,不存在的话找需要插入的位置
		pos := sort.Search(len(mc.overlap), func(i int) bool {
			return mc.overlap[i].start >= start
		})
		// 查看搜索出的位置附件有没有和当前日程重叠的部分
		if pos < len(mc.overlap) && mc.overlap[pos].start < end ||
			pos-1 >= 0 && mc.overlap[pos-1].end > start {
			return false
		}
	}
	for _, v := range mc.calendar {
		if max(v.start, start) < min(v.end, end) {
			it := interval{start: max(start, v.start), end: min(end, v.end)}
			pos := sort.Search(len(mc.overlap), func(i int) bool {
				return mc.overlap[i].start >= it.start
			})
			// 在pos处插入it
			mc.overlap = append(append(mc.overlap[:pos:pos], it), mc.overlap[pos:]...)
		}
	}
	pos := sort.Search(len(mc.calendar), func(i int) bool {
		return mc.calendar[i].start >= start
	})
	it := interval{start: start, end: end}
	// 在pos处插入it
	mc.calendar = append(append(mc.calendar[:pos:pos], it), mc.calendar[pos:]...)
	return true
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
```
[参考实现](ii/d.go)
### [732. 我的日程安排表 III](https://leetcode-cn.com/problems/my-calendar-iii/)
这次不再限制日程重叠度，放开限制随便加日程，但是每次加完要返回一个整数K，表示最大的 K 次预订。
```text
当 K 个日程安排有一些时间上的交叉时（例如K个日程安排都在同一时间内），就会产生 K 次预订。

每次调用 MyCalendar.book方法时，返回一个整数 K ，表示最大的 K 次预订。

请按照以下步骤调用MyCalendar 类:
 MyCalendar cal = new MyCalendar(); 
MyCalendar.book(start, end)

示例 1:

MyCalendarThree();
MyCalendarThree.book(10, 20); // returns 1
MyCalendarThree.book(50, 60); // returns 1
MyCalendarThree.book(10, 40); // returns 2
MyCalendarThree.book(5, 15); // returns 3
MyCalendarThree.book(5, 10); // returns 3
MyCalendarThree.book(25, 55); // returns 3
解释: 
前两个日程安排可以预订并且不相交，所以最大的K次预订是1。
第三个日程安排[10,40]与第一个日程安排相交，最高的K次预订为2。
其余的日程安排的最高K次预订仅为3。
请注意，最后一次日程安排可能会导致局部最高K次预订为2，但答案仍然是3，
原因是从开始到最后，时间[10,20]，[10,40]和[5,15]仍然会导致3次预订。

说明:
每个测试用例，调用 MyCalendar.book 函数最多不超过 400次。
调用函数 MyCalendar.book(start, end)时， start 和 end 的取值范围为 [0, 10^9]。
```
```text
朴素实现。按照时间点来统计：
把原问题想象成在数轴上画线段，如果线段有重合，则重合的部分颜色加深
所幸我们关注的点都是整数。对于数轴上每个点，统计其被多少条线段覆盖，可为点增加一个深度的属性来记录
为了能在常数时间插入点，可用list
时空复杂度都是O(n)
```
```go
type point struct {
	pos  int // 该点在数轴上的位置。
	deep int // 该点的深度——即被多少条线段包含。
}

type MyCalendarThree struct {
	points *list.List
	k      int
}

func Constructor() MyCalendarThree {
	mc := MyCalendarThree{points: list.New()}
	// 结合list特点, 方便后续处理，先预置两个点，无限小点和和无限大点
	// 注意输入的start和end在范围[0, 10^9]内
	mc.points.PushBack(&point{pos: -1, deep: 0})
	mc.points.PushBack(&point{pos: 1e9 + 1, deep: 0})
	return mc
}

func (mc *MyCalendarThree) Book(start int, end int) int {
	var startNode, endNode *list.Element
	// 插入起始点，如果已经存在则不插入
	for e := mc.points.Front(); e.Next() != nil; e = e.Next() {
		p := e.Value.(*point)
		if start == p.pos {
			startNode = e
			break
		}
		// 插入点，注意其深度暂时和其前驱点深度一致
		nextP := e.Next().Value.(*point)
		if start > p.pos && start < nextP.pos {
			p := &point{pos: start, deep: p.deep}
			startNode = mc.points.InsertAfter(p, e)
			break
		}
	}
	// 插入结束点，如果已经存在则不插入
	for e := mc.points.Back(); e.Prev() != nil; e = e.Prev() {
		p := e.Value.(*point)
		if end == p.pos {
			endNode = e
			break
		}
		prevP := e.Prev().Value.(*point)
		if end < p.pos && end > prevP.pos {
			p := &point{pos: end, deep: prevP.deep}
			endNode = mc.points.InsertBefore(p, e)
			break
		}
	}
	// 对于起始和结束点之间的所有点，深度都加一。
	for e := startNode; e != endNode && e != nil; e = e.Next() {
		p := e.Value.(*point)
		p.deep++
		mc.k = max(mc.k, p.deep)
	}
	return mc.k
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
[参考实现](iii/d.go)