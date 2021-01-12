/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package i

/*
面试题51. 数组中的逆序对 https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
*/

func reversePairs(nums []int) int {
	var count int
	mergeSort(nums, &count)
	return count
}

func mergeSort(nums []int, count *int) {
	if len(nums) < 2 {
		return
	}
	mid := len(nums) / 2
	left := make([]int, mid)
	right := make([]int, len(nums)-mid)
	_ = copy(left, nums[:mid])
	_ = copy(right, nums[mid:])
	mergeSort(left, count)
	mergeSort(right, count)
	merge(left, right, nums, count)
}

func merge(left, right, nums []int, count *int) {
	var i, j, k int
	for ; i < len(left) && j < len(right); k++ {
		if left[i] <= right[j] {
			*count += j // left[i] 要比right[0:j]共j个元素大
			nums[k] = left[i]
			i++
		} else {
			nums[k] = right[j]
			j++
		}
	}
	for ; i < len(left); i, k = i+1, k+1 {
		*count += j // 左侧剩余的元素同样要比j个（等于len（right））right部分元素大
		nums[k] = left[i]
	}
	for ; j < len(right); j, k = j+1, k+1 {
		nums[k] = right[j]
	}
}
