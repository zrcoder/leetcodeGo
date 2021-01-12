/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package cherry_pickup

import "testing"

func Test_cherryPickup(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "case 1",
			grid: [][]int{
				{0, 1, -1},
				{1, 0, -1},
				{1, 1, 1}},
			want: 5,
		},
		{
			name: "case 2",
			grid: [][]int{
				{1, 1, 1, 0, 0},
				{0, 0, 1, 0, 1},
				{1, 0, 1, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 1, 1, 1},
			},
			want: 11,
		},
		{
			name: "case 3",
			grid: [][]int{
				{1, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 1},
				{1, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 1, 1, 1, 1},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cherryPickup(tt.grid); got != tt.want {
				t.Errorf("cherryPickup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cherryPickupOnce(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "case 1",
			grid: [][]int{
				{0, 1, -1},
				{1, 0, -1},
				{1, 1, 1}},
			want: 4,
		},
		{
			name: "case 2",
			grid: [][]int{
				{1, 1, 1, 0, 0},
				{0, 0, 1, 0, 1},
				{1, 0, 1, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 1, 1, 1},
			},
			want: 9,
		},
		{
			name: "case 3",
			grid: [][]int{
				{1, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 1},
				{1, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 1, 1, 1, 1},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cherryPickupOnce(tt.grid); got != tt.want {
				t.Errorf("cherryPickupOnce() = %v, want %v", got, tt.want)
			}
		})
	}
}
