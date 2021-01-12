/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package largest_rectangle_in_histogram

import "container/list"

/*
84. 柱状图中最大的矩形 https://leetcode-cn.com/problems/largest-rectangle-in-histogram/

给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。
以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。

示例:
输入: [2,1,5,6,2,3]
输出: 10
*/
/*
朴素实现1，遍历所有宽度形成的矩形，如在[i,j]形成的矩形，面积 = 宽度 * min(heights[i:j+1]

时间复杂度O(n^2), 空间复杂度O(1)
有用例超时
*/
func largestRectangleArea01(heights []int) int {
	res := 0
	for i := range heights {
		for j := i; j < len(heights); j++ {
			s := (j - i + 1) * min(heights[i:j+1])
			if s > res {
				res = s
			}
		}
	}
	return res
}

/*
朴素实现2， 遍历所有高度，对每个高度向左右扩展，直到到达边界或高度小于当前高度
时空复杂度同上, 实测AC了~~
*/
func largestRectangleArea02(heights []int) int {
	res := 0
	for i, h := range heights {
		width := 0
		for left := i; left >= 1 && heights[left-1] >= h; left-- {
			width++
		}
		for right := i; right < len(heights)-1 && heights[right+1] >= h; right++ {
			width++
		}
		res = max(res, width*h)
	}
	return res
}

/*
基于朴素实现2，借助单调栈来找每个位置左侧/右侧位置最近且高度小于当前位置高度的位置
时间复杂度O(n),空间复杂度O(n)
*/
func largestRectangleArea(heights []int) int {
	left, right := calLeft(heights), calRight(heights)
	result := 0
	for i, v := range heights {
		result = max(result, (right[i]-left[i]-1)*v)
	}
	return result
}

// 找到每个位置左侧距离最近且高度小于当前位置高度的位置
func calLeft(heights []int) []int {
	r := make([]int, len(heights))
	stack := list.New()
	for i, v := range heights {
		for stack.Len() > 0 && heights[stack.Back().Value.(int)] >= v {
			_ = stack.Remove(stack.Back())
		}
		if stack.Len() == 0 {
			r[i] = -1
		} else {
			r[i] = stack.Back().Value.(int)
		}
		stack.PushBack(i)
	}
	return r
}

// 找到每个位置右侧距离最近且高度小于当前位置的位置
func calRight(heights []int) []int {
	r := make([]int, len(heights))
	stack := list.New()
	for i := len(heights) - 1; i >= 0; i-- {
		for stack.Len() > 0 && heights[stack.Back().Value.(int)] >= heights[i] {
			_ = stack.Remove(stack.Back())
		}
		if stack.Len() == 0 {
			r[i] = len(heights)
		} else {
			r[i] = stack.Back().Value.(int)
		}
		stack.PushBack(i)
	}
	return r
}

func min(arr []int) int {
	r := arr[0]
	for _, v := range arr {
		if v < r {
			r = v
		}
	}
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
