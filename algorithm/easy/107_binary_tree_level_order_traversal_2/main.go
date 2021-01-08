package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	ans = merge(levelOrderBottom(root.Left), levelOrderBottom(root.Right))
	return append(ans, []int{root.Val})
}

func merge(l, r [][]int) [][]int {
	for i := 0; len(l)-i-1 >= 0 && len(r)-i-1 >= 0; i++ {
		if len(l) > len(r) {
			l[len(l)-i-1] = append(l[len(l)-i-1], r[len(r)-i-1]...)
		} else {
			r[len(r)-i-1] = append(l[len(l)-i-1], r[len(r)-i-1]...)
		}
	}
	if len(l) > len(r) {
		return l
	}
	return r
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
