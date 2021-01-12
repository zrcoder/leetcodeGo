/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package remove_duplicate_letters

import "testing"

func TestRemoveDuplicateLetters(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "1",
			s:    "bcabc",
			want: "abc",
		},
		{
			name: "2",
			s:    "cbacdcbc",
			want: "acdb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateLetters(tt.s); got != tt.want {
				t.Errorf("RemoveDuplicateLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
