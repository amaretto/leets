package main

import (
	"fmt"
	"reflect"
	"sort"
)

func nextPermutation(nums []int) {
	var num int
	for i := len(nums) - 1; i >= 0; i-- {
		num = nums[i]
		for j := i - 1; j >= 0; j-- {
			tmp := nums[j]
			if num == tmp {
				continue
			} else if num > tmp { // [1,2,3]->[1,3,2], [2,2,3,5,4,8] -> [2,2,3,5,8,4], [2,2,4,2,3,3] -> [2,2,4,3,3,2]
				nums[i], nums[j] = nums[j], nums[i]
				return
			}
		}
	}
}

// ######## bellow code exceeded time limit ###########
func nextPermutationOld(nums []int) {
	head := []int{}
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.SliceStable(tmp, func(i, j int) bool { return tmp[i] < tmp[j] })
	result := perm(head, tmp)
	result = removeDup(result)

	for i, a := range result {
		if reflect.DeepEqual(a, nums) {
			if i == len(result)-1 {
				copy(nums, result[0])
				return
			}
			copy(nums, result[i+1])
			return
		}
	}

}

func perm(head, rest []int) [][]int {
	var res [][]int
	if len(rest) == 0 {
		res = append(res, head)
		return res
	}
	for i := 0; i < len(rest); i++ {
		restx := make([]int, len(rest))
		headx := make([]int, len(head)) //
		copy(restx, rest)
		copy(headx, head) //
		headx = append(headx, restx[i])
		restx = append(restx[:i], restx[i+1:]...)
		res = append(res, perm(headx, restx)...)
	}
	return res
}

func removeDup(a [][]int) [][]int {
	b := make([][]int, len(a))
	copy(b, a)
	exist := make(map[string]bool, len(a))
	for i := 0; i < len(b); {
		if !exist[fmt.Sprint(b[i])] {
			exist[fmt.Sprint(b[i])] = true
			i++
		} else {
			b = append(b[:i], b[i+1:]...)
		}
	}
	return b
}
