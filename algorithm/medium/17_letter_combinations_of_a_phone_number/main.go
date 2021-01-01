package main

const (
	two   = "abc"
	three = "def"
	four  = "ghi"
	five  = "jkl"
	six   = "mno"
	seven = "pqrs"
	eight = "tuv"
	nine  = "wxyz"
)

func letterCombinations(digits string) []string {
	var choises, ans []string
	for _, num := range digits {
		switch num {
		case '2':
			choises = append(choises, two)
		case '3':
			choises = append(choises, three)
		case '4':
			choises = append(choises, four)
		case '5':
			choises = append(choises, five)
		case '6':
			choises = append(choises, six)
		case '7':
			choises = append(choises, seven)
		case '8':
			choises = append(choises, eight)
		case '9':
			choises = append(choises, nine)
		}
	}

	for _, choise := range choises {
		if len(ans) == 0 {
			for _, num := range choise {
				ans = append(ans, string(num))
			}
		} else {
			var tmp []string
			for _, a := range ans {
				for _, num := range choise {
					tmp = append(tmp, a+string(num))
				}
			}
			ans = tmp
		}
	}
	return ans
}
