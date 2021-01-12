/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package increasing_subsequences

import (
	"math"
)

/*
491. 递增子序列
https://leetcode-cn.com/problems/increasing-subsequences

给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。

示例:
输入: [4, 6, 7, 7]
输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]

说明:
给定数组的长度不会超过15。
数组中的整数范围是 [-100,100]。
给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。
*/

// 78. 子集 问题的变体

/*
递归与回溯

递归枚举子序列的通用模板：

用 tmp 保存当前选出的子序列， cur 表示当前索引

```
var ans [][]int
var tmp []int
func dfs(cur int) {
    if (cur == len(nums) {
        // 判断是否合法，如果合法判断是否重复，将满足条件的加入答案
        if (isValid() && notVisited()) {
			ans = append(ans, tmp)
        }
        return
    }
    // 选择当前元素
	tmp = append(tmp, nums[cur])
    dfs(cur + 1)
	// 回溯选择
	tmp = tmp[:len(tmp)-1]

    // 不选择当前元素
    dfs(cur + 1);
}
```
其中“选择当前元素”和“不选择当前元素”的顺序可以调换

在 dfs(cur) 开始之前， [0, cur-1] 区间的所有元素已经被考虑过， [cur, n-1] 区间的元素还未被考虑。
执行  dfs(cur) 时， 考虑 cur 处的元素选或不选：如果选，则常规递归回溯；如果不选，直接递归下一位置。


当然，在这个问题中，可以做一些简单优化，使枚举出来的结果都是合法的且不重复：
1. 使序列合法，只需在选择的时候保证当前元素不小于上一个选择的元素即可
2. 怎么保证没有重复？ 需要给“不选择”做一个限定，只有当当前元素不等于上一个选择的元素的时候，才考虑不选择当前元素而递归后边的元素

时间复杂度 O((2^n)*n)
空间复杂度 O(n)
*/
func findSubsequences(nums []int) [][]int {
	var ans [][]int
	var tmp []int
	var dfs func(int, int)
	dfs = func(cur, last int) {
		if cur == len(nums) {
			if len(tmp) > 1 {
				t := make([]int, len(tmp))
				_ = copy(t, tmp)
				ans = append(ans, t)
			}
			return
		}
		// 不选择当前元素
		if nums[cur] != last {
			dfs(cur+1, last)
		}
		// 选择当前元素
		if nums[cur] >= last {
			tmp = append(tmp, nums[cur])
			dfs(cur+1, nums[cur])
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(0, math.MinInt32)
	return ans
}

/*
另一个解

查找递增数组实际就是以每个元素为起点，后面每个比它大的元素都可以发散出一条路径，所以可以用dfs。
数组中重复的元素只需要以前面的元素为起点进行dfs，因为后面元素的所有情况在前面都可以考虑到。
所以对需要dfs的nums元素用set去重；
注意在以i为起点dfs时，还要对后面的元素相互去重
*/
func findSubsequences1(nums []int) [][]int {
	var cur []int
	var res [][]int

	var dfs func(int)
	dfs = func(start int) {
		cur = append(cur, nums[start])
		if len(cur) > 1 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
		}
		visited := make(map[int]struct{}, 0)
		for i := start + 1; i < len(nums); i++ {
			if _, ok := visited[nums[i]]; !ok && nums[start] <= nums[i] {
				visited[nums[i]] = struct{}{}
				dfs(i)
			}
		}
		cur = cur[:len(cur)-1] // 回溯
	}

	visited := make(map[int]struct{}, 0)
	for i := 0; i < len(nums)-1; i++ { // 递增子序列长度至少为2， 没有必要以最后一个元素为起点去dfs遍历了
		// 数组中可能有重复元素，后边的重复元素dfs遍历情况包含在前边的重复元素dfs遍历里边，不用再遍历
		if _, ok := visited[nums[i]]; ok {
			continue
		}
		visited[nums[i]] = struct{}{}
		dfs(i) // 以i为起点深度优先遍历
	}
	return res
}

/*
 二进制枚举 + 哈希

时间复杂度 O((2^n)*n)
空间复杂度 O(2^n)， 最坏情况是整个序列递增，seen 要存放所有子序列
*/
func findSubsequences0(nums []int) [][]int {
	var res [][]int
	end := 1 << (len(nums))
	seen := map[int]bool{}
	for mask := 1; mask < end; mask++ {
		sub := genIncreaseSub(mask, nums)
		if len(sub) < 2 {
			continue
		}
		hashCode := calHash(sub)
		if !seen[hashCode] {
			res = append(res, sub)
			seen[hashCode] = true
		}
	}
	return res
}

func genIncreaseSub(mask int, nums []int) []int {
	var res []int
	for _, v := range nums {
		if 1&mask == 1 && (len(res) == 0 || res[len(res)-1] <= v) {
			res = append(res, v)
		}
		mask >>= 1
	}
	return res
}

/*
Rabin-Karp 编码

参考 1392. 最长快乐前缀

可以把这个编码方法应用到整数数组
*/

func calHash(s []int) int {
	res := 0
	base := 203         // max(v+101)+1 = 202, 选中一个大于它的质数
	mod := int(1e9 + 7) // 大质数， 小于 math.MaxInt32
	for _, v := range s {
		res = (res*base + (v + 101)) % mod // v + 101 保证为正数
	}
	return res
}
