/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package merge_sorted_array

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1: []int{2, 3, 0, 0, 0},
			m:     2,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{2, 2, 3, 5, 6},
		},
		{
			nums1: []int{2, 3, 0, 0, 0},
			m:     2,
			nums2: []int{5, 6},
			n:     2,
			want:  []int{2, 3, 5, 6},
		},
		{
			nums1: []int{3, 0, 0, 0},
			m:     1,
			nums2: []int{5, 6},
			n:     2,
			want:  []int{3, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			if !reflect.DeepEqual(tt.nums1[:tt.n+tt.m], tt.want) {
				t.Errorf("got: %v, want: %v", tt.nums1[:tt.n+tt.m], tt.want)
			}
		})
	}
}
