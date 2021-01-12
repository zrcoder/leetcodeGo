package combination_sum

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var cur []int
	max := 9
	if n < max {
		max = n
	}
	var backtrack func(t, num int)
	backtrack = func(t, num int) {
		if t == 0 && len(cur) == k {
			tmp := make([]int, len(cur))
			_ = copy(tmp, cur)
			res = append(res, tmp)
			return
		}
		if t <= 0 || num > max {
			return
		}

		cur = append(cur, num)
		backtrack(t-num, num+1)
		cur = cur[:len(cur)-1]

		backtrack(t, num+1)
	}
	backtrack(n, 1)
	return res
}
