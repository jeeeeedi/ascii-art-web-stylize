package AsciiArtWeb

import (
	"bufio"
	"fmt"
	"os"
)

// open the file, scan each line of the file content, and return the TXT file as a slice of strings
func ReadArtInput(artFile string) ([]string, error) {
	artFileContent, err := os.Open(artFile)

	if err != nil {
		return nil, fmt.Errorf("Input file doesn't exist")
	}
	defer artFileContent.Close()

	scanner := bufio.NewScanner(artFileContent)
	var artText []string
	for scanner.Scan() {
		artText = append(artText, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Cannot read art source file")
	}
	return artText, nil
}
