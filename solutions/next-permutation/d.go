package permutation

/*
31. 下一个排列
https://leetcode-cn.com/problems/next-permutation

实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/
/*
较难分析清楚，代码容易写

https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-suan-fa-xiang-jie-si-lu-tui-dao-/

时间复杂度 O(n)
*/
func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	// 从右侧开始寻找第一个比右邻居小的元素
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	// 此时 nums[i+1:]必为非递增序列
	if i >= 0 {
		// 从后向前找到第一个大于i处元素的元素
		k := len(nums) - 1
		for nums[i] >= nums[k] {
			k--
		}
		// 交换
		nums[i], nums[k] = nums[k], nums[i]

	}
	// 此时 nums[i+1:]必为非递增序列，逆序变成非递减
	for l, r := i+1, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
