package main

import "fmt"

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

func convert(s string, numRows int) string {
	cycle := numRows + numRows - 2
	if cycle == 0 {
		return s
	}
	rowArray := make([][]byte, numRows)
	var pos int

	for i := 0; i < len(s); i++ {
		b := s[i]
		if pos = i % cycle; pos < numRows {
			rowArray[pos] = append(rowArray[pos], b)
		} else {
			rowArray[cycle-pos] = append(rowArray[cycle-pos], b)
		}
	}

	var result []byte
	for _, ra := range rowArray {
		result = append(result, ra...)
	}

	return string(result)
}

func convertOld(s string, numRows int) string {
	cycle := numRows + numRows - 2
	if cycle == 0 {
		return s
	}
	rowArray := make([]string, numRows)
	var pos int

	for i, b := range s {
		if pos = i % cycle; pos < numRows {
			rowArray[pos] = rowArray[pos] + string(b)
		} else {
			rowArray[cycle-pos] = rowArray[cycle-pos] + string(b)
		}
	}

	var result string
	for _, ra := range rowArray {
		result += ra
	}

	return result
}
