package main

func generateParenthesis(n int) []string {
	return search(0, 0, n, 0, "")
}

func search(l, r, limit, depth int, s string) []string {
	var result []string
	if l == limit && r == limit {
		return append(result, s)
	}
	if r+1 <= limit && depth-1 >= 0 {
		result = append(result, search(l, r+1, limit, depth-1, s+")")...)
	}

	if l+1 <= limit && depth+1 <= limit {
		result = append(result, search(l+1, r, limit, depth+1, s+"(")...)
	}

	return result
}
