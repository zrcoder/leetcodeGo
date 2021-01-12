/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package array_pair_sum

import "sort"

/*
给定长度为 2n 的数组, 你的任务是将这些数分成 n 对,
例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从1 到 n 的 min(ai, bi) 总和最大。

示例 1:

输入: [1,4,3,2]

输出: 4
解释: n 等于 2, 最大总和为 4 = min(1, 2) + min(3, 4).
*/

/*
分析：
先将数组升序排列后，就很好理解了：
最小的元素，必被计算在最后结果中；为使最终结果最大，和最小元素搭对的必为次小元素，则第一对即为最小元素和次小元素
递归地看，剩余元素也是一样
最终结果就是：计算升序排列后的数组索引为2n（0,2,4,..., 2n-2）的元素和
*/

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}
