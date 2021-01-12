/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package sort

// Search uses binary search to find and return the smallest index i
// in [from, to) at which f(i) is true, assuming that on the range [from, to),
// f(i) == true implies f(i+1) == true. That is, Search requires that
// f is false for some (possibly empty) prefix of the input range [from, to)
// and then true for the (possibly empty) remainder; Search returns
// the first true index. If there is no such index, Search returns n.
// It is just another version of the same name function in standard lib, package sort.
// this version is a little different of parameters
func Search(lo, hi int, f func(int) bool) int {
	for lo < hi {
		mid := int(uint(lo+hi) >> 1) // avoid overflow
		// lo ≤ mid < hi
		if !f(mid) {
			lo = mid + 1 // preserves f(from-1) == false
		} else {
			hi = mid // preserves f(to) == true
		}
	}
	// lo == hi, f(hi-1) == false, and f(hi) == true  =>  answer is lo (or hi).
	return lo
}
