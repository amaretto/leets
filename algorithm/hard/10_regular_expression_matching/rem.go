package main

import "fmt"

func main() {
	fmt.Println(isMatch("ab", ".*c"))
}

func isMatch(s string, p string) bool {
	var pos int
	var e, tmp, next byte

	for i := 0; i < len(p); i++ {
		e = p[i]
		if e == '*' { // repeat
			if tmp == '.' {
				if i == len(p)-1 {
					return true
				}
				next = p[i+1]
				for pos < len(s) && s[pos] != next {
					pos++
				}
			} else {
				for pos < len(s) && s[pos] == tmp {
					pos++
				}
			}
		} else { // . or number
			if e != '.' && e != s[pos] {
				if i < len(p)-1 {
					if next = p[i+1]; next != '*' {
						return false
					}
				}
			} else {
				pos++
			}
		}
		tmp = e
		if pos >= len(s) && i < len(p)-1 {
			next = p[i+1]
			if next == '*' {
				continue
			}
			return false
		}
	}
	if pos < len(s) {
		return false
	}
	return true
}

func isMatchOld(s string, p string) bool {
	var e, tmp byte
	var pos int

	for i := 0; i < len(p); i++ {
		e = p[i]
		if pos >= len(s) {
			if e == '*' {
				continue
			}
			return false
		}

		if e == '.' { // single charactor
			tmp = '.'
			pos++
		} else if e == '*' { // repeat
			if tmp == '.' {
				return true
			}
			for pos < len(s) && s[pos] != tmp {
				pos++
			}
		} else { // number
			if s[pos] != e {
				return false
			}
			tmp = e
			pos++
		}
	}
	if pos < len(s) {
		return false
	}
	return true
}
