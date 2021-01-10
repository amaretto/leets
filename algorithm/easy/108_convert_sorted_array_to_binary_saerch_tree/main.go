package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	array := []int{-10, -3, 0, 5, 9}
	ans := sortedArrayToBST(array)
	printResult(ans)
}

func printResult(t *TreeNode) {
	if t != nil {
		printResult(t.Left)
		fmt.Printf("%d,", t.Val)
		printResult(t.Right)
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	var tmp TreeNode
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		tmp.Val = nums[0]
	} else {
		tmp.Val = nums[len(nums)/2]
		tmp.Left = sortedArrayToBST(nums[:len(nums)/2])
		tmp.Right = sortedArrayToBST(nums[len(nums)/2+1:])
	}
	return &tmp
}
