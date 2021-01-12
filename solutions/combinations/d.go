package combinations

/*
77. 组合
https://leetcode-cn.com/problems/combinations

给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入:  n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/
/*
与 78.子集类似，但可以做个剪枝
*/
func combine(n int, k int) [][]int {
	var res [][]int
	if k < 1 || k > n {
		return res
	}
	cur := make([]int, 0, k)
	var backtrack func(int)
	backtrack = func(i int) {
		// 剪枝
		if len(cur)+n-i+1 < k {
			return
		}
		if len(cur) == k {
			tmp := make([]int, k)
			_ = copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		cur = append(cur, i)
		backtrack(i + 1)
		cur = cur[:len(cur)-1]

		backtrack(i + 1)
	}
	backtrack(1)
	return res
}
