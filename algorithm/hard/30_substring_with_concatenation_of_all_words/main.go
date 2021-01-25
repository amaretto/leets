package main

func findSubstring(s string, words []string) []int {
	var result []int
	if s == "" || words == nil {
		return result
	}

	// create map[word]bool
	wlength := len(words[0])
	clength := len(words[0]) * len(words)

	for i := 0; i+clength <= len(s); i += wlength {
		tmp := s[i : i+clength]
		exists := make(map[string]bool, wlength)
		for j := 0; j+wlength <= len(tmp); j += wlength {
			word := tmp[j : j+wlength]
			if contains(word, words) && !exists[word] {
				exists[word] = true
				if len(exists) == len(words) {
					result = append(result, i)
					break
				}
			} else {
				break
			}
		}
	}

	return result
}

func contains(word string, words []string) bool {
	for _, w := range words {
		if word == w {
			return true
		}
	}
	return false
}
