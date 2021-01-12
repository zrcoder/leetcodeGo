package number_of_ships_in_a_rectangle

/*
1274. 矩形内船只的数目
https://leetcode-cn.com/problems/number-of-ships-in-a-rectangle

(此题是 交互式问题 )

在用笛卡尔坐标系表示的二维海平面上，有一些船。每一艘船都在一个整数点上，且每一个整数点最多只有 1 艘船。

有一个函数 Sea.hasShips(topRight, bottomLeft) ，输入参数为右上角和左下角两个点的坐标，
当且仅当这两个点所表示的矩形区域（包含边界）内至少有一艘船时，这个函数才返回 true ，否则返回 false 。

给你矩形的右上角 topRight 和左下角 bottomLeft 的坐标，请你返回此矩形内船只的数目。题目保证矩形内 至多只有 10 艘船。

调用函数 hasShips 超过400次 的提交将被判为 错误答案（Wrong Answer） 。同时，任何尝试绕过评测系统的行为都将被取消比赛资格。

示例：

输入：
ships = [[1,1],[2,2],[3,3],[5,5]], topRight = [4,4], bottomLeft = [0,0]
输出：3
解释：在 [0,0] 到 [4,4] 的范围内总共有 3 艘船。


提示：

ships 数组只用于评测系统内部初始化。你无法得知 ships 的信息，所以只能通过调用 hasShips 接口来求解。
0 <= bottomLeft[0] <= topRight[0] <= 1000
0 <= bottomLeft[1] <= topRight[1] <= 1000
*/
/**
 * // This is Sea's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Sea struct {
 *     func hasShips(topRight, bottomLeft []int) bool {}
 * }
 */
type Sea interface {
	hasShips(topRight, bottomLeft []int) bool
}

/*
分治
将区域划分成四个小区域，减小问题规模
*/
func countShips(sea Sea, topRight, bottomLeft []int) int {
	x1, y1 := topRight[0], topRight[1]
	x2, y2 := bottomLeft[0], bottomLeft[1]
	if x1 < x2 || y1 < y2 || !sea.hasShips(topRight, bottomLeft) {
		return 0
	}
	if x1 == x2 && y1 == y2 {
		return 1
	}
	midX := (x1 + x2) / 2
	midY := (y1 + y2) / 2
	/*
		注意四个小区域的划分，不要把(x1, y1) - (midX, midY)作为一个子区域，
		比如（2，2），（1，1）区域，中点是（1，1），会导致无穷无尽的递归
		应该把(x1, y1) - (midX-1, midY-1)作为一个子区域
	*/
	return countShips(sea, []int{midX, midY}, []int{x2, y2}) +
		countShips(sea, []int{midX, y1}, []int{x2, midY + 1}) +
		countShips(sea, []int{x1, midY}, []int{midX + 1, y2}) +
		countShips(sea, []int{x1, y1}, []int{midX + 1, midY + 1})
}

/*
二分
*/
func countShips0(sea Sea, topRight, bottomLeft []int) int {
	return help(sea, topRight, bottomLeft)
}

func help(sea Sea, topRight, bottomLeft []int) int {
	x1, y1 := topRight[0], topRight[1]
	x2, y2 := bottomLeft[0], bottomLeft[1]
	if x1 < x2 || y1 < y2 || !sea.hasShips(topRight, bottomLeft) {
		return 0
	}
	if x1 == x2 && y1 == y2 {
		return 1
	}
	if x1 == x2 {
		midY := (y1 + y2) / 2
		return help(sea, []int{x1, midY}, []int{x1, y2}) + help(sea, []int{x1, y1}, []int{x1, midY + 1})
	}
	midX := (x1 + x2) / 2
	return help(sea, []int{midX, y1}, []int{x2, y2}) + help(sea, []int{x1, y1}, []int{midX + 1, y2})
}

/*
二分优化
如果将区域仅划分为两个小区域 A 和 B，那么当对 A 区域调用 API 返回 False 时，
可以直接断定，对 B 区域调用 API 一定会返回 True，这样就省去了一次 API 的调用。
*/
func countShips1(sea Sea, topRight, bottomLeft []int) int {
	return help1(sea, topRight, bottomLeft, false)
}

func help1(sea Sea, topRight, bottomLeft []int, claim bool) int {
	x1, y1 := topRight[0], topRight[1]
	x2, y2 := bottomLeft[0], bottomLeft[1]
	if x1 < x2 || y1 < y2 {
		return 0
	}
	if !claim && !sea.hasShips(topRight, bottomLeft) {
		return 0
	}
	if x1 == x2 && y1 == y2 {
		return 1
	}
	if x1 == x2 {
		midY := (y1 + y2) / 2
		a := help1(sea, []int{x1, midY}, []int{x1, y2}, false)
		return a + help1(sea, []int{x1, y1}, []int{x1, midY + 1}, a == 0)
	}
	midX := (x1 + x2) / 2
	a := help1(sea, []int{midX, y1}, []int{x2, y2}, false)
	return a + help1(sea, []int{x1, y1}, []int{midX + 1, y2}, a == 0)
}
