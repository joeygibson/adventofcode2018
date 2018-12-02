package inventory

import (
	"testing"
)

func TestScanLabelForTwos(t *testing.T) {
	data := []struct {
		str     string
		matches bool
	}{
		{"abcdef", false},
		{"bababc", true},
		{"abbcde", true},
		{"abcccd", false},
		{"aabcdd", true},
		{"abcdee", true},
		{"ababab", false},
	}

	for _, item := range data {
		matches := ScanLabel(2, item.str)

		if matches != item.matches {
			t.Fatalf("%s, expected %t, got %t", item.str, item.matches, matches)
		}
	}
}

func TestScanLabelForThrees(t *testing.T) {
	data := []struct {
		str     string
		matches bool
	}{
		{"abcdef", false},
		{"bababc", true},
		{"abbcde", false},
		{"abcccd", true},
		{"aabcdd", false},
		{"abcdee", false},
		{"ababab", true},
	}

	for _, item := range data {
		matches := ScanLabel(3, item.str)

		if matches != item.matches {
			t.Fatalf("%s, expected %t, got %t", item.str, item.matches, matches)
		}
	}
}

func TestScanLabels(t *testing.T) {
	data := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}

	expectedValue := 4
	value := ScanLabels(2, data)

	if value != expectedValue {
		t.Fatalf("expected %d, got %d", expectedValue, value)
	}

	expectedValue = 3
	value = ScanLabels(3, data)

	if value != expectedValue {
		t.Fatalf("expected %d, got %d", expectedValue, value)
	}
}

func TestChecksum(t *testing.T) {
	data := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}

	expectedValue := 12
	value := Checksum(data)

	if value != expectedValue {
		t.Fatalf("expected %d, got %d", expectedValue, value)
	}
}

func TestDiffByOne(t *testing.T) {
	data := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}

	expectedResult := []string{
		"fghij",
		"fguij",
	}

	result := DiffByOne(data)

	if len(result) != len(expectedResult) {
		t.Fatalf("wrong number of results; expected: %d, got %d", len(expectedResult),
			len(result))
	}

FirstLoop:
	for _, v := range expectedResult {
		for _, v1 := range result {
			if v1 == v {
				continue FirstLoop
			}
		}

		t.Errorf("didn't find %s in result", v)
	}
}

func TestGetCommonLetters(t *testing.T) {
	data := []string{
		"fghij",
		"fguij",
	}

	expectedResult := "fgij"

	result := GetCommonLetters(data[0], data[1])

	if result != expectedResult {
		t.Fatalf("expected: %s, gotL %s", expectedResult, result)
	}
}
