/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package before_and_after_puzzle

import (
	"reflect"
	"testing"
)

func Test_beforeAndAfterPuzzles(t *testing.T) {
	tests := []struct {
		name    string
		phrases []string
		want    []string
	}{
		{
			name:    "test1",
			phrases: []string{"writing code", "code rocks"},
			want:    []string{"writing code rocks"},
		},
		{
			name: "test2",
			phrases: []string{
				"mission statement",
				"a quick bite to eat",
				"a chip off the old block",
				"chocolate bar",
				"mission impossible",
				"a man on a mission",
				"block party",
				"eat my words",
				"bar of soap",
			},
			want: []string{
				"a chip off the old block party",
				"a man on a mission impossible",
				"a man on a mission statement",
				"a quick bite to eat my words",
				"chocolate bar of soap",
			},
		},
		{
			name:    "test3",
			phrases: []string{"a", "b", "a"},
			want:    []string{"a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := beforeAndAfterPuzzles(tt.phrases); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("beforeAndAfterPuzzles() = %v, want %v", got, tt.want)
			}
		})
	}
}
