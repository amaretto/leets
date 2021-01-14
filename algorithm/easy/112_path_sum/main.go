package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	return search(root, 0, sum)
}

func search(root *TreeNode, tmpSum, sum int) bool {
	tmpSum += root.Val
	if root.Left == nil && root.Right == nil {
		if tmpSum == sum {
			return true
		}
		return false
	}
	if root.Left == nil {
		return search(root.Right, tmpSum, sum)
	}
	if root.Right == nil {
		return search(root.Left, tmpSum, sum)
	}

	return search(root.Right, tmpSum, sum) || search(root.Left, tmpSum, sum)
}
