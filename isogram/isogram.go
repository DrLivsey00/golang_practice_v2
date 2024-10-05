package isogram

import "unicode"

type empty struct{}

func IsIsogram(word string) bool {

	checked := map[rune]empty{}

	for _, ch := range word {
		letter := unicode.ToLower(ch)
		if unicode.IsLetter(ch) {
			if _, ok := checked[letter]; ok {
				return false
			}
		}
		checked[letter] = empty{}
	}

	return true
}
