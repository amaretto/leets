package main

func isValid(s string) bool {
	var parentheses = map[byte]byte{
		byte('{'): byte('}'),
		byte('('): byte(')'),
		byte('['): byte(']'),
	}

	var stuck []byte

	if len(s)%2 == 1 {
		return false
	}

	var c byte
	for i := 0; i < len(s); i++ {
		c = s[i]
		if c == byte('{') || c == byte('[') || c == byte('(') {
			stuck = append(stuck, c)
		} else {
			if len(stuck) == 0 {
				return false
			}
			if c == parentheses[stuck[len(stuck)-1]] {
				stuck = stuck[:len(stuck)-1]
				continue
			} else {
				return false
			}
		}
	}
	if len(stuck) != 0 {
		return false
	}
	return true
}
