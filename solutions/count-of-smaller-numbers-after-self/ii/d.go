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

var count []int

func countSmaller(nums []int) []int {
	n := len(nums)
	pairs := make([]pair, n) // 记录每个元素的值和索引,以免在排序过程中打乱顺序
	for i, v := range nums {
		pairs[i] = pair{val: v, index: i}
	}
	count = make([]int, n)
	mergeSort(pairs)
	return count
}

func mergeSort(pairs []pair) {
	if len(pairs) < 2 {
		return
	}
	mid := len(pairs) / 2
	left := append([]pair{}, pairs[:mid]...)
	right := append([]pair{}, pairs[mid:]...)
	mergeSort(left)
	mergeSort(right)
	merge(left, right, pairs)
}

func merge(left, right, pairs []pair) {
	var i, j int
	for k := 0; i < len(left) || j < len(right); k++ {
		if j == len(right) || i < len(left) && left[i].val <= right[j].val {
			count[left[i].index] += j // left[i]的值要比 right[0:j]共j个值大
			pairs[k] = left[i]
			i++
		} else {
			pairs[k] = right[j]
			j++
		}
	}
}
