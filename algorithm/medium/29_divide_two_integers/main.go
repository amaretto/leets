package main

func divide(dividend int, divisor int) int {
	var isSigned bool
	var count int
	if dividend < 0 {
		dividend = 0 - dividend
		count++
	}
	if divisor < 0 {
		divisor = 0 - divisor
		count++
	}
	if count == 1 {
		isSigned = true
	}
	div := divideBit(dividend, divisor)
	if isSigned {
		return 0 - div
	}
	if div >= (1 << 31) {
		div = (1 << 31) - 1
	}
	return div

}

func divideBit(dividend int, divisor int) int {
	var count, ans int
	for {
		if divisor<<1 > dividend {
			break
		}
		divisor = divisor << 1
		count++
	}

	for ; count >= 0; count-- {
		if dividend >= divisor {
			dividend -= divisor
			ans += 1 << count
		}
		divisor = divisor >> 1
	}

	return ans
}
