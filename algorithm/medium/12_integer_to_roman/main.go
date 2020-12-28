package main

func intToRoman(num int) string {
	var result string
	// M... 	1000
	for i := num / 1000; i > 0; i-- {
		result += "M"
	}

	num %= 1000
	h := num / 100

	switch {
	// CM			900
	case h == 9:
		result = result + "CM"
	// DC...	500~800
	case h >= 5:
		result = result + "D"
		for i := h - 5; i > 0; i-- {
			result = result + "C"
		}
	// CD			400
	case h == 4:
		result = result + "CD"
	// C...		100~300
	default:
		for i := h; i > 0; i-- {
			result = result + "C"
		}
	}

	num %= 100
	t := num / 10

	switch {
	// XC			90
	case t == 9:
		result = result + "XC"
	// LX...	50~80
	case t >= 5:
		result = result + "L"
		for i := t - 5; i > 0; i-- {
			result = result + "X"
		}
	// XL...	40
	case t == 4:
		result = result + "XL"
	// X...		10~30
	default:
		for i := t; i > 0; i-- {
			result = result + "X"
		}
	}

	num %= 10
	switch {
	// IX			9
	case num == 9:
		result = result + "IX"
	// VI...	5~8
	case num >= 5:
		result = result + "V"
		for i := num - 5; i > 0; i-- {
			result = result + "I"
		}
	// IV			4
	case num == 4:
		result = result + "IV"
	// I...		3
	default:
		for i := num; i > 0; i-- {
			result = result + "I"
		}
	}

	return result
}
