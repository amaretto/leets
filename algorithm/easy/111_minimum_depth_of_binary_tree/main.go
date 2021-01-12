package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	testCase := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(minDepth(testCase))
	testCase2 := &TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{4, nil, &TreeNode{5, nil, &TreeNode{6, nil, nil}}}}}
	fmt.Println(minDepth(testCase2))
	testCase3 := &TreeNode{}
	fmt.Println(minDepth(testCase3))
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	l := minDepth(root.Left)
	r := minDepth(root.Right)
	if (l != 0) && (r != 0) {
		return min(l, r) + 1
	} else if l == 0 {
		return r
	}
	return l
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
