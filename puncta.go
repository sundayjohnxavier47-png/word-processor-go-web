package main

func collapseSpaces(text string) string {
	runes := []rune(text)
	result := []rune{}

	for i := 0; i < len(runes); i++ {
		c := runes[i]
		if c == ' ' {
			if len(result) > 0 && result[len(result)-1] == ' ' {
				continue
			}
		}
		result = append(result, c)
	}
	return string(result)
}

func fixPunctuationSpacing(text string) string {
	runes := []rune(text)
	result := []rune{}
	punctuation := map[rune]bool{',': true, '.': true, '!': true, '?': true, ':': true, ';': true}

	for i := 0; i < len(runes); i++ {
		c := runes[i]
		if punctuation[c] {
			for len(result) > 0 && result[len(result)-1] == ' ' {
				result = result[:len(result)-1]
			}
			result = append(result, c)
			if i+1 < len(runes) {
				next := runes[i+1]
				if next != ' ' && next != '\n' && !punctuation[next] {
					result = append(result, ' ')
				}
			}
		} else {
			result = append(result, c)
		}
	}
	return string(result)
}
