package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abcda"))
}

func longestPalindrome(s string) string {
	max := 1
	slen := len(s)
	ts := ""
	result := ""
	var count int
	var cc, lc byte
	for i := 0; i < slen; i++ {
		if slen-(i+1) < max/2 {
			break
		}
		cc = s[i]

		//even
		if cc == lc {
			count = 2
			for j := 0; i+j+1 < slen; j++ {
				if i-2-j < 0 {
					break
				}
				if s[i-2-j] == s[i+j+1] {
					count += 2
					ts = s[i-2-j : i+j+2]
					continue
				} else {
					if count > max {
						max = count
						if count == 2 {
							ts = string(cc) + string(lc)
						}
						result = ts
					}
					break
				}
			}
			if count > max {
				if count == 2 {
					ts = string(cc) + string(lc)
				}
				result = ts
				max = count
			}
		}
		// odd
		count = 1
		for j := 0; i+j+1 < slen; j++ {
			if i-1-j < 0 {
				break
			}
			if s[i-1-j] == s[i+j+1] {
				count += 2
				ts = s[i-1-j : i+j+2]
				continue
			} else {
				if count > max {
					result = ts
					max = count
				}
				break
			}
		}
		if count > max {
			result = ts
			max = count
		}
		lc = cc
	}

	if max == 1 {
		if slen != 0 {
			result = string(cc)
		}
	}
	return result
}
