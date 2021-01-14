package main

import (
	"testing"
)

var testCases = []struct {
	root   *TreeNode
	want   int
	result bool
}{
	{&TreeNode{5, &TreeNode{4, &TreeNode{11, &TreeNode{7, nil, nil}, &TreeNode{2, nil, nil}}, nil}, &TreeNode{8, &TreeNode{13, nil, nil}, &TreeNode{4, nil, &TreeNode{1, nil, nil}}}}, 22, true},
}

func TestIsSymmetric(t *testing.T) {
	for _, testCase := range testCases {
		result := hasPathSum(testCase.root, testCase.want)
		if result != testCase.result {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
