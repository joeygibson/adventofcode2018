package main

import (
	"fmt"
	"github.com/joeygibson/adventofcode2018/day1/pkg/frequencies"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: day1-1 <data file>")
		os.Exit(1)
	}

	fileName := os.Args[1]

	numbers, err := frequencies.ReadFrequencies(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	value := frequencies.AddFrequencies(0, numbers)

	fmt.Printf("Frequencies from %s: %d\n", fileName, value)

	value = frequencies.Calibration(0, numbers)

	fmt.Printf("Calibration value from %s: %d\n", fileName, value)
}
