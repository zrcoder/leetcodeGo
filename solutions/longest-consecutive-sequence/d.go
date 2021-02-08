package longest_consecutive_sequence

import "sort"

/*
128. 最长连续序列 https://leetcode-cn.com/problems/longest-consecutive-sequence/
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
*/
/*
朴素实现，先排序，再统计
时间复杂度主要是排序，O(n lgn)；空间复杂度O(1)
*/
func longestConsecutive0(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	res, curLen := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			continue
		}
		if nums[i-1]+1 == nums[i] {
			curLen++
		} else {
			curLen = 1
		}
		if curLen > res {
			res = curLen
		}
	}
	return res
}

/*
不排序，将数组中所有元素存入哈希表中，这个哈希表既作为集合来来判断某个数字是否存在于`nums`中，也作为哈希表判断某个数字是否在之前的遍历中已经统计过
具体步骤：遍历哈希表中所有元素
1.对于当前元素num，统计以num开头的连续序列长度：尝试在哈希表中找num+1， num+2... 直到找不到下个数字
2.开始对每个num尝试前，如果num-1存在于集合中，或者num这个数字已经在步骤1里访问过，则直接跳过
这样保证在整个哈希表中，每个元素只被访问一次
时间复杂度是`O(2n)` = `O(n)`, 空间复杂度`O(n)`

实际测试，两种实现耗时都是4ms，主要原因应该是测试用例规模不够大，另外哈希表的查找虽说是O(1)，其实远不如数组的O(1)快
当然内存占用，纯排序不借用额外空间要优一点
*/
func longestConsecutive(nums []int) int {
	memo := make(map[int]bool, len(nums))
	for _, v := range nums {
		memo[v] = false
	}
	res := 0
	for num := range memo {
		if _, exist := memo[num-1]; exist || memo[num] {
			continue
		}
		memo[num] = true
		currLen := 1
		currNum := num + 1
		_, exist := memo[currNum]
		for exist {
			memo[currNum] = true
			currLen++
			currNum++
			_, exist = memo[currNum]
		}
		if currLen > res {
			res = currLen
		}
	}
	return res
}
