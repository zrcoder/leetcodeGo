/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package sort_colors

/*
75. 颜色分类
https://leetcode-cn.com/problems/sort-colors

给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，
使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

注意:
不能使用代码库中的排序函数来解决这道题。

示例:
输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]

进阶：
一个直观的解决方案是使用计数排序的两趟扫描算法。
首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
你能想出一个仅使用常数空间的一趟扫描算法吗？
*/
/*
计数排序
时间复杂度O(n),需要遍历数组两次
空间复杂度O(1)
*/
func sortColors1(nums []int) {
	red, white, blue := 0, 0, 0
	for _, v := range nums {
		if v == 0 {
			red++
		} else if v == 1 {
			white++
		} else if v == 2 {
			blue++
		}
	}
	i := 0
	for ; i < red; i++ {
		nums[i] = 0
	}
	for ; i < red+white; i++ {
		nums[i] = 1
	}
	for ; i < red+white+blue; i++ {
		nums[i] = 2
	}
}

/*
进阶：一次扫描的排序算法，类似快排的思想

分析：
最终结果是{0,0,...0,1,1,...,1,2,2,...2}
可用三个指针i0, i2 和 i，i0 维持指向左侧一堆 0 的右边界，i2 维持右侧一堆 2 的左侧边界；i则指向遍历过程中的当前元素
遍历时，根据 nums[i] 的值决定应该将其加入左侧还是右侧，让 i0 和 i2 稍微跨界一步，就可以通过交换边界处和i处元素做到这一点
需要注意的是，如果把 i 和 i2 交换，这时候的 nums[i] 可能是 0,1,2 三种情况，为了减少判断，这种情况 i 不向后移动

时间复杂度O(n),只遍历数组一次
空间复杂度O(1)
*/
func sortColors(nums []int) {
	i0, i2 := 0, len(nums)-1
	for i := i0; i <= i2; {
		switch nums[i] {
		case 0:
			nums[i0], nums[i] = nums[i], nums[i0]
			i0++
			i++
		case 1:
			i++
		case 2:
			nums[i2], nums[i] = nums[i], nums[i2]
			i2--
		}
	}
}
