package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameVal(root.Left, root.Right)
}

func isSameVal(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	if isSameVal(p.Left, q.Right) && isSameVal(p.Right, q.Left) {
		return true
	}
	return false
}
