/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package ii

import (
	"testing"
)

func TestMyCalendarTwo_Book(t *testing.T) {
	cases := []struct {
		intervals [][]int
		want      []bool
	}{
		{
			intervals: [][]int{{10, 20}, {50, 60}, {10, 40}, {5, 15}},
			want:      []bool{true, true, true, false},
		},
		{
			intervals: [][]int{{24, 40}, {43, 50}, {27, 43}, {5, 21}, {30, 40}},
			want:      []bool{true, true, true, true, false},
		},
	}
	for _, c := range cases {
		myCal := Constructor()
		for i, v := range c.intervals {
			got := myCal.Book(v[0], v[1])
			if got != c.want[i] {
				t.Errorf("%v, got %v, want %v", v, got, c.want[i])
			}
		}
	}
}
