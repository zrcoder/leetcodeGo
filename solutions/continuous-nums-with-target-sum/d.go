/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package continuous_nums_with_target_sum

import "math"

/*
[面试题57 - II. 和为s的连续正数序列](https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/)

输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

示例 1：
输入：target = 9
输出：[[2,3,4],[4,5]]

示例 2：
输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]

限制：
1 <= target <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
双指针
时间复杂度O(target);两个指针移动均单调不减，且最多移动 target/2 次
空间复杂度O(1)，除了答案数组只需要常数的空间
*/
func findContinuousSequence(target int) [][]int {
	var r [][]int
	for left, right := 1, 2; left < right; {
		size := right - left + 1
		sum := (left + right) * size / 2
		if sum == target {
			t := make([]int, size)
			for i := range t {
				t[i] = left + i
			}
			r = append(r, t)
			left++
			right++
		} else if sum < target {
			right++
		} else {
			left++
		}
	}
	return r
}

/*
类似双指针的解法，额外引入一个队列，空间复杂度增加为O(target)
*/
func findContinuousSequence0(target int) [][]int {
	var r [][]int
	var queue []int
	sum := 0
	for i := 1; i < target; i++ {
		queue = append(queue, i)
		sum += i
		for sum > target && len(queue) > 0 {
			sum -= queue[0]
			queue = queue[1:]
		}
		if sum == target {
			t := make([]int, len(queue))
			_ = copy(t, queue)
			r = append(r, t)
		}
	}
	return r
}

/*
解一元二次方程
如上边的解法，假设用x、y分别表示left，right
为了使sum == target <=> (x+y) * (y-x+1) / 2 == target <=> y^2 + y -x^2 + x - 2*target == 0
可以看成是一个关于y的一元二次方程，其中a=1，b=1, c = -x^2 + x - 2*target
可以用求根公式解得y；需要满足：
判别式delta = b^2 - 4ac 需要大于0且平方根需为整数
*/
func findContinuousSequence1(target int) [][]int {
	var r [][]int
	limit := (target - 1) / 2 // 等效于 target/2 向下取整
	for x := 1; x <= limit; x++ {
		delta := 1 - 4*(-x*x+x-2*target)
		if delta < 0 {
			continue
		}
		deltaSqrt := int(math.Sqrt(float64(delta)))
		if deltaSqrt*deltaSqrt == delta && (deltaSqrt-1)%2 == 0 {
			y := (deltaSqrt - 1) / 2 // 另一个解(-1-delta_sqrt)/2必然小于0，不用考虑
			if x >= y {
				continue
			}
			t := make([]int, y-x+1)
			for i := range t {
				t[i] = x + i
			}
			r = append(r, t)
		}
	}
	return r
}
