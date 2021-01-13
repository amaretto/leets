package main

import (
	"testing"
)

var testCases = []struct {
	root *TreeNode
	want int
}{
	{&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}, 2},
	{&TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{4, nil, &TreeNode{5, nil, &TreeNode{6, nil, nil}}}}}, 5},
}

func TestIsSymmetric(t *testing.T) {
	for _, testCase := range testCases {
		result := minDepth(testCase.root)
		if result != testCase.want {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
