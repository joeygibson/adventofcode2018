package inventory

import (
	"bufio"
	"os"
	"strings"
)

func Checksum(input []string) int {
	twos := ScanLabels(2, input)
	threes := ScanLabels(3, input)

	return twos * threes
}

func ScanLabels(numberToMatch int, input []string) int {
	matchCount := 0

	for _, str := range input {
		if ScanLabel(numberToMatch, str) {
			matchCount++
		}
	}

	return matchCount
}

func ScanLabel(numberToMatch int, input string) bool {
	hits := make(map[int32]int)

	for _, c := range input {
		hits[c] += 1
	}

	for _, v := range hits {
		if v == numberToMatch {
			return true
		}
	}

	return false
}

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

func DiffByOne(input []string) (words []string) {
	wordMap := make(map[string]bool)

	for i, wordOne := range input {
		if len(wordOne) == 0 {
			continue
		}

		SecondLoop:
		for j, wordTwo := range input {
			if i == j {
				continue
			}

			if len(wordTwo) == 0 {
				continue
			}
			
			alreadyHit := false

			for ci := range wordOne {
				if wordOne[ci] != wordTwo[ci] {
					if alreadyHit {
						continue SecondLoop
					} else {
						alreadyHit = true
					}
				}
			}

			if alreadyHit {
				wordMap[wordOne] = true
				wordMap[wordTwo] = true
			}
		}
	}

	for k := range wordMap {
		words = append(words, k)
	}

	return
}

func GetCommonLetters(wordOne, wordTwo string) string {
	letters := make([]string, 0)

	for i, letter := range wordOne {
		if wordOne[i] == wordTwo[i] {
			letters = append(letters, string(letter))
		}
	}

	return strings.Join(letters, "")
}