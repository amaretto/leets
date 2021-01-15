package main

func generate(numRows int) [][]int {
	var pre []int
	var result [][]int
	for row := 1; row <= numRows; row++ {
		tmp := make([]int, row)
		for i := 0; i < row; i++ {
			if i == 0 || i == row-1 {
				tmp[i] = 1
			} else {
				tmp[i] = pre[i-1] + pre[i]
			}
		}
		result = append(result, tmp)
		pre = tmp
	}

	return result
}
