/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package partition_to_k_equal_sum_subsets

import (
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	sum, max := 0, 0 // 注意到输入限制为非负，且和不会溢出
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	target := sum / k // 尝试把所有元素放到k个组中，每组元素和为target
	if sum%k != 0 || max > target {
		return false
	}
	used := make([]bool, len(nums))
	return backTracking(k, 0, 0, target, nums, used)
}

/*
backTracking通过穷举回溯，尝试把nums所有元素放到k个组中，每组元素和为target
currSubSum 记录当前组累加的结果，start指明从nums的哪个位置开始

如果nums是非递增序列，将大大减小递归次数；在某些场景，可以考虑对nums先做排序处理，再调用backTracking
*/
func backTracking(k, currSubSum, start, target int, nums []int, used []bool) bool {
	if k == 0 { // 说明所有数字都放入了对应组
		return true
	}
	if currSubSum == target { // 已经构建了一组
		// 构建下一组
		return backTracking(k-1, 0, 0, target, nums, used)
	}
	for i := start; i < len(nums); i++ {
		if !used[i] && currSubSum+nums[i] <= target {
			used[i] = true                                                    // 当前值放入当前构建的组
			if backTracking(k, currSubSum+nums[i], i+1, target, nums, used) { // currSubSum本身不改，回溯时也不必改
				return true
			}
			used[i] = false // 说明将当前值放入当前组不能得到结果，回溯
		}
	}
	return false
}

/*
穷举搜索：
对于 nums 中的每个数字，我们可以将其添加到 k 个子集中的一个，只要该子集的和不会超过目标值。
对于每一个选择，我们都递归地用一个更小的数字进行搜索，以便在nums中考虑。如果我们成功地放置了每个数字，那么我们的搜索就成功了。
时间复杂度：O(k^（N-k） * k!)。其中 N 指的是 nums 的长度。
空间复杂度：O(N)，递归调用 search 所使用的堆栈空间。

作者：LeetCode
链接：https://leetcode-cn.com/problems/partition-to-k-equal-sum-subsets/solution/hua-fen-wei-kge-xiang-deng-de-zi-ji-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func canPartitionKSubsets0(nums []int, k int) bool {
	// 根据题目输入限制，累加不会溢出
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k // 尝试把所有元素放到k个组中，每组元素和为target
	sort.Ints(nums)   // 排序，以便尝试先放置较大的元素，减少递归，尽快得出结论
	curr := len(nums) - 1
	if nums[curr] > target {
		return false
	}
	for k >= 0 && nums[curr] == target {
		curr--
		k--
	}
	groups := make([]int, k)
	return search(groups, curr, nums, target)
}

func search(groups []int, curr int, nums []int, target int) bool {
	if curr < 0 { // 说明所有数字都放入了对应组
		return true
	}
	num := nums[curr]
	for i := 0; i < len(groups); i++ {
		if groups[i]+num <= target {
			groups[i] += num                          // 当前值放入组i
			if search(groups, curr-1, nums, target) { // 尝试放入下一个值到合适的组
				return true
			}
			groups[i] -= num // 说明将当前值num放入组i不能得到结果，回溯
		}
		/*
			确保每个 group 的所有 0 值都出现在数组 groups 的末尾。这大大减少了重复的工作。
			例如，在第一次运行搜索时，我们只进行一次递归调用，而不是 k 次。
			还可以通过跳过 group[i] 的重复值来加快速度，但这是不必要的。
		*/
		if groups[i] == 0 {
			break
		}
	}
	return false
}
