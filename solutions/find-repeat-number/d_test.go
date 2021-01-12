/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package find_repeat_number

import "testing"

func Test_findRepeatNumber1(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{nums: []int{1, 0, 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumber1(tt.nums); got != tt.want {
				t.Errorf("findRepeatNumber1() = %v, want %v", got, tt.want)
			}
		})
	}
}
