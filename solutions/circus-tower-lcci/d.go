/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package circus_tower_lcci

import (
	"sort"
)

type Person struct {
	height, weight int
}

func bestSeqAtIndex1(height []int, weight []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	persons := make([]Person, n)
	for i := range persons {
		persons[i].height = height[i]
		persons[i].weight = weight[i]
	}
	sort.Slice(persons, func(i, j int) bool {
		// 身高高的在前边，身高相等则体重轻的在前边
		if persons[i].height == persons[j].height {
			return persons[i].weight < persons[j].weight
		}
		return persons[i].height > persons[j].height
	})
	k := 0
	for _, p := range persons {
		// 在结果中找到第一个不能叠在p上面的人, 二分法
		j := sort.Search(k, func(i int) bool {
			c := persons[i]
			return c.height <= p.height || c.weight <= p.weight
		})
		// 将这个个人替换成p
		persons[j] = p
		if j == k {
			k++
		}
	}
	return k
}
