package main

import (
	"testing"
)

var testCases = []struct {
	tree   []int
	Output bool
}{
	{[]int{1, 2, 2, 3, 4, 4, 3}, true},
	{[]int{1, 2, 2, -100000, 3, -100000, 3}, true},
}

func TestIsSymmetric(t *testing.T) {
	for _, testCase := range testCases {
		// preapre root
		tarray := make([]TreeNode, len(testCase.tree))
		for i, num := range testCase.tree {
			tarray[i] = TreeNode{num, nil, nil}
		}
		for i := 0; i < len(tarray); i++ {
			if 2*(i+1) < len(tarray) && tarray[2*(i+1)].Val != -100000 {
				tarray[i].Left = &tarray[2*(i+1)]
			}
			if 2*(i+1)+1 < len(tarray) && tarray[2*(i+1)+1].Val != -100000 {
				tarray[i].Right = &tarray[2*(i+1)+1]
			}
		}

		result := isSymmetric(&tarray[0])
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
