/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package count_and_say

import "testing"

func Test_countAndSay(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{name: "n=1", n: 1, want: "1"},
		{name: "n=2", n: 2, want: "11"},
		{name: "n=3", n: 3, want: "21"},
		{name: "n=4", n: 4, want: "1211"},
		{name: "n=5", n: 5, want: "111221"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
