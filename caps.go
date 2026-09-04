package main

import "unicode"

func capitalizeSentences(text string) string {
	runes := []rune(text)
	capitalizeNext := true

	for i := 0; i < len(runes); i++ {
		c := runes[i]

		if c == ' ' || c == '\t' {
			continue
		}

		if capitalizeNext && unicode.IsLetter(c) {
			runes[i] = unicode.ToUpper(c)
			capitalizeNext = false
		} else if !unicode.IsSpace(c) {
			capitalizeNext = false
		}

		if c == '.' || c == '!' || c == '?' || c == '\n' {
			capitalizeNext = true
		}
	}
	return string(runes)
}
