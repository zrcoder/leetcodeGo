package best_sightseeing_pair

/*
1014. 最佳观光组合
给定正整数数组 A，A[i] 表示第 i 个观光景点的评分，并且两个景点 i 和 j 之间的距离为 j - i。
一对景点（i < j）组成的观光组合的得分为（A[i] + A[j] + i - j）：景点的评分之和减去它们两者之间的距离。
返回一对观光景点能取得的最高分。

示例：
输入：[8,1,5,2,6]
输出：11
解释：i = 0, j = 2, A[i] + A[j] + i - j = 8 + 5 + 0 - 2 = 11

提示：
2 <= A.length <= 50000
1 <= A[i] <= 1000
*/
/*
朴素实现非常容易，可惜是O(n^2)的复杂度
从得分公式分析，看看能不能优化：
point == A[i] + A[j] + i - j == (A[i] + i) + (A[j] + j)
对于子数组 A[0:j+1], 最后一个元素 A[j] 的得分实际就是 max(A[i]+i) + A[j] -j(其中0 <= i < j)
A[j] - j 确定，max(A[i]+i) 在之前遍历的时候可以统计
这样可以优化为一次遍历，实际复杂度下降到O(n)
*/
func maxScoreSightseeingPair(A []int) int {
	result, maxLeft := 0, 0
	for j, v := range A {
		point := maxLeft + v - j
		result = max(result, point)
		maxLeft = max(maxLeft, v+j)
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
另一个思路：
注意到题目输入数据范围限制了是1000，所以可以直接用影响范围来优化，
即A[i]的影响范围仅限于i+A[i]，距离再远一点，A[i]的影响值就是0了
*/
func maxScoreSightseeingPair1(A []int) int {
	start, result := 0, 0
	for i := 1; i < len(A); i++ {
		for start+A[start] < i {
			start++
		}
		for j := start; j < i; j++ {
			point := A[i] + A[j] + j - i
			result = max(result, point)
		}
	}
	return result
}
