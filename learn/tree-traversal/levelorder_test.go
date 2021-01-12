/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package tree_traversal

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	if got := levelOrder(testTree); !reflect.DeepEqual(got, levelOrderWant) {
		t.Errorf("levelOrder() = %v, want %v", got, levelOrderWant)
	}
}
func Test_levelOrder1(t *testing.T) {
	if got := levelOrder1(testTree); !reflect.DeepEqual(got, levelOrderWant) {
		t.Errorf("levelOrder1() = %v, want %v", got, levelOrderWant)
	}
}

func Test_simpleLevelOrder(t *testing.T) {
	if got := simpleLevelOrder(testTree); !reflect.DeepEqual(got, simpleLevelOrderWant) {
		t.Errorf("simpleLeverOrder() = %v, want %v", got, simpleLevelOrderWant)
	}
}
