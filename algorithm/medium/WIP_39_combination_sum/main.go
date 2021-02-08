package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	nm := make(map[int]int, len(candidates))
	var na []int
	var result [][]int

	for _, c := range candidates {
		nm[c]++
		if na == nil || na[len(na)-1] != c {
			na = append(na, c)
		}
	}
	fmt.Println(nm)
	fmt.Println(na)

	for _, n := range na {
		if n > target {
			break
		}
	}

	return result
}
