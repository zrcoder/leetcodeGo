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
题意有点难理解。可以加点背景说明：一开始人们随便站成了一队，然后班长统计了每个人的身高 h 以及排在其前边不比自己矮的人的个数 k。突然这些人一哄而散跑去看美女了。问题是恢复原来的队列。

这个问题总体思路和上边的信封套娃及叠罗汉问题类似。都是排序，经历粗排和细排两轮。

很自然的思路：越高的人k 值理应越小。先按照身高降序，在身高相等的时候怎么排呢？k 小的排前边。

在构建结果数组的时候，如果当前人的 k 不小于结果数组的长度，直接把他追加到对尾，否则，用二分法找到他该插入的位置，当然后边的人要一一后移。
这里需要注意，其实用二分法反而有点浪费，先二分再一一后移一些人，不如放弃二分，一开始直接从结果数组后边向前找，类似冒泡的方法，将当前人插入队里，而他后边的人在冒泡的过程中已经移动好了。
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

func reconstructQueue2(people [][]int) [][]int {
	// 高的排前边，一样高的按照k升序排列
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	result := make([][]int, len(people))
	length := 0
	for _, p := range people {
		k := p[1]
		i := length
		for i > k { // 根据前边的排序，实际不会出现 k > length 的情况
			result[i] = result[i-1]
			i--
		}
		result[i] = p
		length++
	}
	return result
}
