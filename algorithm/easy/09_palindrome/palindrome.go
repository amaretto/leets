package main

func isPalindrome(x int) bool {
	y := x
	var num int
	if y < 0 {
		return false
	}

	for ; y != 0; y /= 10 {
		num *= 10
		num += y % 10
	}
	if x == num {
		return true
	}
	return false
}
