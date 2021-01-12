/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package check_if_a_number_is_majority_element_in_a_sorted_array

import "testing"

func Test_isMajorityElement(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{nums: []int{2, 4, 5, 5, 5, 5, 5, 6, 6}, target: 5}, want: true},
		{args: args{nums: []int{10, 100, 101, 101}, target: 101}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMajorityElement(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("isMajorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getInsertionPos(nums []int, target int, isFromRight bool) int {
	if isFromRight {
		return searchFromRight(nums, target)
	}
	return search(nums, target)
}

func Test_getInsertionPos(t *testing.T) {
	type args struct {
		nums      []int
		target    int
		fromRight bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{nums: []int{5, 5, 7, 7, 8}, target: 7, fromRight: true},
			want: 4,
		},
		{
			args: args{nums: []int{5, 5, 7, 7, 8}, target: 5, fromRight: true},
			want: 2,
		},
		{
			args: args{nums: []int{5, 5, 7, 7}, target: 7, fromRight: true},
			want: 4,
		},
		{
			args: args{nums: []int{5, 5, 7, 7, 8}, target: 7, fromRight: false},
			want: 2,
		},
		{
			args: args{nums: []int{5, 5, 7, 7}, target: 5, fromRight: false},
			want: 0,
		},
		{
			args: args{nums: []int{5, 5, 7, 7}, target: 6, fromRight: false},
			want: 2,
		},
		{
			args: args{nums: []int{5, 5, 7, 7}, target: 6, fromRight: true},
			want: 2,
		},
		{
			args: args{nums: []int{5, 5, 7, 7}, target: 3, fromRight: true},
			want: 0,
		},
		{
			args: args{nums: []int{5, 5, 7, 7}, target: 8, fromRight: true},
			want: 4,
		},
		{
			args: args{nums: []int{5, 5, 7, 7}, target: 3, fromRight: false},
			want: 0,
		},
		{
			args: args{nums: []int{5, 5, 7, 7}, target: 8, fromRight: false},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getInsertionPos(tt.args.nums, tt.args.target, tt.args.fromRight); got != tt.want {
				t.Errorf("getInsertionPos() = %v, want %v", got, tt.want)
			}
		})
	}
}
