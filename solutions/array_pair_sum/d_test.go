/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package array_pair_sum

import "testing"

func TestArrayPairSum(t *testing.T) {
	nums := []int{1, 4, 3, 2}
	if arrayPairSum(nums) != 4 {
		t.Error("failed")
	}
	t.Log(nums)
}
