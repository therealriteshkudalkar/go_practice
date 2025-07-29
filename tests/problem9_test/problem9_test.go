package problem9_test

import (
	"testing"

	"github.com/therealriteshkudalkar/go_practice/src/problem9"
)

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name string
		root *problem9.TreeNode
		k    int
		want int
	}{
		{
			"Test Case 1",
			&problem9.TreeNode{3, &problem9.TreeNode{1, nil, &problem9.TreeNode{2, nil, nil}}, &problem9.TreeNode{4, nil, nil}},
			1,
			1,
		},
		{
			"Test Case 2",
			&problem9.TreeNode{5, &problem9.TreeNode{3, &problem9.TreeNode{2, &problem9.TreeNode{1, nil, nil}, nil}, &problem9.TreeNode{4, nil, nil}}, &problem9.TreeNode{6, nil, nil}},
			3,
			3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := problem9.KthSmallest(test.root, test.k); got != test.want {
				t.Errorf("KthSmallest(): %v, but want: %v", got, test.want)
			}
		})
	}
}
