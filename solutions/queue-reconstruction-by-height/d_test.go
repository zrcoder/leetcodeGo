/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package queue_reconstruction_by_height

import (
	"reflect"
	"testing"
)

func Test_reconstructQueue(t *testing.T) {
	type args struct {
		people [][]int
	}
	tests := []struct {
		name   string
		people [][]int
		want   [][]int
	}{
		{
			name:   "1",
			people: [][]int{{8, 0}, {4, 4}, {8, 1}, {5, 0}, {6, 1}, {5, 2}},
			want:   [][]int{{5, 0}, {8, 0}, {5, 2}, {6, 1}, {4, 4}, {8, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reconstructQueue1(tt.people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconstructQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
