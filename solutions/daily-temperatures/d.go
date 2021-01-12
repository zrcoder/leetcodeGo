package daily_temperatures

/*
739. 每日温度 https://leetcode-cn.com/problems/daily-temperatures/
请根据每日 气温 列表，重新生成一个列表。
对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。
如果气温在这之后都不会升高，请在该位置用 0 来代替。

例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，
你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
*/

// 朴素实现，时间复杂度 O(n^2)
func dailyTemperatures0(T []int) []int {
	result := make([]int, len(T))
	for i, v := range T {
		j := i + 1
		for j < len(T) && T[j] <= v {
			j++
		}
		if j < len(T) {
			result[i] = j - i
		}
	}
	return result
}

/*
单调栈, 思想和朴素实现几乎一样。不过秒的是大大降低了时间复杂度。
从后向前遍历T，维持一个单调递减栈来装入T中元素的索引
对于位置i，其对应的元素v如果不大于栈顶元素，则栈顶元素出栈，一直出，直到栈空或者栈顶元素大于v
大于i处元素的元素索引就是栈顶的值
每个元素最多进栈出栈一次，总体时间复杂度为 O(n)
当然，引入了额外的栈，空间复杂度上升为 O(n)
*/
func dailyTemperatures(T []int) []int {
	result := make([]int, len(T))
	var stack []int
	for i := len(T) - 1; i >= 0; i-- {
		for len(stack) > 0 && T[stack[len(stack)-1]] <= T[i] {
			stack = stack[:len(stack)-1]
		}
		// if len(stack) == 0, result[i] = 0
		if len(stack) > 0 {
			result[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return result
}
