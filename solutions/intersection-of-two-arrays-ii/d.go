/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package intersection_of_two_arrays_ii

import (
	"math"
	"sort"
)

/*
给定两个数组，编写一个函数来计算它们的交集。

示例 1:
输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2,2]
示例 2:
输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [4,9]
说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
我们可以不考虑输出结果的顺序。
进阶:
如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-arrays-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
1. 借助字典
用一个字典，记录一个数组里元素出现的个数；键为元素，值为其出现的个数
然后遍历另一个数组，如果元素在字典里存在则写入结果，并将字典中对应个数减1
为使字典较小，让其记录长度小当数组即可；另外可将已有数组作为结果，来优化空间复杂度
假设数组的长度分别为m，n，则时间复杂度O(m+n)；空间复杂度均为O(min(m, n))
*/
func intersect1(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	if len(nums1) > len(nums2) { // 为了使后背字典尽量小
		nums1, nums2 = nums2, nums1
	}
	m := make(map[int]int, len(nums1))
	for _, v := range nums1 {
		m[v]++
	}
	k := 0
	for _, v := range nums2 {
		if m[v] == 0 {
			continue
		}
		nums1[k] = v // 将nums1作为结果数组；换成nums2也行
		k++
		m[v]--
	}
	return nums1[:k]
}

/*
2.空间复杂度O(1)
将两个数组排序；遍历其中一个，不断在另一个数组里边用两次二分法查找元素出现的次数，同时统计在当前数组里出现的次数，以得到结果
另一个优化是可以利用当前数组作为结果数组

综合看，空间复杂度为常数级O(1)
时间复杂度O(nlogn+mlogm)，主要为排序和最后的筛选结果
*/
func intersect2(nums1 []int, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	if m == 0 || n == 0 {
		return nil
	}
	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	time1 := 1 // 统计nums1里某个元素的个数，下边遍历使用
	k := 0
	// 借用nums1的空间作为返回数组
	for i := 0; i < m-1; i++ {
		v := nums1[i]
		if v == nums1[i+1] {
			time1++
			continue
		}
		time2 := count(nums2, v)  // nums2中v出现的个数
		time := min(time1, time2) // v在结果里的个数， 包含time2为0的情况
		k = write(nums1, k, v, time)
		time1 = 1
	}
	// 处理最后一个元素
	last := nums1[m-1]
	time2 := count(nums2, last)
	if time2 == 0 {
		return nums1[:k]
	}
	if m == 1 || last != nums1[m-2] {
		nums1[k] = last
		k++
	} else {
		time1++
		time := min(time1, time2)
		k = write(nums1, k, last, time)
	}
	return nums1[:k]
}

/*
3. 两个数组排序后，可以一次遍历搞定结果
分别用i，j两个指针遍历两个数组
当发现i处元素比j处小，i++；
i处元素比j处大，j++
相等的时候，则写入结果
注意可以用已有数组做结果

时间复杂度O(nlogn+mlogm)，空间复杂度O(1)

复杂度同2，但是代码量大大减少
*/
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	m, n := len(nums1), len(nums2)
	var i, j, k int
	for i < m && j < n {
		switch {
		case nums1[i] < nums2[j]:
			i++
		case nums1[i] > nums2[j]:
			j++
		default:
			nums1[k] = nums1[i]
			k++
			i++
			j++
		}
	}
	return nums1[:k]
}

// arr 已经排序
func count(arr []int, x int) int {
	left := sort.SearchInts(arr, x)
	if left == len(arr) || arr[left] != x {
		return 0
	}
	return searchFromRight(arr, x) - left
}

// arr 已经排序
func searchFromRight(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] <= target:
			left = mid + 1
		case nums[mid] > target:
			right = mid
		}
	}
	return left
}

//从索引k开始向结果数组里写入time个v, 返回写完后的索引
func write(r []int, k, v, time int) int {
	for j := 0; j < time; j++ {
		r[k] = v
		k++
	}
	return k
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
