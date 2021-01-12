/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package matrix_traversal

import "math"

/*
解法一：
1。每一趟对角线中元素的坐标（r, c）相加的和是递增的，且正好是遍历的趟数——r代表row，c代表clounm。
第0趟：r+c == 0
第1趟：r+c == 1
第2趟：r+c == 2
第3趟：r+c == 3
第4趟：r+c == 4
。。。
且趟数上限是m+n-2（即m-1 + n-1）

2。每一趟，要么横坐标递增，纵坐标递减；要么横坐标递减，纵坐标递增。
第0趟：r，c都是从0到0
第1趟：r从0到1，c从1到0
第2趟：r从2到0，c从0到2
第3趟：r从1到2，c从2到1
第4趟：r和c都是从2到2
即偶数趟从左下到右上，r递减，c递增；奇数趟从右上到左下r递增，c递减

根据以上分析，常规遍历即可
时空复杂度都是O(m*n)，我们遍历了每个元素，且将每个元素放入一个结果数组里
*/
func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	if len(matrix[0]) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	result := make([]int, m*n)
	r, c := 0, 0
	for i := 0; i < len(result); i++ {
		result[i] = matrix[r][c]
		if (r+c)%2 == 0 { // 偶数趟，从左下到右上
			if c == n-1 {
				r++
			} else {
				c++
				r = max(r-1, 0)
			}
		} else { // 奇数趟，从右上到左下
			if r == m-1 {
				c++
			} else {
				r++
				c = max(c-1, 0)
			}
		}
	}
	return result
}

/*
解法二：

1。每一趟对角线中元素的坐标（r, c）相加的和是递增的，且正好是遍历的趟数——r代表row，c代表clounm。
第0趟：r+c == 0
第1趟：r+c == 1
第2趟：r+c == 2
第3趟：r+c == 3
第4趟：r+c == 4
。。。
且趟数上限是m+n-2（即m-1 + n-1）

2。每一趟，要么横坐标递增，纵坐标递减；要么横坐标递减，纵坐标递增。
第0趟：r，c都是从0到0
第1趟：r从0到1，c从1到0
第2趟：r从2到0，c从0到2
第3趟：r从1到2，c从2到1
第4趟：r和c都是从2到2
即偶数趟r递减，c递增；奇数趟r递增，c递减

3。确定初始值。
假设是第time趟，递减的坐标初始值要尽量大，应为time，但如果time超过上限，则初始值应为上限；另一个坐标的初始值则是time-该坐标初始值：
偶数趟r递减，r初始值为time，但如果 time>m-1, 初始值则为m-1，即min(time, m-1)； 而c=time-r
奇数趟c递减，c初始值为time，但如果 time>n-1, 则初始值为n-1，即min(time, n-1)； 而r=time-c

时空复杂度都是O(m*n)，我们遍历了每个元素，且将每个元素放入一个结果数组里
*/
func findDiagonalOrder1(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	if len(matrix[0]) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	result := make([]int, m*n)
	i := 0
	for time := 0; time <= m+n-2; time++ {
		if time%2 == 0 { // 偶数趟，r递减
			r := min(time, m-1)
			c := time - r
			for r >= 0 && c < n {
				result[i] = matrix[r][c]
				i++
				r--
				c++
			}
		} else { // 奇数趟, c递减
			c := min(time, n-1)
			r := time - c
			for c >= 0 && r < m {
				result[i] = matrix[r][c]
				i++
				c--
				r++
			}
		}
	}
	return result
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
