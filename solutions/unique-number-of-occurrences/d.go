/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package unique_number_of_occurrences

import "sort"

/*
1207. 独一无二的出现次数 https://leetcode-cn.com/problems/unique-number-of-occurrences
给你一个整数数组 arr，请你帮忙统计数组中每个数的出现次数。
如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。
示例 1：
输入：arr = [1,2,2,1,1,3]
输出：true
解释：在该数组中，1 出现了 3 次，2 出现了 2 次，3 只出现了 1 次。没有两个数的出现次数相同。

示例 2：
输入：arr = [1,2]
输出：false

示例 3：
输入：arr = [-3,0,1,-3,1,1,1,-3,10,0]
输出：true
*/
/*
借助一个map统计元素个数
再借助一个set去重，判读是否有重复个数
时空复杂度都是O(n)
*/
func uniqueOccurrences(arr []int) bool {
	counts := make(map[int]int, 0)
	for _, v := range arr {
		counts[v]++
	}
	set := make(map[int]bool, 0)
	for _, count := range counts {
		if set[count] {
			return false
		}
		set[count] = true
	}
	return true
}

/*
先排序，再用二分法计算每个元素的数量（对于一个特定元素，遍历时左边界确定，右边界用二分法计算，最后计算长度）
另用一个set，不断存入计算出来的元素个数，当发现有重复的时候返回false，遍历完毕还没有重复即true
时间复杂度O(nlgn)，空间复杂度O(n)，但是要比前一个方法少申请一个map
*/
func uniqueOccurrences1(arr []int) bool {
	if len(arr) == 0 {
		return true
	}
	sort.Ints(arr)
	left := 0
	set := make(map[int]bool, 0)
	for left < len(arr) {
		tmp := arr[left:]
		count := sort.Search(len(tmp), func(i int) bool {
			return tmp[i] > arr[left]
		})
		if set[count] {
			return false
		}
		set[count] = true
		left += count
	}
	return true
}
