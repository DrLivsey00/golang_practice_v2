package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	distance := 0
	if len(a) != len(b) {
		return 0, errors.New("inputs are not the same length")
	}
	runes1 := []rune(a)
	runes2 := []rune(b)

	for i, runee := range runes1 {
		if runee != runes2[i] {
			distance++
		}

	}

	return distance, nil
}
