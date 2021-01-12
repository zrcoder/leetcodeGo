package ii

/*
454. 四数相加 II

给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。

为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。
所有整数的范围在 -228 到 228 - 1 之间，最终结果不会超过 231 - 1 。

例如:

输入:
A = [ 1, 2]
B = [-2,-1]
C = [-1, 2]
D = [ 0, 2]

输出:
2

解释:
两个元组如下:
1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
*/
/*
分治思想，巧用哈希表
朴素实现会有四层循环，时间复杂度`O(n^4)`,不理想
可以用额外的两个容器分别存A、B中两两元素的和，以及C、D中两两元素的和；然后在这两个容器里统计和为0的组合数
如果这两个容器用切片，则大小都会是`n*n`， 最后遍历实质还是`O(n^4)`
关键来了，这两个容器用哈希表。元素两两之和可能有重复，哈希表可以记录对于一个特定的和，出现了多少次
最后，只需要遍历其中一个哈希表，看每个元素的相反数是否在另一个哈希表里，如果在，那么结果就加上当前和出现的个数乘以相反数在另一个哈希表里出现的次数
这样时间复杂度一下子降低为`O(n^2)`了
*/
func fourSumCount1(A []int, B []int, C []int, D []int) int {
	x, y := combile(A, B), combile(C, D)
	result := 0
	for vX, cX := range x {
		result += cX * y[-vX]
	}
	return result
}

func combile(a, b []int) map[int]int {
	r := make(map[int]int, 0)
	for _, vA := range a {
		for _, vB := range b {
			r[vA+vB]++
		}
	}
	return r
}

/*
还可以进一步优化
两个哈希表，可以只要一个
第一个哈希表统计了A、B两两元素和之后
直接在遍历C、D的过程中确定C、D中元素和的相反数是不是在第一个哈希表中即可，假设在，且出现了x次，则结果加上x就行
这样虽然综合时间、空间复杂度不变，但能节省一个哈希表，同时少一次循环。
*/
func fourSumCount(A []int, B []int, C []int, D []int) int {
	count := make(map[int]int, 0)
	for _, v1 := range A {
		for _, v2 := range B {
			count[v1+v2]++
		}
	}
	result := 0
	for _, v1 := range C {
		for _, v2 := range D {
			result += count[-v1-v2]
		}
	}
	return result
}
