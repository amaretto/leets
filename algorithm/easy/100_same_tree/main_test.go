package main

import (
	"testing"
)

var testCases = []struct {
	p      []int
	q      []int
	Output bool
}{
	{[]int{1, 2, 3}, []int{1, 2, 3}, true},
	{[]int{1, 2}, []int{1, -100000, 2}, false},
	{[]int{1, 2, 1}, []int{1, 1, 2}, false},
}

func TestIsSameTree(t *testing.T) {
	for _, testCase := range testCases {
		// preapre p
		parray := make([]TreeNode, len(testCase.p))
		for i, num := range testCase.p {
			parray[i] = TreeNode{num, nil, nil}
		}
		for i := 0; i < len(parray); i++ {
			if 2*(i+1) < len(parray) && parray[2*(i+1)].Val != -100000 {
				parray[i].Left = &parray[2*(i+1)]
			}
			if 2*(i+1)+1 < len(parray) && parray[2*(i+1)+1].Val != -100000 {
				parray[i].Right = &parray[2*(i+1)+1]
			}
		}
		// preapre q
		qarray := make([]TreeNode, len(testCase.q))
		for i, num := range testCase.q {
			qarray[i] = TreeNode{num, nil, nil}
		}
		for i := 0; i < len(qarray); i++ {
			if 2*(i+1) < len(qarray) && qarray[2*(i+1)].Val != -100000 {
				qarray[i].Left = &qarray[2*(i+1)]
			}
			if 2*(i+1)+1 < len(qarray) && qarray[2*(i+1)+1].Val != -100000 {
				qarray[i].Right = &qarray[2*(i+1)+1]
			}
		}

		result := isSameTree(&parray[0], &qarray[0])
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
