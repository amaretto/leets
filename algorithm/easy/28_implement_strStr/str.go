package main

func strStr(haystack string, needle string) int {
	nl := len(needle)
	hl := len(haystack)
	if nl == 0 {
		return 0
	}

	for i := 0; i < hl; i++ {
		if (needle[0] == haystack[i]) && (nl+i <= hl) {
			for j := 0; j < nl; j++ {
				if needle[j] == haystack[i+j] {
					if j == nl-1 {
						return i
					}
					continue
				} else {
					break
				}
			}
		}
	}
	return -1
}
