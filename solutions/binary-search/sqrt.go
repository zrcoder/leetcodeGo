/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package binary_search

func mySqrt(x int) int {
	left, right := 1, x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		}
		if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func mySqrt1(x int) int {
	left, right := 1, x+1
	for left < right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		}
		if mid*mid < x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
