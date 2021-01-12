/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package trapping_rain_water

/*
42. 接雨水 https://leetcode-cn.com/problems/trapping-rain-water

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:
输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
*/
/*
左右指针向中间凑，对于左右的峰值，因为短板效应，哪边小处理哪边
假设处理的是左边，用一个变量记录峰值， 在元素小于等于峰值的情况下左指针一直右移且累加雨水面积（即峰值与元素值的差值）
如果出现新的峰值，更新记录峰值的变量，且和右边峰值比较决定处理哪边
右边的处理同理
时间复杂度O(n), 空间复杂度O(1)
*/
func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	left, right := 0, n-1
	leftPeek, rightPeek := 0, 0
	sum := 0
	for left < right {
		leftVal, rightVal := height[left], height[right]
		if leftVal < rightVal { // 处理左侧
			if leftVal >= leftPeek {
				leftPeek = leftVal
			} else {
				sum += leftPeek - leftVal
			}
			left++
		} else { // 处理右侧
			if rightVal >= rightPeek {
				rightPeek = rightVal
			} else {
				sum += rightPeek - rightVal
			}
			right--
		}
	}
	return sum
}

/*
借助栈的解法
https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode

遍历数组时维护一个栈，记录可能存水的条形块的索引。
如果当前条形块小于或等于栈顶索引对应的条形块，将条形块的索引入栈，意思是当前的条形块被栈中的前一个条形块界定。
如果当前条形块大于栈顶索引对应的条形块，可以确定栈顶的条形块被当前条形块和栈的前一个条形块界定，因此可以弹出栈顶元素并且累加答案

时间复杂度：O(n)。
单次遍历O(n) ，每个条形块最多访问两次（由于栈的弹入和弹出），并且弹入和弹出栈都是 O(1)的。
空间复杂度：O(n)。 栈最多在阶梯型或平坦型条形块结构中占用 O(n)的空间。
*/
func trap1(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	sum := 0
	var stack []int // 记录可能存水的条形块的索引
	for i, v := range height {
		for len(stack) > 0 && v > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			newTop := stack[len(stack)-1]
			distance := i - newTop - 1
			boundHeight := min(v, height[newTop]) - height[top]
			sum += distance * boundHeight
		}
		stack = append(stack, i)
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
