/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package twosum

func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	for i, element := range nums {
		if j, found := index[target-element]; found {
			return []int{j, i} // not {i, j}, but {j, i}; let's think, j < i
		}
		index[element] = i
	}
	return nil
}
