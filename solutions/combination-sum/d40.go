package combination_sum

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var cur []int
	var backtrack func(t, start int)
	backtrack = func(t, start int) {
		if t == 0 {
			tmp := make([]int, len(cur))
			_ = copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		// 1.常规回溯写法

		if t < 0 || start == len(candidates) {
			return
		}
		// 选择 start 处的元素
		cur = append(cur, candidates[start])
		backtrack(t-candidates[start], start+1)
		cur = cur[:len(cur)-1]
		// 不选择 start 处的元素
		// 也不能选择紧跟 start 后与 start 处元素相同的元素
		i := start+1
		for i < len(candidates) && candidates[i] == candidates[start] {
			i++
		}
		backtrack(t, i)

		// 2.另一个写法

		//for i := start; i < len(candidates); i++ {
		//	if t - candidates[i] < 0 {
		//		return
		//	}
		//	if i > start && candidates[i] == candidates[i-1] {
		//		continue
		//	}
		//	cur = append(cur, candidates[i])
		//	backtrack(t-candidates[i], i+1)
		//	cur = cur[:len(cur)-1]
		//}
	}
	sort.Ints(candidates)
	backtrack(target, 0)
	return res
}