/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package is_graph_bipartite

import "testing"

func Test_isBipartite(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		want  bool
	}{
		{
			name:  "case 1",
			graph: [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}},
			want:  true,
		},
		{
			name:  "case 2",
			graph: [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}},
			want:  false,
		},
		{
			name:  "case 3",
			graph: [][]int{{2, 3}, {2}, {0, 1}, {0}},
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBipartite(tt.graph); got != tt.want {
				t.Errorf("isBipartite() = %v, want %v", got, tt.want)
			}
		})
	}
}
