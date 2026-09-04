package main

func fixQuotes(text string) string {
	runes := []rune(text)
	result := []rune{}

	doubleOpen := true
	singleOpen := true

	for i := 0; i < len(runes); i++ {
		c := runes[i]

		if c == '"' || c == '\'' {
			isDouble := c == '"'
			isOpening := (isDouble && doubleOpen) || (!isDouble && singleOpen)

			if isOpening {
				// keep whatever spacing came before it (don't touch result)
				result = append(result, c)
				// skip any spaces right after the opening quote
				for i+1 < len(runes) && runes[i+1] == ' ' {
					i++
				}
			} else {
				// closing quote: strip trailing spaces before it
				for len(result) > 0 && result[len(result)-1] == ' ' {
					result = result[:len(result)-1]
				}
				result = append(result, c)
			}

			if isDouble {
				doubleOpen = !doubleOpen
			} else {
				singleOpen = !singleOpen
			}
			continue
		}

		result = append(result, c)
	}
	return string(result)
}
