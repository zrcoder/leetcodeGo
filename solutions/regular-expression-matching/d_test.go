/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package regular_expression_matching

import "testing"

func Test_isMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		{s: "aa", p: "a", want: false},
		{s: "aa", p: "a*", want: true},
		{s: "ab", p: ".*", want: true},
		{s: "aab", p: "c*a*b", want: true},
		{s: "mississippi", p: "mis*is*p*."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.s, tt.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
