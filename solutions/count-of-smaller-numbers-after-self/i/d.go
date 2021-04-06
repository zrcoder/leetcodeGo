/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package i

/*
面试题51. 数组中的逆序对 https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
*/

var count int

func reversePairs(nums []int) int {
	count = 0
	mergeSort(nums)
	return count
}

func mergeSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	mid := len(nums) / 2
	left := append([]int{}, nums[:mid]...)
	right := append([]int{}, nums[mid:]...)
	mergeSort(left)
	mergeSort(right)
	merge(left, right, nums)
}

func merge(left, right, nums []int) {
	var i, j int
	for k := 0; k < len(nums); k++ {
		if j == len(right) || i < len(left) && left[i] <= right[j] {
			count += j // left[i] 要比right[0:j]共j个元素大
			nums[k] = left[i]
			i++
		} else {
			nums[k] = right[j]
			j++
		}
	}
}
