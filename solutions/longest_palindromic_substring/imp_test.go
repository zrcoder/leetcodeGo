/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package longest_palindromic_substring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		input  string
		expect string
	}{
		{input: "babab", expect: "babab"},
		{input: "babad", expect: "bab"},
		{input: "cbbd", expect: "bb"},
		{input: "cbc", expect: "cbc"},
		{input: "bb", expect: "bb"},
		{input: "x", expect: "x"},
		{input: "", expect: ""},
	}
	for _, c := range cases {
		r := longestPalindrome(c.input)
		if c.expect != r {
			t.Log("input:", c.input, "expect:", c.expect, "got:", r)
		}
	}
}
