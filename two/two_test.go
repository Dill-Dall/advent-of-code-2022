package two_test

import (
	"2022/two"
	"strconv"
	"testing"
)

func TestPartOneSample(t *testing.T) {
	result := two.PartOne("test_input.txt")
	if result != 15 {
		t.Fatalf("Part one failed was not %v -- instead it was %v\n", 15, result)
	}
}

func PartOneTest(t *testing.T) {
	result := two.PartOne("input.txt")
	println("Result is: " + strconv.Itoa(result))
}

func TestPartTwoSample(t *testing.T) {
	result := two.PartTWO("test_input.txt")
	if result != 12 {
		t.Fatalf("Part one failed was not %v -- instead it was %v\n", 12, result)
	}
}

func TestPartTwo(t *testing.T) {
	result := two.PartTWO("input.txt")
	println("Result is: " + strconv.Itoa(result))
}
