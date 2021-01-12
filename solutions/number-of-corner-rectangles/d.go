/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package rectangles

/*
给定一个只包含 0 和 1 的网格，找出其中角矩形的数量。
一个 角矩形 是由四个不同的在网格上的 1 形成的轴对称的矩形。注意只有角的位置才需要为 1。并且，4 个 1 需要是不同的。

示例 1：
输入：grid =
[[1, 0, 0, 1, 0],
 [0, 0, 1, 0, 1],
 [0, 0, 0, 1, 0],
 [1, 0, 1, 0, 1]]
输出：1
解释：只有一个角矩形，角的位置为 grid[1][2], grid[1][4], grid[3][2], grid[3][4]。

示例 2：
输入：grid =
[[1, 1, 1],
 [1, 1, 1],
 [1, 1, 1]]
输出：9
解释：这里有 4 个 2x2 的矩形，4 个 2x3 和 3x2 的矩形和 1 个 3x3 的矩形。

示例 3：
输入：grid =
[[1, 1, 1, 1]]
输出：0
解释：矩形必须有 4 个不同的角。

注：

网格 grid 中行和列的数目范围为 [1, 200]。
Each grid[i][j] will be either 0 or 1.
网格中 1 的个数不会超过 6000。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-corner-rectangles
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
题目没有描述清楚，根据实际测试用例，轴对称完全没有体现，像这样的竟然返回1：
[[1,1,0,1],
[0,0,0,0],
[1,0,0,1]]
强烈要求修改题目描述为：角矩形指四个角上均为1的矩形
*/

/*
暴力解
假设grid是m行n列，时间复杂度O(m^2*n^2),空间复杂度O(1)
*/
func countCornerRectangles1(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	count := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 0 {
				continue
			}
			for cc := c + 1; cc < n; cc++ {
				if grid[r][cc] == 0 {
					continue
				}
				for rr := r + 1; rr < m; rr++ {
					if grid[rr][c] == 1 && grid[rr][cc] == 1 {
						count++
					}
				}
			}
		}
	}
	return count
}

/*
暴力解改进
可以固定两行，然后在这两行里边查找，时间复杂度降为O(m^2*n), 空间复杂度仍然为O(1)
*/
func countCornerRectangles11(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	count := 0
	for r1 := 0; r1 < len(grid)-1; r1++ {
		for r2 := r1 + 1; r2 < len(grid); r2++ {
			bothOneCount := 0
			for c := 0; c < len(grid[r2]); c++ {
				if grid[r1][c] == 1 && grid[r2][c] == 1 {
					bothOneCount++
				}
			}
			/*
				假设f(n)表示r1和r2两行里有n对满足条件的1能构成的矩形数量
				有一个比较明显的递推公式： f(n) = f(n-1) + n - 1, 又f(1) = 0, 可以推出通项公式
				f(n) = n*(n-1)/2
			*/
			count += bothOneCount * (bothOneCount - 1) / 2
		}
	}
	return count
}

/*
借用一个map，记录一行中同时为1的两列（假设列为c1， c2）， key为这两列的标志，有个技巧，可以让200*c1+c2来作为key；
遍历行，对应的两个列再次出现均为1的情况则map里的值+1

假设grid是m行n列，时间复杂度O(m*n^2),空间复杂度O(n^2)
*/
func countCornerRectangles2(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	count := 0
	counter := make(map[int]int, 0)
	for _, row := range grid {
		for c1 := 0; c1 < len(row); c1++ {
			if row[c1] == 0 {
				continue
			}
			for c2 := c1 + 1; c2 < len(row); c2++ {
				if row[c2] == 0 {
					continue
				}
				pos := c1*200 + c2 // 一个简单的哈希，代表了c1和c2这两个位置; 题目约束列的上限为200
				c := counter[pos]
				count += c
				counter[pos] = c + 1
			}
		}
	}
	return count
}

/*
将上面的map改为数组

时间复杂度O(m*n^2); 实际测试比上边用map的方法快小一半的时间
空间复杂度O(n^2)
*/
func countCornerRectangles(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	n := len(grid[0])
	count := 0
	counter := make([][]int, n)
	for i := 0; i < n; i++ {
		counter[i] = make([]int, n)
	}
	for _, row := range grid {
		for c1 := 0; c1 < n; c1++ {
			if row[c1] == 0 {
				continue
			}
			for c2 := c1 + 1; c2 < n; c2++ {
				if row[c2] == 0 {
					continue
				}
				count += counter[c1][c2]
				counter[c1][c2]++
			}
		}
	}
	return count
}
