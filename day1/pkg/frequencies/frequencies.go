package frequencies

import (
	"fmt"
	"github.com/joeygibson/adventofcode2018/util"
	"strconv"
)

func ReadFrequencies(fileName string) (freqs []int64, err error) {
	lines, err := util.ReadValues(fileName)
	if err != nil {
		return
	}

	for _, val := range lines {
		num, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			fmt.Printf("converting %s to int64: %v\n", val, err)
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
