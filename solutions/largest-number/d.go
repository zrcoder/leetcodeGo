/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package largest_number

import (
	"sort"
	"strconv"
	"strings"
)

/*
179. 最大数 https://leetcode-cn.com/problems/largest-number/
给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。

示例 1:

输入: [10,2]
输出: 210
示例 2:

输入: [3,30,34,5,9]
输出: 9534330
说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。
*/
/*
思路很简单，细节是魔鬼~

尝试1：
将所有int转化成string，然后降序排序，再一一将这些string链接起来
问题：比如输入3,30那么预期应该是330而不是303

尝试2：
多考虑一下，其实排序的依据是这样：
对于两个string， a和b 如果a+b > b+a， 那么应该把a排在前边~
问题：
输入一堆0， 预期输出0，而不是一堆0

最终，踩完这些坑后，终于成了~
时间复杂度O(nlgn),空间复杂度O(n)
*/
func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, v := range nums {
		strs[i] = strconv.Itoa(v)
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i] // 1：排序比较逻辑
	})
	result := strings.Join(strs, "")
	if result[0] == '0' { // 2: 一堆0的输入有这个结果
		return "0"
	}
	return result
}
