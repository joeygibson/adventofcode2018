package frequencies

import (
	"path/filepath"
	"testing"
)

func TestReadFrequencies(t *testing.T) {
	fileName := filepath.Join("testdata", "test0.txt")

	numbers, err := ReadFrequencies(fileName)
	if err != nil {
		t.Fatalf("reading file %s: %v", fileName, err)
	}

	data := []int64{
		-10,
		23,
		-1,
		-1,
		-1,
		-1,
		-15,
		99,
		1234567,
	}

	for i := 0; i < len(data); i++ {
		if numbers[i] != data[i] {
			t.Errorf("expected %d, got %d", data[i], numbers[i])
		}
	}
}

func TestAddFrequencies(t *testing.T) {
	data := []int64{
		-10,
		23,
		-1,
		-1,
		-1,
		-1,
		-15,
		99,
		1234567,
	}

	var expectedValue int64 = 1234660

	value := AddFrequencies(0, data)
	if value != expectedValue {
		t.Fatalf("expected %d, got %d", expectedValue, value)
	}
}

func TestAddFrequenciesWithNonZeroStartingValue(t *testing.T) {
	data := []int64{
		-10,
		23,
		-1,
		-1,
		-1,
		-1,
		-15,
		99,
		1234567,
	}

	var expectedValue int64 = 1234683

	value := AddFrequencies(23, data)
	if value != expectedValue {
		t.Fatalf("expected %d, got %d", expectedValue, value)
	}
}

func TestCalibration(t *testing.T) {
	data := []int64{
		+1,
		-2,
		+3,
		+1,
	}

	var expectedValue int64 = 2
	value := Calibration(0, data)
	if value != expectedValue {
		t.Fatalf("expected %d, got %d", expectedValue, value)
	}
}
