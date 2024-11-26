package AsciiArtWeb

import "fmt"

// The convToArt function takes the input string and ascii art TXT filename and converts the input string to ascii art by accessing the corresponding lines from the TXT file. we accessed the lineNum using the integer value of the current character minus the integer of the space character, multiplied by a factor of 9 (since the TXT file has 8 lines per character + 1 space in between) and incremented by n for every loop.
func ConvToArt(input string, fileName string) (string, error) {
	result := ""
	line := ""
	var err error

	inputLines := splitInputToLines(input)
	artLines, err := ReadArtInput(fileName)
	if err != nil {
		return "", fmt.Errorf("Cannot read art input")
	}

	for _, singleInputLine := range inputLines {

		if singleInputLine == "" {
			result = result + "\n"
			continue
		}

		for n := 1; n < 9; n++ {

			for _, c := range singleInputLine {
				lineNum := (int(c)-int(' '))*9 + n
				line = line + artLines[lineNum]
			}
			line = line + "\n"
			result = result + line
			line = ""
		}
	}
	return result, nil
}
