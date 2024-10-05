package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {

	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}
	for _, char := range id {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	sum := 0
	for i, char := range id {
		digit, _ := strconv.Atoi(string(char))
		if (len(id)-i)%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return sum%10 == 0
}
