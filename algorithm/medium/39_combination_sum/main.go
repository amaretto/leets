package main

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int

	sort.Ints(candidates)

	search(target, candidates, nil, &result)

	return result
}

func search(target int, candidates, pre []int, result *[][]int) {
	tmp := make([]int, len(pre))
	copy(tmp, pre)

	if target-candidates[0] == 0 {
		tmp = append(tmp, candidates[0])
		*result = append(*result, tmp)
		return
	} else if target-candidates[0] > 0 {
		if len(candidates) > 1 {
			search(target, candidates[1:], tmp, result)
		}
		target -= candidates[0]
		tmp = append(tmp, candidates[0])
		search(target, candidates, tmp, result)
	} else {
		return
	}

}
