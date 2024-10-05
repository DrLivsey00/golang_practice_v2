package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	for i := 0; i < len(word); i++ {
		if word[i] != '-' && word[i] != ' ' {
			if strings.Count(word, string(word[i])) > 1 {
				return false
			}
		}
	}
	return true
}
