/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package tree_traversal

import (
	"reflect"
	"testing"
)

func Test_preorder(t *testing.T) {
	if got := preorder(testTree); !reflect.DeepEqual(got, preOrderWant) {
		t.Errorf("preOrder() = %v, want1 %v", got, preOrderWant)
	}
}

func Test_preorder1(t *testing.T) {
	if got := preorder1(testTree); !reflect.DeepEqual(got, preOrderWant) {
		t.Errorf("preOrder() = %v, want1 %v", got, preOrderWant)
	}
}
