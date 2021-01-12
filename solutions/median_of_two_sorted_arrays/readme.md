# [4. 寻找两个有序数组的中位数](https://leetcode-cn.com/problems/median-of-two-sorted-arrays)

## 题目
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:
```text
nums1 = [1, 3]
nums2 = [2]
则中位数是 2.0
```
示例 2:
```text
nums1 = [1, 2]
nums2 = [3, 4]
则中位数是 (2 + 3)/2 = 2.5
```
## 分析
对于一个有序数组，如果元素个数为奇数，中位数即中间元素的值；若元素个数为偶数，中位数为中间两个元素的平均值。<br>
对于两个或多个有序数组，其合并后的中位数并非每个数组中位数的平均值，如：
```
[1, 3, 5] // 中位数3
[8, 10] // 中位数9
// 合并后的数组
[1, 3, 5, 8, 10] // 中位数5, 并非3和9的平均数
```
所以，必须对两个数组合并，合并后依然有序<br>

#### 0. 朴素实现（时间与空间复杂度均为O(m+n)）

```
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	return medianOf(merge(nums1, nums2))
}

func merge(nums1, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	r := make([]int, m+n)
	for i, j, k := 0, 0, 0; i < m || j < n; k++ {
		if j == n {
			r[k] = nums1[i]
			i++
		} else if i == m {
			r[k] = nums2[j]
			j++
		} else if nums1[i] < nums2[j] {
			r[k] = nums1[i]
			i++
		} else {
			r[k] = nums2[j]
			j++
		}
	}
	return r
}

func medianOf(nums []int) float64 {
	length := len(nums)
	if length == 0 {
		return 0.0
	}
	if length%2 == 0 {
		return float64(nums[length/2]+nums[length/2-1]) * 0.5
	}
	return float64(nums[length/2])
}
```
#### 1. 朴素实现的改进，不用真的merge
```
func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	lastR, currentR := -1, -1
	start1, start2 := 0, 0
	for i := 0; i <= (m+n)/2; i++ {
		lastR = currentR
		if start1 < m && (start2 >= n || nums1[start1] < nums2[start2]) {
			currentR = nums1[start1]
			start1++
		} else {
			currentR = nums2[start2]
			start2++
		}
	}
	if (m+n)%2 == 1 {
		return float64(currentR)
	}
	return float64(lastR+currentR) * 0.5
}
```
#### 2. 时间O(log(m+n))，空间O(1)

原理参考：<br>
https://cloud.tencent.com/developer/article/1483811 <br>
问题转化为求数组第k个元素： 对于两个数组，假设长度分别是m、n，求合并后的中位数即求：<br>
```text
i. 合并后第(m+n)/2 + 1 个元素（m+n为奇数）
ii. 合并后第(m+n)/2 个元素与第(m+n)/2 + 1个元素的平均值（m+n为偶数）
```
对于两个数组，求合并后的第k个元素，可以分别取两个数组第k/2个元素，通过比较这两个元素的大小，可以批量地减少搜索范围
```text
1.如果nums1[k/2] < nums2[k/2], 说明合并后的第k个元素肯定不在nums[0:k/2+1]区间里
可以继续在nums1[k/2:]和nums2中搜索第k-(k/2+1)个元素
2.反之，可以排除nums2的前k/2个元素继续搜索
需要注意边界
```
```
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size := len(nums1) + len(nums2)
	if size == 0 {
		return 0.0
	}
	if size%2 == 1 {
		return getKth(nums1, nums2, size/2+1)
	}
	return (getKth(nums1, nums2, size/2) + getKth(nums1, nums2, size/2+1)) * 0.5
}
func getKth(nums1, nums2 []int, k int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return getKth(nums2, nums1, k)
	}
	if m == 0 {
		return float64(nums2[k-1])
	}
	if k == 1 {
		return float64(min(nums1[0], nums2[0]))
	}
	i, j := min(m-1, k/2-1), min(n-1, k/2-1)
	if nums1[i] > nums2[j] {
		return getKth(nums1, nums2[j+1:], k-(j+1))
	}
	return getKth(nums1[i+1:], nums2, k-(i+1))
}
```

#### 3. 时间O(log(min(m,n)))，空间O(1)
原理参考：<br>
https://blog.csdn.net/bjweimengshu/article/details/97717144 <br>
用i，j两个指针将两个数组分别划分为两部分，将nums1的左半部分和nums2的左半部分合起来看作合并后的左半部分，同样可以得到合并后右半部分<br>
```text
                        |
nums1       0, ..., i-1,| i, ..., m-1
                        |
nums2 0, 1, ...,    j-1,| j, ..., n-1
                        |
              左半部分   |  右半部分
```
如果能保证左右部分的大小相当（m+n为偶数则左右部分大小相等；为奇数则左半部分比右半部分多一个），也就找到了合并后的中位数
```text
m+n为偶数时：
i+j = m-i + n-j 即i+j = （m+n）/2
m+n为奇数时：
i+j = m-i + n-j + 1也就是 i+j = (m+n+1)/2
因整数除法特性，可以统一为i+j = (m+n+1)/2
```
注意到确定了i，就确定了j， j = (m+n+1)/2 - i；<br>
数组已排序，用二分搜索法来确定i:
```text
因为两个数组都是有序的，所以 nums1[i-1] <= nums1[i]，nums2[i-1] <= nums2[i] 是天然具备的，
所以只需要保证 nums2[j-1] < = nums1[i] 和 nums1[i-1] <= nums2[j];对不满足的情况分两种情况讨论：
nums2[j-1] > nums1[i]
此时需要增加i
nums1[i-1] > nums2[j]
此时要减少i
```
```
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays1(nums2, nums1) // 方便处理后边二分搜索的边界情况
	}

	// 二分搜索，寻找合适的 i 和 j
	left, right := 0, m
	for left <= right {
		i := left + (right-left)/2
		j := (m+n+1)/2 - i
		if i < m && j > 0 && nums2[j-1] > nums1[i] { // i 小了
			left = i + 1
		} else if i > 0 && j < n && nums1[i-1] > nums2[j] { // i 大了
			right = i - 1
		} else { // i 正好
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = min(nums1[i], nums2[j])
			}
			return float64(maxLeft+minRight) / 2
		}
	}
	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
