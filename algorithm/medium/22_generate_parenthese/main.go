package main

import "math"

func generateParenthesis(n int) []string {
	var ans []string
	if n == 1 {
		ans = append(ans, "()")
		return ans
	}

	for i := int64(math.Pow(2, float64(n)*2-2)); i > 0; i-- {
		if int(i) == n*2-2 {

		}
	}

	return nil
}
