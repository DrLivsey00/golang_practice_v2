package strand

import "strings"

var transcription = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	runes := []rune(dna)
	var rna strings.Builder
	for _, v := range runes {
		rna.WriteRune(transcription[v])
	}
	return rna.String()
}
