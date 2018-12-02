package frequencies

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFrequencies(fileName string) (freqs []int64, err error) {
	fileHandle, err := os.Open(fileName)
	if err != nil {
		return
	}

	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)

	for scanner.Scan() {
		text := scanner.Text()
		num, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			fmt.Printf("converting %s to int64: %v\n", text, err)
			continue
		}

		freqs = append(freqs, num)
	}

	return
}

func AddFrequencies(startingValue int64, numbers []int64) int64 {
	value := startingValue

	for _, num := range numbers {
		value += num
	}

	return value
}

func Calibration(startingValue int64, numbers []int64) int64 {
	value := startingValue

	freqsReached := make(map[int64]struct{})

	for {
		for _, num := range numbers {
			value += num

			_, ok := freqsReached[value]
			if !ok {
				freqsReached[value] = struct{}{}
			} else {
				return value
			}
		}
	}
}
