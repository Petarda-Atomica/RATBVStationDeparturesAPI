package main

import (
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func removeDiacritics(s string) (string, error) {
	// 1. Decompose characters (e.g., 'ș' becomes 's' + 'comma')
	// 2. Filter out non-spacing marks (the diacritics)
	// 3. Recompose into a normal string
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, err := transform.String(t, s)
	if err != nil {
		return "", err
	}
	return result, nil
}
