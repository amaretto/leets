package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("   -42"))
}

func myAtoi(str string) int {
	var c byte
	var result int
	var isBegin, isNegative bool
	for i := 0; i < len(str); i++ {
		// check charactor
		c = str[i]
		if c < '0' || '9' < c {
			if isBegin {
				if isNegative {
					result *= -1
				}
				return result
			}

			// pre string
			switch c {
			case ' ': // white space
				continue
			case '+': // plus
				if i+1 < len(str) && '0' <= str[i+1] && str[i+1] <= '9' {
					continue
				} else {
					return 0
				}
			case '-': // minus
				if i+1 < len(str) && '0' <= str[i+1] && str[i+1] <= '9' {
					isNegative = true
				} else {
					return 0
				}
			default:
				return result
			}
		} else {
			if !isBegin {
				isBegin = true
			}
			result *= 10
			result += int(c) - '0'
			if result > math.MaxInt32 {
				if isNegative {
					return math.MinInt32
				}
				return math.MaxInt32
			}
		}
	}

	if isNegative {
		result *= -1
	}
	return result
}

func myAtoiOld(str string) int {
	var c byte
	var result int
	var isBegin, isNegative bool
	for i := 0; i < len(str); i++ {
		// check charactor
		c = str[i]
		if c < 48 || 57 < c {
			if isBegin {
				if isNegative {
					result *= -1
				}
				return result
			}

			// pre string
			switch c {
			case 32: // white space
				continue
			case 43: // plus
				if i+1 < len(str) && 47 < str[i+1] && str[i+1] < 58 {
					continue
				} else {
					return 0
				}
			case 45: // minus
				if i+1 < len(str) && 47 < str[i+1] && str[i+1] < 58 {
					isNegative = true
				} else {
					return 0
				}
			default:
				return result
			}
		} else {
			if !isBegin {
				isBegin = true
			}
			result *= 10
			result += int(c) - 48
			if result > math.MaxInt32 {
				if isNegative {
					return math.MinInt32
				}
				return math.MaxInt32
			}
		}
	}

	if isNegative {
		result *= -1
	}
	return result
}
