package main

func longestValidParentheses(s string) int {
	var result int
	if s == "" {
		return result
	}

	var sum, count int
	for i := 0; i < len(s); i++ {
		ss := s[i:]
		for _, c := range ss {
			if c == '(' {
				sum++
				count++
			} else if c == ')' {
				sum--
				count++
			}
			if sum < 0 {
				break
			}
			if sum == 0 && count > result {
				result = count
			}
		}
		sum = 0
		count = 0
	}

	return result
}

// This code is reffered to discuttion on LeetCode
func longestValidParenthesesImprove(s string) int {
	var result, left int
	d := make([]int, len(s))
	if s == "" {
		return 0
	}

	for i, c := range s {
		if c == '(' {
			left++
		} else if left > 0 {
			d[i] = d[i-1] + 2
			if i-d[i] > 0 {
				d[i] += d[i-d[i]]
			}
			result = max(result, d[i])
			left--
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
