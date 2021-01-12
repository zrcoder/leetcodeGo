/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package lexicographically_smallest_equivalent_string

import "testing"

func Test_smallestEquivalentString(t *testing.T) {
	type args struct {
		A string
		B string
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{A: "parker", B: "morris", S: "parser"}, want: "makkek"},
		{name: "2", args: args{A: "hello", B: "world", S: "hold"}, want: "hdld"},
		{name: "3", args: args{A: "leetcode", B: "programs", S: "sourcecode"}, want: "aauaaaaada"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestEquivalentString(tt.args.A, tt.args.B, tt.args.S); got != tt.want {
				t.Errorf("smallestEquivalentString() = %v, want %v", got, tt.want)
			}
		})
	}
}
