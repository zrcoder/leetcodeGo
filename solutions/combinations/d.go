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
	var cur []int
	var dfs func(i int)
	dfs = func(i int) {
		if len(cur)+n-i+1 < k {
			return
		}
		if len(cur) == k {
			res = append(res, copySlice(cur))
			return
		}
		dfs(i + 1)
		cur = append(cur, i)
		dfs(i + 1)
		cur = cur[:len(cur)-1]
	}
	dfs(1)
	return res
}

func copySlice(s []int) []int {
	res := make([]int, len(s))
	copy(res, s)
	return res
}
