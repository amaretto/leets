package main

func longestCommonPrefix(strs []string) string {

	var prefix string

	for _, str := range strs {
		if str == "" {
			return ""
		}
		if prefix == "" {
			prefix = str
			continue
		}
		for i, c := range str {
			if len(prefix)-1 < i {
				break
			}
			if rune(prefix[i]) == c {
				continue
			} else {
				prefix = prefix[:i]
				break
			}
		}
		if prefix == "" {
			return ""
		}
		if len(prefix) > len(str) {
			prefix = str
		}
	}

	return prefix
}
