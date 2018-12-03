package util

import (
	"bufio"
	"os"
)

/// Open the given file, read all its lines, and return them as an array of strings.
func ReadValues(fileName string) (input []string, err error) {
	fileHandle, err := os.Open(fileName)
	if err != nil {
		return
	}

	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)

	for scanner.Scan() {
		text := scanner.Text()
		input = append(input, text)
	}

	return
}
