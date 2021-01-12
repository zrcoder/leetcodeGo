/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package add_binary

import "testing"

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{a: "10", b: "1"}, want: "11"},
		{args: args{a: "11", b: "1"}, want: "100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{a: 4, b: 5}, want: 9},
		{args: args{a: 4, b: -5}, want: -1},
		{args: args{a: -4, b: -5}, want: -9},
		{args: args{a: -4, b: 5}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
