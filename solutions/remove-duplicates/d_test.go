/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package remove_duplicates

import "testing"

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "", nums: []int{1, 1, 2}, want: 2},
		{name: "", nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
