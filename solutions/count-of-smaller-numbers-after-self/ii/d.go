/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package ii

/*
[315. 计算右侧小于当前元素的个数](https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/)
*/
type pair struct {
	val, index int
}

func countSmaller(nums []int) []int {
	n := len(nums)
	pairs := make([]pair, n) // 记录每个元素的值和索引,以免在排序过程中打乱顺序
	for i, v := range nums {
		pairs[i] = pair{val: v, index: i}
	}
	count := make([]int, n)
	mergeSort(pairs, count)
	return count
}

func mergeSort(pairs []pair, count []int) {
	if len(pairs) < 2 {
		return
	}
	mid := len(pairs) / 2
	left := make([]pair, mid)
	right := make([]pair, len(pairs)-mid)
	_ = copy(left, pairs[:mid])
	_ = copy(right, pairs[mid:])
	mergeSort(left, count)
	mergeSort(right, count)
	merge(left, right, pairs, count)
}

func merge(left, right, pairs []pair, count []int) {
	var i, j, k int
	for ; i < len(left) && j < len(right); k++ {
		if left[i].val <= right[j].val {
			count[left[i].index] += j // left[i]的值要比 right[0:j]共j个值大
			pairs[k] = left[i]
			i++
		} else {
			pairs[k] = right[j]
			j++
		}
	}
	for ; i < len(left); i, k = i+1, k+1 {
		count[left[i].index] += j // 左侧剩余的元素同样要比j个（等于len（right））right部分元素大
		pairs[k] = left[i]
	}
	for ; j < len(right); j, k = j+1, k+1 {
		pairs[k] = right[j]
	}
}
