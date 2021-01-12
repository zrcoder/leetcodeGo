/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package maximum_product_of_word_lengths

/*

https://leetcode-cn.com/problems/maximum-product-of-word-lengths

给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，并且这两个单词不含有公共字母。
你可以认为每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。
*/

/*
暴力解决
假设单词最长长度为l,最短长度为s
时间复杂度O(n^2*l)，空间复杂度O(n^2*s)；haveCommonChar优化后变成空间复杂度变成了O(n^2)
*/
func maxProduct1(words []string) int {
	max := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if !haveCommonChar(words[i], words[j]) && len(words[i])*len(words[j]) > max {
				max = len(words[i]) * len(words[j])
			}
		}
	}
	return max
}

func haveCommonChar(s, t string) bool {
	if len(s) > len(t) {
		return haveCommonChar(t, s)
	}
	set := make(map[rune]struct{}, len(s))
	for _, c := range s {
		set[c] = struct{}{}
	}
	for _, c := range t {
		if _, ok := set[c]; ok {
			return true
		}
	}
	return false
}

/*
因word只包含小写字母，haveCommonChar里使用的set是可以优化的
一个转化是将map修改为数组实现：
*/
func haveCommonChar1(s, t string) bool {
	const n = 26
	set := make([]bool, n)
	for _, c := range s {
		set[c-'a'] = true
	}
	for _, c := range t {
		if set[c-'a'] {
			return true
		}
	}
	return false
}

/*
继续优化是使用bitset代替数组，更进一步，因为容量26，一个int32变量即可
*/
func haveCommonChar2(s, t string) bool {
	const n = 26
	set := 0
	for _, c := range s {
		mask := 1 << uint(c-'a')
		set |= mask
	}
	for _, c := range t {
		mask := 1 << uint(c-'a')
		if set&mask != 0 {
			return true
		}
	}
	return false
}

/*
另一个思路：上面对循环调用导致每个单词重复计算了n次掩码
可以在一开始就计算每个单词掩码并存储起来, 后边循环比较是否有公共字符时直接使用，两个掩码做与运算即可，为0说明没有公共字母
时间复杂度O(n^2 + L), 空间复杂度O(n); L为所以单词对总长度
*/
func maxProduct(words []string) int {
	masks := make([]int, len(words))
	for i, v := range words {
		masks[i] = calMask(v)
	}
	result := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if masks[i]&masks[j] == 0 && len(words[i])*len(words[j]) > result {
				result = len(words[i]) * len(words[j])
			}
		}
	}
	return result
}

/*
另一个可能对优化如下：
https://leetcode-cn.com/problems/maximum-product-of-word-lengths/solution/zui-da-dan-ci-chang-du-cheng-ji-by-leetcode

在计算每个单词的掩码时，同时存储对应的长度，如果多个单词有相同的掩码，则存储最大的长度，如 aabb 和 ab， 存储长度为 4

时间复杂度：O(N^2+L)，其中N 是单词数量，L 是所有单词的所有字母的数量。当N>2^26时，时间复杂度为O(L)。
空间复杂度：O(N)，使用一个长度为N 的 HashMap。

这样做比较次数可能会减小，但是 map 处理效率较低，实际测试没有得到优化，反而有些降低性能
*/
func maxProduct2(words []string) int {
	maskLenInfo := make(map[int]int, len(words))
	for _, v := range words {
		mask := calMask(v)
		if len(v) > maskLenInfo[mask] {
			maskLenInfo[mask] = len(v)
		}
	}
	result := 0
	for mask1, length1 := range maskLenInfo {
		for mask2, length2 := range maskLenInfo {
			if mask1&mask2 == 0 && length1*length2 > result {
				result = length1 * length2
			}
		}
	}
	return result
}

func calMask(word string) int {
	mask := 0
	for _, c := range word {
		mask |= 1 << uint(c-'a')
	}
	return mask
}
