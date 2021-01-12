/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package integer_to_roman

import (
	"strings"
	"testing"
)

func Test_intToRoman(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{num: 3, want: "III"},
		{num: 4, want: "IV"},
		{num: 9, want: "IX"},
		{name: "58", num: 58, want: "LVIII"},
		{name: "1994", num: 1994, want: "MCMXCIV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isValid(s string) bool {
	left, right := "([{", ")]}"
	var stack []rune
	for _, v := range s {
		if strings.ContainsRune(left, v) {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 || strings.IndexRune(left, stack[len(stack)-1]) != strings.IndexRune(right, v) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
