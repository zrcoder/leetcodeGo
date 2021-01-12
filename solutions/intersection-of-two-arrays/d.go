/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package intersection_of_two_arrays

import "sort"

/*
给定两个数组，编写一个函数来计算它们的交集。

示例 1:
输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2]

示例 2:
输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [9,4]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
1。借助一个set解决
复杂度：
假设两个数组但长度分别为m，n
时间复杂度是O(m+n)
空间复杂度是O(2*min(m, n)) = O(2*min(m, n)), set和结果数组需要的长度
*/
func intersection1(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	if len(nums1) > len(nums2) { // 保证nums1长度较小，后边set和result分配但空间将较小
		nums1, nums2 = nums2, nums1
	}
	set := make(map[int]struct{}, len(nums1))
	for _, v := range nums1 {
		set[v] = struct{}{}
	}
	result := make([]int, 0, len(nums1))
	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			result = append(result, v)
			delete(set, v)
		}
	}
	return result
}

/*
2. 空间复杂度尽量低的实现
将两个数组排序，其中一个排序是为了方便二分查找，另一排序是为了在遍历时判断是否有重复元素（当前元素与上一个比较即可）

复杂度：
综合时间复杂度O(n*logn + mlogm)
可以利用已有数组，不用新建数组,空间复杂度： O(1)
*/
func intersection2(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	k := 0
	last := nums2[0]
	for i := 0; i < len(nums2); i++ {
		if isSameAsLast(nums2[i], last, i) || !has(nums1, nums2[i]) {
			continue
		}
		nums2[k] = nums2[i]
		k++
		last = nums2[i]
	}
	return nums2[:k]
}

func has(arr []int, val int) bool {
	i := sort.SearchInts(arr, val)
	return i < len(arr) && arr[i] == val
}
func isSameAsLast(last, current, index int) bool {
	return index > 0 && last == current
}

/*
3. 两个数组排序后，可以一次遍历搞定结果
分别用i，j两个指针遍历两个数组
当发现i处元素比j处小，i++；
i处元素比j处大，j++
相等的时候，则写入结果，且i和j右移到不等于当前值的位置
注意可以用已有数组做结果
*/
func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var i, j, k int
	for i < len(nums1) && j < len(nums2) {
		switch {
		case nums1[i] > nums2[j]:
			j++
		case nums1[i] < nums2[j]:
			i++
		default:
			v := nums1[i]
			nums1[k] = v // 用nums1作为结果数组，也可换用nums2
			k++
			for i < len(nums1) && nums1[i] == v {
				i++
			}
			for j < len(nums2) && nums2[j] == v {
				j++
			}
		}
	}
	return nums1[:k]
}
