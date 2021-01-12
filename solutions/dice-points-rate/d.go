/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package dice_points_rate

import "math"

/*
面试题60. n个骰子的点数 https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof/
把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。

你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。

示例 1:
输入: 1
输出: [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]

示例 2:
输入: 2
输出: [0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0.05556,0.02778]

限制：
1 <= n <= 11
*/

/*
这题竟然是简单级，其实还蛮困难的~

首先，一个骰子可能出现1-6点里的一种情况，共n个骰子，所有情况共有6^n种，
最小的点数是所有骰子都是1点，即n；最大的点数是所有骰子都是6点，即6n
所以所有点数在区间[n, 6n]中, 且中间所有的点都可能投出，所以个数为6n-n+1 = 5n+1

主要的工作是计算出各个点数可能出现多少次：
比如共2个骰子，点数3会出现的情况有多少种？
第一个骰子1点，第二个骰子2点或者第一个骰子2点，第二个骰子1点；共两种情况
可以模拟一个骰子一个骰子地掷
假设掷到第i个骰子， 前i-1个骰子掷出来的每个结果需要分别加上1-6的点数
最终，每个点数出现的概率即该点数出现的次数除以所有点数出现的情况的总数
*/
func twoSum(n int) []float64 {
	const (
		minPoint = 1
		maxPoint = 6
	)
	count := make(map[int]int, 0) // 统计每个点数出现的次数
	// 模拟掷第一个骰子
	for i := minPoint; i <= maxPoint; i++ { // 1-6每个点数都可能出现
		count[i] = 1
	}
	// 模拟掷后边的n-1个骰子
	for i := 1; i < n; i++ {
		nextCount := map[int]int{}
		for k, v := range count {
			for i := minPoint; i <= maxPoint; i++ {
				nextCount[k+i] += v
			}
		}
		count = nextCount
	}

	sum := math.Pow(float64(maxPoint), float64(n)) // 或者遍历count，累加其value也可得到
	result := make([]float64, n*maxPoint-n+1)
	for k, v := range count {
		result[k-n] = float64(v) / sum
	}
	return result
}
