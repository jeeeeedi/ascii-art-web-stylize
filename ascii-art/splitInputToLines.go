package AsciiArtWeb

// take the input and convert it into a slice of strings with "\n" as the separator
func splitInputToLines(inputStr string) []string {
	inputlines := []string{}
	line := ""
	for i := 0; i < len(inputStr); i++ {

		if inputStr[i] == '\n' || inputStr[i] == '\r' {
			// Append the current "line" and reset for the next line
			inputlines = append(inputlines, line)
			line = ""
			i++

			if i == len(inputStr)-1 {
				inputlines = append(inputlines, "")
			}

			continue
		}

		line = line + string(inputStr[i])

		if i == len(inputStr)-1 {
			inputlines = append(inputlines, line)
		}

	}
	return inputlines
}
