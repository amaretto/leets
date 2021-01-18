package main

func getRow(rowIndex int) []int {
	var row []int
	tmp := []int{1, 1}
	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return tmp
	}
	for i := 2; i <= rowIndex; i++ {
		row = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				row[j] = 1
			} else {
				row[j] = tmp[j-1] + tmp[j]
			}
		}
		tmp = row
	}
	return row
}
