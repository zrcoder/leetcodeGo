/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package bisect_squares_lcci

import "math"

/*
面试题 16.13. 平分正方形 https://leetcode-cn.com/problems/bisect-squares-lcci/
给定两个正方形及一个二维平面。请找出将这两个正方形分割成两半的一条直线。假设正方形顶边和底边与 x 轴平行。

每个正方形的数据square包含3个数值，正方形的左下顶点坐标[X,Y] = [square[0],square[1]]，以及正方形的边长square[2]。
所求直线穿过两个正方形会形成4个交点，请返回4个交点形成线段的两端点坐标
（两个端点即为4个交点中距离最远的2个点，这2个点所连成的线段一定会穿过另外2个交点）。
2个端点坐标[X1,Y1]和[X2,Y2]的返回格式为{X1,Y1,X2,Y2}，要求若X1 != X2，需保证X1 < X2，否则需保证Y1 <= Y2。

若同时有多条直线满足要求，则选择斜率最大的一条计算并返回（与Y轴平行的直线视为斜率无穷大）。

示例：

输入：
square1 = {-1, -1, 2}
square2 = {0, -1, 2}
输出： {-1,0,2,0}
解释： 直线 y = 0 能将两个正方形同时分为等面积的两部分，返回的两线段端点为[-1,0]和[2,0]
提示：

square.length == 3
square[2] > 0
*/

/*
要均分两个正方形，则必须经过两个正方形的中心`o1`,`o2`
对直线`o1o2`与坐标轴平行的情况，可简单计算返回
其他情况，根据斜率可以判断最终的结果应该是直线`o1o2`与上下边界的交点还是与左右边界的交点
计算交点可以用两点式或点斜式
*/
func cutSquares(square1 []int, square2 []int) []float64 {
	o1x := float64(square1[0]) + float64(square1[2])/2
	o1y := float64(square1[1]) + float64(square1[2])/2
	o2x := float64(square2[0]) + float64(square2[2])/2
	o2y := float64(square2[1]) + float64(square2[2])/2
	minX := min(square1[0], square2[0])
	maxX := max(square1[0]+square1[2], square2[0]+square2[2])
	minY := min(square1[1], square2[1])
	maxY := max(square1[1]+square1[2], square2[1]+square2[2])
	if equal(o1x, o2x) {
		return []float64{o1x, minY, o1x, maxY}
	}
	if equal(o1y, o2y) {
		return []float64{minX, o1y, maxX, o2y}
	}
	k := (o1y - o2y) / (o1x - o2x)
	// 与上下两边交
	//由两点式： (x-x1)(y2-y1)=(y-y1)(x2-x1)知，y=y0时，x=(y0-y1)(x2-x1)/(y2-y1) + x1
	if k > 1 { // 左下右上走势
		return []float64{(minY-o1y)*(o2x-o1x)/(o2y-o1y) + o1x, minY, (maxY-o1y)*(o2x-o1x)/(o2y-o1y) + o1x, maxY}
	}
	if k < -1 { // 左上右下走势
		return []float64{(maxY-o1y)*(o2x-o1x)/(o2y-o1y) + o1x, maxY, (minY-o1y)*(o2x-o1x)/(o2y-o1y) + o1x, minY}
	}
	// 与左右两边交
	//由两点式： (x-x1)(y2-y1)=(y-y1)(x2-x1)知，x=x0时， y=(x0-x1)(y2-y1)/(x2-x1) + y1
	return []float64{minX, (minX-o1x)*(o2y-o1y)/(o2x-o1x) + o1y, maxX, (maxX-o1x)*(o2y-o1y)/(o2x-o1x) + o1y}
}

/*
最后计算交点用了两点式，也可以用点斜式：
	// 点斜式： y - y2 = k*(x-x2)
	if k > 1 {
		return []float64{(minY-o2y)/k+o2x, minY, (maxY-o2y)/k+o2x, maxY}
	}
	if k < -1 {
		return []float64{(maxY-o2y)/k+o2x, maxY, (minY-o2y)/k+o2x, minY}
	}
	return []float64{minX, o2y + k*(minX-o2x), maxX, o2y + k*(maxX-o2x)}
*/

func min(a, b int) float64 {
	if a < b {
		return float64(a)
	}
	return float64(b)
}
func max(a, b int) float64 {
	if a > b {
		return float64(a)
	}
	return float64(b)
}
func equal(a, b float64) bool {
	return math.Abs(a-b) < 1e-6
}
