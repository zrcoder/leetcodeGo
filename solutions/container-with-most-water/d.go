package water

/*
11. 盛最多水的容器
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例：

输入：[1,8,6,2,5,4,8,3,7]
输出：49
*/
/*
* 朴素解法
两层循环遍历，求每两个柱子所能盛放的水量并实时更新结果，时间复杂度是`O(n^2)`
* 贪心策略
用两个指针向中间凑。每次计算两个柱子所能盛放的水量，实时更新结果。
如果左指针代表的柱子高度较小，左指针加1，反之，右指针减1,
如果左右柱子一样高，不论左指针右移还是右指针左移，围成的面积都会比当前小，可以将两个指针同时向中间移动一步
只需遍历一遍数组，时间复杂度降为`O(n)`
*/
func maxArea(height []int) int {
	result := 0
	left, right := 0, len(height)-1
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		result = max(result, area)
		switch {
		case height[left] == height[right]:
			left++
			right--
		case height[left] < height[right]:
			left++
		default:
			right--
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
