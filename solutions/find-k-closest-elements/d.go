/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */
package find_k_closest_elements

import (
	"math"
	"sort"
)

/*
658. 找到 K 个最接近的元素 https://leetcode-cn.com/problems/find-k-closest-elements

给定一个排序好的数组，两个整数 k 和 x，从数组中找到最靠近 x（两数之差最小）的 k 个数。
返回的结果必须要是按升序排好的。如果有两个数与 x 的差值一样，优先选择数值较小的那个数。

示例 1:

输入: [1,2,3,4,5], k=4, x=3
输出: [1,2,3,4]


示例 2:

输入: [1,2,3,4,5], k=4, x=-1
输出: [1,2,3,4]


说明:

k 的值为正数，且总是小于给定排序数组的长度。
数组不为空，且长度不超过 104
数组里的每个元素与 x 的绝对值不超过 104
*/

// 用二分法找到应该插入x的位置；根据k确定结果的左右边界，并通过一个循环不断比较左端还是右端更接近x，缩小目标直到个数为k
func findClosestElements(arr []int, k int, x int) []int {
	i := sort.SearchInts(arr, x)
	var minLeft, maxRight int
	if i < len(arr) && arr[i] == x {
		minLeft, maxRight = max(0, i-k+1), min(i+k-1, len(arr)-1)
	} else {
		minLeft, maxRight = max(0, i-k), min(i+k, len(arr)-1)
	}
	for maxRight-minLeft+1 > k {
		if abs(arr[minLeft]-x) <= abs(arr[maxRight]-x) {
			maxRight--
		} else {
			minLeft++
		}
	}
	result := make([]int, k)
	for j := 0; j < k; j++ {
		result[j] = arr[minLeft+j]
	}
	return result
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
