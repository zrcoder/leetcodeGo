/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package median_of_two_sorted_arrays

// 朴素实现，先merge再找中间的， 代码略

// 朴素实现的改进，不用真的merge
func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	lastR, currentR := -1, -1
	start1, start2 := 0, 0
	for i := 0; i <= (m+n)/2; i++ {
		lastR = currentR
		if start1 < m && (start2 < n && nums1[start1] <= nums2[start2] || start2 == n) {
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

// 2. 时间O(log(m+n))，空间O(1)
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
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

// 3. 时间O(log(min(m,n)))，空间O(1)

var (
	m, n   int
	s1, s2 []int
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n = len(nums1), len(nums2)
	if m*n == 0 {
		return 0.0
	}
	s1, s2 = nums1, nums2
	if m > n {
		// 保证 nums1 长度较小，方便处理后边二分搜索的边界情况
		m, n = n, m
		s1, s2 = nums2, nums1
	}
	return binarySearch()
}

func binarySearch() float64 {
	left, right := 0, m
	for left <= right {
		i := left + (right-left)/2
		j := (m+n+1)/2 - i
		if i < m && j > 0 && s2[j-1] > s1[i] { // i 小了
			left = i + 1
		} else if i > 0 && j < n && s1[i-1] > s2[j] { // i 大了
			right = i - 1
		} else { // i, j 正好
			return cal(i, j)
		}
	}
	return 0.0
}

func cal(i, j int) float64 {
	maxLeft := 0
	if i == 0 {
		maxLeft = s2[j-1]
	} else if j == 0 {
		maxLeft = s1[i-1]
	} else {
		maxLeft = max(s1[i-1], s2[j-1])
	}

	if (m+n)%2 == 1 {
		return float64(maxLeft)
	}

	minRight := 0
	if i == m {
		minRight = s2[j]
	} else if j == n {
		minRight = s1[i]
	} else {
		minRight = min(s1[i], s2[j])
	}
	return float64(maxLeft+minRight) / 2
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
