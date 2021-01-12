/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package find_smallest_common_element_in_all_rows

import "sort"

/*
1198. 找出所有行中最小公共元素
https://leetcode-cn.com/problems/find-smallest-common-element-in-all-rows

给你一个矩阵 mat，其中每一行的元素都已经按 递增 顺序排好了。请你帮忙找出在所有这些行中 最小的公共元素。

如果矩阵中没有这样的公共元素，就请返回 -1。



示例：

输入：mat = [[1,2,3,4,5],[2,4,5,8,10],[3,5,7,9,11],[1,3,5,7,9]]
输出：5


提示：

1 <= mat.length, mat[i].length <= 500
1 <= mat[i][j] <= 10^4
mat[i] 已按递增顺序排列。
*/
/*
遍历第一行，看其他行是否有第一行的元素
时间复杂度O(m*n*lgn),m、n分别为行数、列数
空间复杂度O(1)
*/
func smallestCommonElement(mat [][]int) int {
	for _, v := range mat[0] {
		ok := true
		for i := 1; i < len(mat); i++ {
			if !has(mat[i], v) {
				ok = false
				break
			}
		}
		if ok {
			return v
		}
	}
	return -1
}

func has(nums []int, target int) bool {
	i := sort.SearchInts(nums, target) //O(lgn)
	return i < len(nums) && nums[i] == target
}

/*
用一个额外的数组 pos 记录当前每一行元素的索引
对于当前元素 cur，如果每行都含有，即其个数为行数则找到了结果
*/
func smallestCommonElement1(mat [][]int) int {
	n, m := len(mat), len(mat[0])
	pos := make([]int, n)
	cur, cnt := -1, 0
	for {
		for r, row := range mat {
			i := pos[r]
			k := sort.Search(m-i, func(j int) bool {
				return row[i:][j] >= cur
			})
			i += k
			if i == m {
				return -1
			}
			if cur == row[i] {
				cnt++
				if cnt == n {
					return cur
				}
			} else {
				cnt = 1
			}
			cur = row[i]
			pos[r] = i
		}
	}
}
