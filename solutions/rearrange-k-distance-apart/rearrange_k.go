/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package rearrange_k_distance_apart

import (
	"sort"
)

func rearrangeString(s string, k int) string {
	if k <= 1 {
		return s
	}
	result := []byte(s)
	pairs := count(result)
	cmp := func(i, j int) bool {
		if pairs[i].count == pairs[j].count {
			return pairs[i].char < pairs[j].char
		}
		return pairs[i].count > pairs[j].count
	}
	sort.Slice(pairs, cmp)
	j := 0
	for pairs[0].count > 0 {
		for i := 0; i < k; i++ {
			if pairs[0].count == 0 {
				break
			}
			if i >= len(pairs) || pairs[i].count == 0 {
				return ""
			}
			result[j] = pairs[i].char
			j++
			pairs[i].count--
		}
		sort.Slice(pairs, cmp)
	}
	return string(result)
}

type Pair struct {
	count int
	char  byte
}

func count(s []byte) []Pair {
	pairs := make([]Pair, 26)
	for _, b := range s {
		pairs[b-'a'].char = b
		pairs[b-'a'].count++
	}
	return pairs
}
