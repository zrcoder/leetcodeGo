/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package rearrange_k_distance_apart

import (
	"sort"
	"strings"
)

func reorganizeString0(s string) string {
	return rearrangeString(s, 2)
}

func reorganizeString(s string) string {
	result := []byte(s)
	pairs := count(result)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})
	if 2*pairs[0].count-1 > len(s) {
		return ""
	}
	j := 0
	for i := 0; i < len(pairs) && pairs[i].count > 0; i++ {
		for t := 0; t < pairs[i].count; t++ {
			result[j] = pairs[i].char
			j += 2
			if j >= len(s) {
				j = 1
			}
		}
	}
	return string(result)
}

func reorganizeString1(S string) string {
	count := make(map[rune]int)
	n := len(S)
	var maxChar rune
	var maxCount int
	for _, r := range S {
		count[r]++
		if count[r] > maxCount {
			maxChar = r
			maxCount = count[r]
		}
	}
	if 2*maxCount-1 > n {
		return ""
	}
	builder := &strings.Builder{}
	builder.Grow(n)
	repeat(builder, maxChar, maxCount)
	delete(count, maxChar)
	for r, c := range count {
		repeat(builder, r, c)
	}
	tmp := builder.String()
	result := make([]byte, n)
	limited := (n + 1) / 2
	for i, k := 0, 0; i < limited; i++ {
		for j := i; j < n; j += limited {
			result[k] = tmp[j]
			k++
		}
	}
	return string(result)
}

func repeat(builder *strings.Builder, r rune, n int) {
	for i := 0; i < n; i++ {
		builder.WriteRune(r)
	}
}
