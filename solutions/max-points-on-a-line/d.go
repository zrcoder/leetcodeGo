package line

/*
149. 直线上最多的点数 https://leetcode-cn.com/problems/max-points-on-a-line/

给定一个二维平面，平面上有 n 个点，求最多有多少个点在同一条直线上。

示例 1:
输入: [[1,1],[2,2],[3,3]]
输出: 3
解释:
^
|
|        o
|     o
|  o
+------------->
0  1  2  3  4

示例 2:
输入: [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
输出: 4
解释:
^
|
|  o
|     o        o
|        o
|  o        o
+------------------->
0  1  2  3  4  5  6
*/
/*
两层循环，保证所有点两两相遇
外层循环索引i从0开始，内存循环索引j从i+1开始
对于点points[i]，借助一个哈希表统计其他点与其形成的直线斜率即可，键为斜率，值为这样的点的个数
可以用float64类型表示斜率，且将两点x坐标相同的情况记为斜率math.MaxFloat64
但是浮点数有精度损失，作为哈希表的key会不太准确，实际测试有一个用例失败
一个可行的方法是用分数表示斜率，且为了能区分斜率是否相同，需要把分数化简
这时候要注意的是在计算△x和△y的时候，可能发生溢出，这里可以用big包；就这个问题，输入的点坐标其实都是int32，可以将其转换成int64
*/
func maxPoints(points [][]int) int {
	n := len(points)
	if n < 3 {
		return n
	}

	res := 0
	for i := 0; i < n-1; i++ {
		m := map[Q]int{}
		same := 0
		cnt := 0
		for j := i + 1; j < n; j++ {
			if isSame(points[i], points[j]) {
				same++
				continue
			}
			k := getK(points[i], points[j])
			m[k]++
			cnt = max(cnt, m[k])
		}
		res = max(res, cnt+1+same)
	}
	return res
}

func isSame(p1, p2 []int) bool {
	return p1[0] == p2[0] && p1[1] == p2[1]
}

func getK(p1, p2 []int) Q {
	if p1[0] == p2[0] {
		return Q{1, 0}
	}
	if p1[1] == p2[1] {
		return Q{0, 1}
	}
	dx := int64(p1[0]) - int64(p2[0])
	dy := int64(p1[1]) - int64(p2[1])
	d := gcd(dx, dy)
	return Q{dx / d, dy / d}
}

// 有理数
type Q struct {
	m, n int64 // 分子分母
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
