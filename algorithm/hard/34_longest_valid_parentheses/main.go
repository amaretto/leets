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
