package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	testCase := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(isBalanced(testCase))
	testCase2 := &TreeNode{1, &TreeNode{2, &TreeNode{3, &TreeNode{4, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{3, nil, nil}}, &TreeNode{2, nil, nil}}
	fmt.Println(isBalanced(testCase2))
	testCase3 := &TreeNode{}
	fmt.Println(isBalanced(testCase3))
}

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	} else if root.Left == nil && root.Right == nil {
		return true
	} else if root.Left == nil && root.Right != nil {
		return isBalanced(root.Right)
	} else if root.Left != nil && root.Right == nil {
		return isBalanced(root.Left)
	}
	// node have both subtree
	if root.Left.Val == root.Right.Val {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}
