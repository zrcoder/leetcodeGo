package combination_sum

import (
	"sort"
	"strconv"
	"strings"
)

// 回溯， 参考 77. 组合、 78. 子集

// 1，需要去重
func combinationSum0(candidates []int, target int) [][]int {
	var res [][]int
	var cur []int
	var dfs func(t int)
	dfs = func(t int) {
		if t < 0 {
			return
		}
		if t == 0 {
			res = append(res, copySlice(cur))
			return
		}
		for _, v := range candidates {
			cur = append(cur, v)
			dfs(t - v)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(target)

	// 出现了如 [2,2,3] 和 [2,3,2] 这样重复的结果，去重
	for _, v := range res {
		sort.Ints(v)
	}
	set := map[string]bool{}
	var ans [][]int
	for _, v := range res {
		key := hash(v)
		if !set[key] {
			ans = append(ans, v)
			set[key] = true
		}
	}
	return ans
}

func hash(s []int) string {
	buf := strings.Builder{}
	for _, v := range s {
		buf.WriteString(strconv.Itoa(v))
		buf.WriteString(",")
	}
	return buf.String()
}

// 分析重复的原因并修正
// https://leetcode-cn.com/problems/combination-sum/solution/hui-su-suan-fa-jian-zhi-python-dai-ma-java-dai-m-2
func combinationSum1(candidates []int, target int) [][]int {
	var res [][]int
	var cur []int
	var dfs func(t, start int)
	dfs = func(t, start int) {
		if t < 0 {
			return
		}
		if t == 0 {
			res = append(res, copySlice(cur))
			return
		}
		// 从 start 开始，而不是从 0 开始，防止重复的组合出现
		for j := start; j < len(candidates); j++ {
			cur = append(cur, candidates[j])
			dfs(t-candidates[j], j)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(target, 0)
	return res
}

// 通用穷举的写法
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var cur []int
	var dfs func(t, i int)
	dfs = func(t, i int) {
		if t < 0 || i == len(candidates) {
			return
		}
		if t == 0 {
			res = append(res, copySlice(cur))
			return
		}

		// 使用 i 处元素
		cur = append(cur, candidates[i])
		// 元素可无限制重复使用，这里 i 不加1
		dfs(t-candidates[i], i)
		cur = cur[:len(cur)-1]

		// 不使用 i 处元素，这里 i + 1
		dfs(t, i+1)
	}
	dfs(target, 0)
	return res
}

func copySlice(s []int) []int {
	res := make([]int, len(s))
	copy(res, s)
	return res
}
