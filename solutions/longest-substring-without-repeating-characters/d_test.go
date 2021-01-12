/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package longest_substring_without_repeating_characters

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		input  string
		expect int
	}{
		{input: "abcabcbb", expect: 3},
		{input: "bbbbb", expect: 1},
		{input: "pwwkew", expect: 3},
		{input: "abcdefg123456789999", expect: 16},
		{input: "abcadefg123a456789999", expect: 14},
		{input: "", expect: 0},
		{input: "x", expect: 1},
	}
	for _, c := range cases {
		r := lengthOfLongestSubstring(c.input)
		if c.expect != r {
			t.Error("input:", c.input, "expect:", c.expect, "got:", r)
		}
	}
}
