package main

func fixQuotes(text string) string {
	runes := []rune(text)
	result := []rune{}

	doubleOpen := true // tracks whether the next " is opening or closing
	singleOpen := true // tracks whether the next ' is opening or closing

	for i := 0; i < len(runes); i++ {
		c := runes[i]

		if c == '"' {
			if doubleOpen {
				// opening quote: strip trailing space just added before it
				for len(result) > 0 && result[len(result)-1] == ' ' {
					result = result[:len(result)-1]
				}
				result = append(result, c)
				// then also skip any space right after it in the source
				for i+1 < len(runes) && runes[i+1] == ' ' {
					i++
				}
			} else {
				// closing quote: strip space just before it
				for len(result) > 0 && result[len(result)-1] == ' ' {
					result = result[:len(result)-1]
				}
				result = append(result, c)
			}
			doubleOpen = !doubleOpen
			continue
		}

		if c == '\'' {
			if singleOpen {
				for len(result) > 0 && result[len(result)-1] == ' ' {
					result = result[:len(result)-1]
				}
				result = append(result, c)
				for i+1 < len(runes) && runes[i+1] == ' ' {
					i++
				}
			} else {
				for len(result) > 0 && result[len(result)-1] == ' ' {
					result = result[:len(result)-1]
				}
				result = append(result, c)
			}
			singleOpen = !singleOpen
			continue
		}

		result = append(result, c)
	}
	return string(result)
}
