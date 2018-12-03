package util

import (
	"path/filepath"
	"testing"
)

func TestReadValues(t *testing.T) {
	fileName := filepath.Join("testdata", "test0.txt")

	lines, err := ReadValues(fileName)
	if err != nil {
		t.Fatalf("opening %s: %v", fileName, err)
	}

	expectedLineCount := 12
	lineCount := len(lines)

	if lineCount != expectedLineCount {
		t.Fatalf("wrong line count; expected: %d, got: %d", expectedLineCount, lineCount)
	}
}

func TestReadValuesWithEmptyFile(t *testing.T) {
	fileName := filepath.Join("testdata", "empty.txt")

	lines, err := ReadValues(fileName)
	if err != nil {
		t.Fatalf("opening %s: %v", fileName, err)
	}

	expectedLineCount := 0
	lineCount := len(lines)

	if lineCount != expectedLineCount {
		t.Fatalf("wrong line count; expected: %d, got: %d", expectedLineCount, lineCount)
	}
}