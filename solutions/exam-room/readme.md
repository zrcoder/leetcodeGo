# [855. 考场就座](https://leetcode-cn.com/problems/exam-room)

```
在考场里，一排有 N 个座位，分别编号为 0, 1, 2, ..., N-1 。

当学生进入考场后，他必须坐在能够使他与离他最近的人之间的距离达到最大化的座位上。
如果有多个这样的座位，他会坐在编号最小的座位上。(另外，如果考场里没有人，那么学生就坐在 0 号座位上。)

返回 ExamRoom(int N) 类，它有两个公开的函数：其中，函数 ExamRoom.seat() 会返回一个 int （整型数据），代表学生坐的位置；
函数 ExamRoom.leave(int p) 代表坐在座位 p 上的学生现在离开了考场。每次调用 ExamRoom.leave(p) 时都保证有学生坐在座位 p 上。

示例：

输入：["ExamRoom","seat","seat","seat","seat","leave","seat"], [[10],[],[],[],[],[4],[]]
输出：[null,0,9,4,2,null,5]
解释：
ExamRoom(10) -> null
seat() -> 0，没有人在考场里，那么学生坐在 0 号座位上。
seat() -> 9，学生最后坐在 9 号座位上。
seat() -> 4，学生最后坐在 4 号座位上。
seat() -> 2，学生最后坐在 2 号座位上。
leave(4) -> null
seat() -> 5，学生最后坐在 5 号座位上。

提示：

1 <= N <= 10^9
在所有的测试样例中 ExamRoom.seat() 和 ExamRoom.leave() 最多被调用 10^4 次。
保证在调用 ExamRoom.leave(p) 时有学生正坐在座位 p 上。
```

```
最直观的解法是用一个bool数组模拟这一排座位；坐了人标为true，没坐人标为false<br>
Seat方法遍历一遍数组找到目标即可，注意判断最大距离时，开头和结束位置如果没人，距离计算不用折半<br>
空间复杂度O(n)；Leave的时间复杂度O(1),Seat的时间复杂度O(n)<br>
在leetcode测试发现128个用例，只过了120个；接下来一个用例因为超过内存限制而失败<br>
数组换用bitset， 将空间复杂度减少为原来的1/8, 不过在执行到第119个用例的时候又超过了时间限制——虽然bitset的读写也是常熟级复杂度，但毕竟不如数组读写来得直接<br>
换用list，list里只存已经坐了人的座位，这样能节省空间和遍历的时间；测试通过——不过这也只是因为用例设计问题，如果所有座位都坐满人，或多数座位坐了人，list同样会超时<br>
官方题解使用有序集合（如Java的TreeSet），可以和list的实现做个对比。假设集合中已有元素个数为P：
对于Seat，TreeSeat的复杂度是 O(P+lg(P)）= O(P）， 而list的复杂度是真正的O(P）；<br>
对于Leave，TreeSet的是O(lg(P))， list的是O(P)；<br>
综合复杂度：如果Seat和Leave的操作次数相当，list的实现复杂度综合为O(2P), TreeSet是O(P+2lg(P))) —— 都可归约为O(P)
另外，可以借助一个哈希表存储，使得Leave的复杂度为O(1)， 不过这个又有额外P的空间增加了<br>
## bitset的实现（基本同数组的实现）
```

```go
type BitSet []byte

func NewBitSetWithSize(size int) BitSet {
	const byteLen = 8
	realSize := 1 + (size-1)/byteLen
	return make([]byte, realSize)
}

// Set true at the index
func (bs BitSet) Set(index int) {
	i, mask := bs.caculateInnerIndexAndMask(index)
	bs[i] |= mask
}

// Set false at the index
func (bs BitSet) Unset(index int) {
	i, mask := bs.caculateInnerIndexAndMask(index)
	bs[i] &= ^mask
}

// Returns the bool value at the index
func (bs BitSet) Get(index int) bool {
	i, mask := bs.caculateInnerIndexAndMask(index)
	return bs[i]&mask != 0
}

func (bs BitSet) caculateInnerIndexAndMask(index int) (int, byte) {
	const byteLen = 8
	return index / byteLen, 1 << uint(index%byteLen)
}

type ExamRoom struct {
	seated BitSet
	count  int
	n      int
}

func Constructor(N int) ExamRoom {
	return ExamRoom{
		seated: NewBitSetWithSize(N),
		count:  0,
		n:      N,
	}
}

func (room *ExamRoom) Seat() int {
	if room.count == 0 { // 还没有人入座，直接将0插入
		room.seated.Set(0)
		room.count++
		return 0
	}
	prevSeated := -1
	for i := 0; i < room.n; i++ {
		if room.seated.Get(i) {
			prevSeated = i
			break
		}
	}
    target := 0
	maxDist := prevSeated // 入座后距离最近的人的最大距离，当前是从位置0到第一个坐了人的位置的距离
	for i := prevSeated + 1; i < room.n; i++ {
		if !room.seated.Get(i) {
			continue
		}
		dist := (i - prevSeated) / 2 // 插入到中间后距离两边的距离
		if dist > maxDist {
			maxDist = dist
			target = prevSeated + dist
		}
		prevSeated = i
	}
	if !room.seated.Get(room.n-1) && room.n-1-prevSeated > maxDist { // 处理最后座位没坐人的情况
		target = room.n - 1
	}
	room.seated.Set(target)
	room.count++
	return target
}

func (room *ExamRoom) Leave(p int) {
	if !room.seated.Get(p) {
		return
	}
	room.seated.Unset(p)
	room.count--
}
```

## list的实现

```go
import "container/list"

type ExamRoom struct {
	seated *list.List // 装坐着同学的位置
	last   int        // 最后一个座位， 就是总座位数减一
}

func Constructor(N int) ExamRoom {
	return ExamRoom{
		seated: list.New(),
		last:   N - 1,
	}
}

func (room *ExamRoom) Seat() int {
	if room.seated.Len() == 0 { // 还没有人入座，选座位0
		room.seated.PushFront(0)
		return 0
	}
	prevSeated := room.seated.Front().Value.(int)
	targetVal := 0                           // 需要插入的座位
	maxDist := prevSeated                    // 入座后距离最近的人的最大距离，当前是从位置0到第一个坐了人的位置的距离
	targetNextElement := room.seated.Front() // 需要插入的点的后一个元素。方便找到后直接插入
	for e := room.seated.Front().Next(); e != nil; e = e.Next() {
		currSeated := e.Value.(int)
		distant := (currSeated - prevSeated) / 2 // 两点之间的最远距离
		if distant > maxDist {
			maxDist = distant
			targetNextElement = e
			targetVal = prevSeated + distant
		}
		prevSeated = currSeated
	}
	if room.last-prevSeated > maxDist { // 尾部特殊判断
		room.seated.PushBack(room.last)
		return room.last
	}
	room.seated.InsertBefore(targetVal, targetNextElement)
	return targetVal
}

func (room *ExamRoom) Leave(p int) {
	for e := room.seated.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == p {
			room.seated.Remove(e)
			return
		}
	}
	return
}
```