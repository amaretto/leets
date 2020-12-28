package main

// Reverse is
func Reverse(x int) int {
	sign := 1
	var result int

	if x < 0 {
		sign = -1
		x = x * -1
	}

	for ; x != 0; x /= 10 {
		result *= 10
		result += x % 10
	}

	if result > 2147483647 {
		return 0
	}

	return sign * result
}
