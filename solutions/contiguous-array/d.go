package contiguous_array

/*
525. 连续数组
给定一个二进制数组, 找到含有相同数量的 0 和 1 的最长连续子数组（的长度）。

示例 1:

输入: [0,1]
输出: 2
说明: [0, 1] 是具有相同数量0和1的最长连续子数组。
示例 2:

输入: [0,1,0]
输出: 2
说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。

注意: 给定的二进制数组的长度不会超过50000。
*/
// 暴力解法
func findMaxLength0(nums []int) int {
	var res int
	for start := 0; start < len(nums); start++ {
		var zeros, ones int
		for end := start; end < len(nums); end++ {
			if nums[end] == 0 {
				zeros++
			} else {
				ones++
			}
			if zeros == ones && end-start+1 > res {
				res = end - start + 1
			}
		}
	}
	return res
}

/*
数组里只有两种元素，不妨记录其中一个为 -1， 另一个为1，
则连续子数组两元素个数相等 <=> 该连续数组和为0
可以用上前缀和技巧
*/
func findMaxLength(nums []int) int {
	indices := map[int]int{0: -1}
	var res, prefixSum int
	for i, v := range nums {
		if v == 0 {
			v = -1
		}
		prefixSum += v
		if j, ok := indices[prefixSum]; !ok {
			indices[prefixSum] = i
		} else if i-j > res {
			res = i - j
		}
	}
	return res
}
