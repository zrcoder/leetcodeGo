/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package symmetric_tree

import "testing"

func Test_isMirror1(t *testing.T) {
	type args struct {
		t1 *TreeNode
		t2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMirror1(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("isMirror1() = %v, want %v", got, tt.want)
			}
		})
	}
}
