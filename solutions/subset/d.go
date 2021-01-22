/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package subset

/*
78. 子集
https://leetcode-cn.com/problems/subsets

给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

/*
朴素实现
时空复杂度均为O(n*2^n)
*/
func subsets(nums []int) [][]int {
	res := [][]int{{}} // 空集也是子集之一
	for _, num := range nums {
		for _, r := range res {
			tmp := make([]int, len(r)+1)
			copy(tmp, r)
			tmp[len(tmp)-1] = num
			res = append(res, tmp)
		}
	}
	return res
}

/*
回溯1
参考 491. 递增子序列 解法二，使用递归枚举子序列的通用模板
*/
func subsets1(nums []int) [][]int {
	var res [][]int
	var cur []int
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(nums) {
			res = append(res, copySlice(cur))
			return
		}
		// 不选择当前元素
		backtrack(i + 1)
		// 选择当前元素
		cur = append(cur, nums[i])
		backtrack(i + 1)
		cur = cur[:len(cur)-1]
	}
	backtrack(0)
	return res
}

// 回溯2
func subsets2(nums []int) [][]int {
	var res [][]int
	var cur []int
	var backtrack func(start int)
	backtrack = func(start int) {
		res = append(res, copySlice(cur))
		for i := start; i < len(nums); i++ {
			cur = append(cur, nums[i])
			backtrack(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(0)
	return res
}

func copySlice(s []int) []int {
	r := make([]int, len(s))
	copy(r, s)
	return r
}

/*
二进制枚举

nums 里的每个元素，要么在结果中，要么不在结果中
用一个 n 位的 bitset 来表示各个元素在不在结果中，
如 000...000 表示所有元素都不在结果中，000..011 表示后边两个元素在结果中

局限：len(nums)不能大于64， 否则无法用一个int做mask

时空复杂度均为O(n*2^n)
*/
func subsets3(nums []int) [][]int {
	var res [][]int
	max := 1 << len(nums)
	for state := 0; state < max; state++ {
		var cur []int
		for i, v := range nums {
			if (1<<i)&state != 0 {
				cur = append(cur, v)
			}
		}
		res = append(res, cur)
	}
	return res
}

/*
## [1593. 拆分字符串使唯一子字符串的数目最大](https://leetcode-cn.com/problems/split-a-string-into-the-max-number-of-unique-substrings/)

难度中等

给你一个字符串 `s` ，请你拆分该字符串，并返回拆分后唯一子字符串的最大数目。

字符串 `s` 拆分后可以得到若干 **非空子字符串** ，这些子字符串连接后应当能够还原为原字符串。但是拆分出来的每个子字符串都必须是 **唯一的** 。

注意：**子字符串** 是字符串中的一个连续字符序列。



**示例 1：**

```
输入：s = "ababccc"
输出：5
解释：一种最大拆分方法为 ['a', 'b', 'ab', 'c', 'cc'] 。像 ['a', 'b', 'a', 'b', 'c', 'cc'] 这样拆分不满足题目要求，因为其中的 'a' 和 'b' 都出现了不止一次。
```

**示例 2：**

```
输入：s = "aba"
输出：2
解释：一种最大拆分方法为 ['a', 'ba'] 。
```

**示例 3：**

```
输入：s = "aa"
输出：1
解释：无法进一步拆分字符串。
```



**提示：**

- `1 <= s.length <= 16`
- `s` 仅包含小写英文字母

函数签名：

```go
func maxUniqueSplit(s string) int
```

### 分析

#### 回溯

对子集问题加了限制：所有子集不能相同。可以在回溯的过程中用哈希表去重。

> 另需注意，s 本身就是一个子集，但是这里空集不算。如示例3。

*/

func maxUniqueSplit(s string) int {
	type Set = map[string]bool
	var dfs func(i int)
	set := make(Set, 0)
	res := 1
	dfs = func(start int) {
		if start == len(s) {
			res = max(res, len(set))
			return
		}
		for i := start; i < len(s); i++ {
			sub := s[start : i+1]
			if !set[sub] {
				set[sub] = true
				dfs(i + 1)
				delete(set, sub)
			}
		}
	}
	dfs(0)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
