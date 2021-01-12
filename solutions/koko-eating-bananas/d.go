/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package koko_eating_bananas

import (
	"sort"

	s "github.com/zrcoder/leetcodeGo/util/sort"
)

/*
875. 爱吃香蕉的珂珂 https://leetcode-cn.com/problems/koko-eating-bananas
珂珂喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。

珂珂可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。
如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。

珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。

返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。

示例 1：
输入: piles = [3,6,7,11], H = 8
输出: 4

示例 2：
输入: piles = [30,11,23,4,20], H = 5
输出: 30
示例 3：

输入: piles = [30,11,23,4,20], H = 6
输出: 23


提示：

1 <= piles.length <= 10^4
piles.length <= H <= 10^9
1 <= piles[i] <= 10^9
*/
/*朴素实现
O(n*max),超时
n为piles长度，max为piles最大元素
*/
func minEatingSpeed1(piles []int, H int) int {
	if H < len(piles) {
		return -1
	}
	max := -1
	for _, v := range piles {
		if v > max {
			max = v
		}
	}
	if H == len(piles) {
		return max
	}

	for k := 1; k < max; k++ {
		if calCost(piles, k) <= H {
			return k
		}
	}
	return -1
}

/*
二分法实现
O(n*lg(max)), n为piles长度，max为piles最大元素
*/
func minEatingSpeed2(piles []int, H int) int {
	n := len(piles)
	if H < n {
		return -1
	}
	max := -1
	for _, v := range piles {
		if v > max {
			max = v
		}
	}
	if H == n {
		return max
	}
	left, right := 1, max // 注意取值为max的情况以及在前边返回，所以不再包含max
	for left < right {
		mid := left + (right-left)/2
		if calCost(piles, mid) <= H { // 可能有多个值使得calCost(piles, mid) == H， 需要最左侧的那个值
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 使用标准库，减少代码量
func minEatingSpeed(piles []int, H int) int {
	n := len(piles)
	if H < n {
		return -1
	}
	max := -1
	for _, v := range piles {
		if v > max {
			max = v
		}
	}
	if H == n {
		return max
	}
	return sort.Search(max, func(i int) bool { // 注意取值为max的情况以及在前边返回，所以不再包含max
		if i == 0 {
			return false
		}
		return calCost(piles, i) <= H
	})
}

// 仿照标准库改版Search函数
func minEatingSpeed3(piles []int, H int) int {
	n := len(piles)
	if H < n {
		return -1
	}
	max := -1
	for _, v := range piles {
		if v > max {
			max = v
		}
	}
	if H == n {
		return max
	}
	return s.Search(1, max, func(i int) bool {
		return calCost(piles, i) <= H
	})
}

func calCost(piles []int, k int) int {
	sum := 0
	for _, v := range piles {
		sum += (v-1)/k + 1
		/* 另一个实现
		sum += v / k
		if v%k > 0 {
			sum += 1
		}
		*/
	}
	return sum
}
