/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package subarray_sums_divisible_by_k

/*
974. 和可被 K 整除的子数组 https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/
给定一个整数数组 A，返回其中元素之和可被 K 整除的（连续、非空）子数组的数目。

示例：

输入：A = [4,5,0,-2,-3,1], K = 5
输出：7
解释：
有 7 个子数组满足其元素之和可被 K = 5 整除：
[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]

提示：

1 <= A.length <= 30000
-10000 <= A[i] <= 10000
2 <= K <= 10000
*/
/*
朴素实现，时间复杂度O(n^2),空间复杂度O(1)
*/
func subarraysDivByK0(A []int, K int) int {
	result := 0
	for i := 0; i < len(A); i++ {
		sum := 0
		for j := i; j < len(A); j++ {
			sum += A[j]
			if sum%K == 0 {
				result++
			}
		}
	}
	return result
}

/*
怎么优化朴素实现？
对于一个数组， `........i.....j......`
假如前`i`个元素相加的和`prefixSum(i)`除以`K`的余数是`x`，前`j`个元素相加的和`prefixSum(j)`除以`K`的余数也是`x`，
考虑`i`到`j`之间的元素，不包含`i`处元素但是包含`j`处元素，其和肯定整除`K`；
可以简单理解下： `prefixSum(i) = m*K + x`, `prefixSum(j) = n*K + x`, 显然相减后消除了`x`，能被`K`整除

所以，只需要统计数组所有前缀和，出现与`K`取模相同的，结果累加即可；借助一个哈希表来统计会很方便
时间复杂度`O(n)`,  空间复杂度`O(n)`
*/
func subarraysDivByK(A []int, K int) int {
	record := map[int]int{0: 1} // 如果A中有元素能整除K， 那么相关元素单独能构成符合题意的子数组，需要统计
	prefixSum, result := 0, 0
	for _, v := range A {
		prefixSum += v
		mod := prefixSum % K
		if mod < 0 { // prefixSum可能是负数，导致mod为负数， 取余的结果mod为负，跟取余结果为mod+K是等价的
			mod += K
		}
		result += record[mod]
		record[mod]++
	}
	return result
}

/*
小优化
1. 关于mod为负的处理，代码可以精简
2. 因为mod的取值范围确定，是[0, K-1]，所以record也可以用数组，比哈希表更快
*/
func subarraysDivByK1(A []int, K int) int {
	record := make([]int, K)
	record[0] = 1
	prefixSum, result := 0, 0
	for _, v := range A {
		prefixSum += v
		mod := (prefixSum%K + K) % K
		result += record[mod]
		record[mod]++
	}
	return result
}
