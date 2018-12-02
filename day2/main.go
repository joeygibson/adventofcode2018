package main

import (
	"fmt"
	"github.com/joeygibson/adventofcode2018/day2/pkg/inventory"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: day2 <data file>")
		os.Exit(1)
	}

	fileName := os.Args[1]

	input, err := inventory.ReadValues(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	value := inventory.Checksum(input)

	fmt.Printf("Checksum: %d\n", value)

	selectedLabels := inventory.DiffByOne(input)
	commonLetters := inventory.GetCommonLetters(selectedLabels[0], selectedLabels[1])

	fmt.Printf("Initial: %d, slected: %d, common letters: %s\n",
		len(input), len(selectedLabels), commonLetters)
}
