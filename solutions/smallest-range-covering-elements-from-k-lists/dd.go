package smallest_range_covering_elements_from_k_lists

import "math"

/*
解法二：借助哈希表的滑档窗口解法
先统计出所有数字里的最小值 xMin 和最大值 xMax， [xMin, xMax]区间包含了所有数字，答案不会比这个区间更大
接下来可以用滑动窗口的方式遍历[xMin, xMax]区间来尝试
窗口左右边界一开始都为xMin，增加right指针，使得对于窗口[left, right]， k 个列表中的每个列表至少有一个数包含在其中；则left， right可能为一个答案
但这时候可以缩小窗口，增加左边界， 一直维持k 个列表中的每个列表至少有一个数包含在窗口中的性质，直到不再满足这个限制后，停止增加左边界，开始增加右边界

为了迅速判断是否 k 个列表中的每个列表至少有一个数包含在窗口中，可以事先用一个哈希表统计每个数字都在哪些列表出现
假设所有数字共n个， 最大数字与最小数字差为abs，时间复杂度O(max(n, k*abs))
*/
func smallestRange1(nums [][]int) []int {
	size := len(nums)
	indices := map[int][]int{}                 // 记录每个数字所在列表的索引
	xMin, xMax := math.MaxInt32, math.MinInt32 // 分别记录所有数字中的最小和最大元素
	for i, v := range nums {
		for _, x := range v {
			indices[x] = append(indices[x], i)
			xMin = min(xMin, x)
			xMax = max(xMax, x)
		}
	}
	freq := make([]int, size)
	inside := 0
	start, end := xMin, xMax
	for left, right := xMin, xMin; right < xMax; right++ {
		if len(indices[right]) == 0 { // k个列表里都不含right这个数字
			continue
		}
		for _, index := range indices[right] {
			freq[index]++
			if freq[index] == 1 { // right这个数字在列表index里出现一次了，即列表index至少有一个元素包含在窗口里了
				inside++
			}
		}
		for ; inside == size; left++ {
			if right-left < end-start {
				start, end = left, right
			}
			for _, index := range indices[left] {
				freq[index]--
				if freq[index] == 0 {
					inside--
				}
			}
		}
	}
	return []int{start, end}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
