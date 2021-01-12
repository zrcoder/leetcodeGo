package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		array  []int
		target int
		expect []int
	}{
		{array: []int{2, 7, 11, 5}, target: 9, expect: []int{0, 1}},
		{array: []int{2, 7, 11, 5}, target: 7, expect: []int{0, 3}},
		{array: []int{2, 7, 11, 5}, target: 16, expect: []int{2, 3}},
		{array: []int{2, 7, 11, 5}, target: 18, expect: []int{1, 2}},
		{array: []int{2, 7, 11, 5}, target: 100, expect: nil},
	}

	for _, c := range cases {
		r := twoSum(c.array, c.target)
		if !reflect.DeepEqual(r, c.expect) {
			t.Error("expected:", c.expect, "got:", r)
		}
	}
}
