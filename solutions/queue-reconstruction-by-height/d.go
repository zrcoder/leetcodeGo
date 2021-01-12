/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package queue_reconstruction_by_height

import (
	"sort"
)

/*
406. 根据身高重建队列 https://leetcode-cn.com/problems/queue-reconstruction-by-height

假设有打乱顺序的一群人站成一个队列。
每个人由一个整数对(h, k)表示，其中h是这个人的身高，k是排在这个人前面且身高大于或等于h的人数。
编写一个算法来重建这个队列。

注意：
总人数少于1100人。

示例
输入:
[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
输出:
[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
*/
/*
想起中小学时代排队跑操~

题意有点不好理解，这样想一下：
本来所有人站成一队（不一定有序），这时候统计下每个人前边有几个身高大于等于自己的人
突然，打乱了这些人的顺序~~~
问题是恢复这些人的顺序

很自然的思路：
先按照k升序排序（或者按照身高降序排序），再微调， 原地排序
时间复杂度O(n^2),空间复杂度O(1)
*/
func reconstructQueue(people [][]int) [][]int {
	// 先根据k从小到大排序
	sort.Slice(people, func(i, j int) bool {
		return people[i][1] < people[j][1]
	})
	// 由h、k微调顺序
	for i := 1; i < len(people); i++ { // 如果一开始是按照身高降序排序的，这里微调需要从后往前调整
		p := people[i]
		countK := 0 // 统计前边比p高的人数
		j := 0
		// 如果countK 大于 k，需要把这个娃往前移动，j记录需要移动到的位置
		// 如果countK 等于 k，则无需移动;因一开始排序的原因，不会出现countK 小于 k的情况
		for ; j < i; j++ {
			if people[j][0] >= p[0] {
				countK++
				if countK > p[1] {
					break
				}
			}
		}
		if countK > p[1] {
			_ = copy(people[j+1:i+1], people[j:i])
			people[j] = []int{p[0], p[1]}
		}
	}
	return people
}

/*
如果新开辟一个数组，不用在原地排序，代码会简单一点
预处理时不但要按身高降序排列，身高相同的时候还要按照k升序排列
然后从头开始将人们一一放入新开辟的数组，放的时候处理逻辑变得简单
*/
func reconstructQueue1(people [][]int) [][]int {
	// 高的排前边，一样高的按照k升序排列
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	result := make([][]int, 0, len(people))
	for _, p := range people {
		k := p[1]
		// 在索引k处插入 p, 实际不会出现 k > len(result) 的情况
		result = insert(result, k, p)
	}
	return result
}

// 在索引 i 处插入v, 原来 i 及后边的元素后移
func insert(s [][]int, i int, v []int) [][]int {
	return append(append(s[:i:i], v), s[i:]...)
}
