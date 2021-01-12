package redundant_connection

import (
	"reflect"
	"testing"
)

func Test_findRedundantDirectedConnection(t *testing.T) {
	tests := []struct {
		name  string
		edges [][]int
		want  []int
	}{
		{
			edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}},
			want:  []int{4, 1},
		},
		{
			edges: [][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}},
			want:  []int{2, 1},
		},
		{
			edges: [][]int{{1, 2}, {1, 3}, {1, 4}, {4, 3}},
			want:  []int{4, 3},
		},
		{
			edges: [][]int{{3, 2}, {1, 2}, {1, 3}, {1, 4}},
			want:  []int{1, 2},
		},
		{
			edges: [][]int{{1, 2}, {3, 2}, {1, 3}},
			want:  []int{3, 2},
		},
		{
			edges: [][]int{{3, 2}, {1, 2}, {1, 3}},
			want:  []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRedundantDirectedConnection(tt.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRedundantDirectedConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
