/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package path_sum

import (
	"reflect"
	"testing"
)

func Test_getPath(t *testing.T) {
	caps := []int{10, 2, 4, 3, 5, 10, 2, 18, 9, 7, 2, 2, 1, 3, 12, 1, 8, 6, 2, 2}
	relations := map[int][]int{
		0:  {1, 2, 3, 4},
		2:  {5},
		3:  {11, 12, 13},
		4:  {6, 7},
		6:  {9},
		7:  {8, 10},
		13: {14, 16, 17},
		16: {15},
		17: {18, 19},
	}
	const s = 24
	want := []string{
		"10 5 2 7",
		"10 4 10",
		"10 3 3 6 2",
		"10 3 3 6 2",
	}
	if got := getPath(caps, relations, s); !reflect.DeepEqual(got, want) {
		t.Errorf("getPath() = %v, want %v", got, want)
	}
}
