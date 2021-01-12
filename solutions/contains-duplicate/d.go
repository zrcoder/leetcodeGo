/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package contains_duplicate

import "math"

/*
219. 存在重复元素 II https://leetcode-cn.com/problems/contains-duplicate-ii

给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，
使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。

示例 1:
输入: nums = [1,2,3,1], k = 3
输出: true

示例 2:
输入: nums = [1,0,1,1], k = 1
输出: true

示例 3:
输入: nums = [1,2,3,1,2,3], k = 2
输出: false
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if index, ok := m[v]; ok && i-index <= k {
			return true
		}
		m[v] = i
	}
	return false
}

/*
220. 存在重复元素 III https://leetcode-cn.com/problems/contains-duplicate-iii

给定一个整数数组，判断数组中是否有两个不同的索引 i 和 j，
使得 nums [i] 和 nums [j] 的差的绝对值最大为 t，并且 i 和 j 之间的差的绝对值最大为 ķ。

示例 1:
输入: nums = [1,2,3,1], k = 3, t = 0
输出: true

示例 2:
输入: nums = [1,0,1,1], k = 1, t = 2
输出: true
示例 3:

输入: nums = [1,5,9,1,5,9], k = 2, t = 3
输出: false
*/

/*
线性搜索

时间复杂度O(n * min(n, k))
空间复杂度O(1)
*/
func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	n := len(nums)
	for i := 0; i < n; i++ {
		end := min(n-1, i+k)
		for j := i + 1; j <= end; j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

/*
滑动窗口+桶
维持一个长度为k的滑动窗口；用若干桶来装各个数字；

不断右移这个窗口，先根据当前数字的值计算出它应该放入的桶，如果桶里已有数字，表示找到了差值绝对值在t之内的数字直接返回，否则放入对应的桶

如t = 2, 每个桶内应该有3个数字，每个桶内的数字如下：
桶的编号     ...     -2            -1             0           1      ...
                  -------        -------       -------     -------
桶内数字范围      | -6 ~ -4  |    | -3 ~ -1 |   | 0 ~ 2 |   | 3 ~ 5 |
                  -------        -------       -------     -------

时间复杂度O(n)
空间复杂度O(min(n,k)), 最坏情况下所有元素都放入map
*/
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}
	// 所有的桶，键为桶id，值为数字，没必要记录桶里所有的数字只需要记录一个，后续再向桶里插入的时候如果桶里已有数字可以直接返回true
	buckets := map[int]int{}
	bucketSize := t + 1 // 桶的大小，即每个桶里应该放几个数字
	for i, v := range nums {
		// 删除窗口中第一个数字对应的桶
		if i > k {
			firstBucket := getBucket(nums[i-k-1], bucketSize)
			delete(buckets, firstBucket)
		}
		// v应该放入的桶编号
		bucket := getBucket(v, bucketSize)
		// 检查当前桶里是否已经有值
		if _, ok := buckets[bucket]; ok {
			return true
		}
		// 检查前一个桶里是否有满足题意的值
		n, ok := buckets[bucket-1]
		if ok && abs(v-n) <= t {
			return true
		}
		// 检查后一个桶里是否有满足题意的值
		n, ok = buckets[bucket+1]
		if ok && abs(v-n) <= t {
			return true
		}
		buckets[bucket] = v
	}
	return false
}

func getBucket(num, size int) int {
	if num >= 0 {
		return num / size
	}
	// num加1, 把负数移动到从0开始，这样计算出的标号最小是0，需要再减1
	return (num+1)/size - 1
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
