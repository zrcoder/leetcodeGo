/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package permutations

import "sort"

func permute(nums []int) [][]int {
	if len(nums) < 2 {
		return [][]int{nums}
	}
	var res [][]int
	for _, v := range permute(nums[:len(nums)-1]) {
		for i := 0; i <= len(v); i++ {
			t := append(append(v[:i:i], nums[len(nums)-1]), v[i:]...)
			res = append(res, t)
		}
	}
	return res
}

func permute1(nums []int) [][]int {
	n := len(nums)
	var result [][]int
	// 保持start之前的元素固定不变，将其及其之后的元素全排列
	var dfs func(int)
	dfs = func(start int) {
		if start == n {
			r := make([]int, n)
			_ = copy(r, nums)
			result = append(result, r)
			return
		}
		for i := start; i < n; i++ { // 将i及其i之后的元素全排列，注意不能漏了i
			nums[start], nums[i] = nums[i], nums[start] // 通过交换做排列
			dfs(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	dfs(0)
	return result
}

func permuteUnique1(nums []int) [][]int {
	n := len(nums)
	var res [][]int
	// 保持start之前的元素固定不变，将其及其之后的元素全排列
	var dfs func(int)
	dfs = func(start int) {
		if start == n {
			r := make([]int, n)
			_ = copy(r, nums)
			res = append(res, r)
			return
		}
		visited := make(map[int]bool, n-start)
		for i := start; i < n; i++ { // 将start及其之后的元素全排列，注意不能漏了start
			if visited[nums[i]] {
				continue
			}
			visited[nums[i]] = true
			nums[start], nums[i] = nums[i], nums[start] // 通过交换做排列
			dfs(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	dfs(0)
	return res
}

// 填空法
func permuteUnique(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	n := len(nums)
	cur := []int{}
	seen := make([]bool, n)
	var dfs func()
	dfs = func() {
		if len(cur) == n {
			r := make([]int, n)
			_ = copy(r, cur)
			res = append(res, r)
			return
		}
		for i, v := range nums {
			if seen[i] || i > 0 && !seen[i-1] && v == nums[i-1] {
				continue
			}
			cur = append(cur, v)
			seen[i] = true
			dfs()
			seen[i] = false
			cur = cur[:len(cur)-1]
		}
	}
	dfs()
	return res
}
