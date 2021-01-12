package exam

/*
1349. 参加考试的最大学生数
https://leetcode-cn.com/problems/maximum-students-taking-exam

给你一个 m * n 的矩阵 seats 表示教室中的座位分布。如果座位是坏的（不可用），就用 '#' 表示；否则，用 '.' 表示。

学生可以看到左侧、右侧、左上、右上这四个方向上紧邻他的学生的答卷，但是看不到直接坐在他前面或者后面的学生的答卷。
请你计算并返回该考场可以容纳的一起参加考试且无法作弊的最大学生人数。

学生必须坐在状况良好的座位上。


示例 1：
输入：seats = [["#",".","#","#",".","#"],
              [".","#","#","#","#","."],
              ["#",".","#","#",".","#"]]
输出：4
解释：教师可以让 4 个学生坐在可用的座位上，这样他们就无法在考试中作弊。

示例 2：
输入：seats = [[".","#"],
              ["#","#"],
              ["#","."],
              ["#","#"],
              [".","#"]]
输出：3
解释：让所有学生坐在可用的座位上。

示例 3：
输入：seats = [["#",".",".",".","#"],
              [".","#",".","#","."],
              [".",".","#",".","."],
              [".","#",".","#","."],
              ["#",".",".",".","#"]]
输出：10
解释：让学生坐在第 1、3 和 5 列的可用座位上。

提示：
seats 只包含字符 '.' 和'#'
m == seats.length
n == seats[i].length
1 <= m <= 8
1 <= n <= 8
*/
// 用一个int来代表座位是否可用、座位上是否坐了学生，代码会简化不少
// 时间复杂度O(m*(4^n)), 题目里的m和n最大为8
func maxStudents(seats [][]byte) int {
	m, n := len(seats), len(seats[0])
	/*
		对于某行,要安排学生，n个位置共有 2^n （即1<<n）种可能性
		1表示安排落座，0表示不安排
		所有的安排状态为：

		000...000
		000...001
		...
		111...111

		其中每行为长度为n的01串
	*/
	total := 1 << n
	//  student[i]表示前i排共安排了多少学生，student[i[j]表示在第i排按照状态j安排时的结果
	//  为了记录好处理边界，开辟m+1行
	students := make([][]int, m+1)
	for i := range students {
		students[i] = make([]int, total)
	}
	badSeats := make([]int, m)
	for i := range badSeats {
		for j := 0; j < n; j++ {
			// 坏座位用1表示
			if seats[i][j] == '#' {
				badSeats[i] |= 1 << j
			}
		}
	}
	// 事先计算并缓存每种状态能安排的学生数
	stateCount := make([]int, total)
	for i := range stateCount {
		if i&(i<<1) != 0 || i&(i>>1) != 0 { // i这样安排会导致左边有人或右边有人，抛弃安排状态i
			continue
		}
		stateCount[i] = count(i)
	}
	result := 0
	for r := 0; r < m; r++ {
		// 对于第 r 排的安排方式s作判断
		for s := 0; s < total; s++ {
			// 有坏座位/左右有人
			if s&badSeats[r] != 0 || s&(s>>1) != 0 || s&(s<<1) != 0 {
				continue
			}
			for k := 0; k < total; k++ {
				// 对于每个坐了人的座位，左前和右前都没有人
				if s&(k>>1) == 0 && s&(k<<1) == 0 {
					students[r+1][s] = max(students[r+1][s], students[r][k]+stateCount[s])
					result = max(result, students[r+1][s])
				}
			}
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 对于状态s，统计其二进制中1的个数，即安排学生的个数
func count(s int) int {
	r := 0
	for s > 0 {
		s &= s - 1
		r++
	}
	return r
}
