/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package count_primes

import "testing"

func Test_countPrimes(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{n: 10, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimes(tt.n); got != tt.want {
				t.Errorf("countPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
