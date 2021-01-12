/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package random_pick_with_weight

import (
	"math/rand"
	"sort"
)

/*
528. 按权重随机选择 https://leetcode-cn.com/problems/random-pick-with-weight
给定一个正整数数组 w ，其中 w[i] 代表位置 i 的权重，
请写一个函数 pickIndex ，它可以随机地获取位置 i，选取位置 i 的概率与 w[i] 成正比。

说明:
1 <= w.length <= 10000
1 <= w[i] <= 10^5
pickIndex 将被调用不超过 10000 次

示例1:
输入:
["Solution","pickIndex"]
[[[1]],[]]
输出: [null,0]

示例2:
输入:
["Solution","pickIndex","pickIndex","pickIndex","pickIndex","pickIndex"]
[[[1,3]],[],[],[],[],[]]
输出: [null,0,1,1,1,0]

输入语法说明：
输入是两个列表：调用成员函数名和调用的参数。Solution 的构造函数有一个参数，即数组 w。
pickIndex 没有参数。输入参数是一个列表，即使参数为空，也会输入一个 [] 空列表。
*/

/*
将w[i]看作一个长度为w[i]的区间
则这些长度区间连起来后随机在里边取一个数，问题转化为求这个数所在的区间

用一个数组total，total[i] = w[0] + ... + w[i]
则total[n-1]即w里所有元素和；用库函数获取一个在0到total[n-1]的随机值target
找到这个随机值对应total数组的索引即可：total[i-1] <= target < total[i], i即所求
考虑到效率，用二分法找i
*/
type Solution struct {
	total []int
}

func Constructor(w []int) Solution {
	sum := 0
	total := make([]int, len(w))
	for i, v := range w {
		sum += v
		total[i] = sum
	}
	return Solution{total: total}
}

// 朴素实现
func (s *Solution) PickIndex1() int {
	target := rand.Intn(s.total[len(s.total)-1])
	for i, v := range s.total {
		if v > target {
			return i
		}
	}
	return 0
}

// 改用二分法
func (s *Solution) PickIndex2() int {
	target := rand.Intn(s.total[len(s.total)-1])
	left, right := 0, len(s.total)-1
	for left != right {
		mid := (left + right) / 2
		switch {
		case s.total[mid] <= target:
			left = mid + 1
		case s.total[mid] > target:
			right = mid
		}
	}
	return left
}

// 二分法使用标准库，减少代码量
func (s *Solution) PickIndex() int {
	target := rand.Intn(s.total[len(s.total)-1])
	return sort.Search(len(s.total), func(i int) bool {
		return s.total[i] > target
	})
}
