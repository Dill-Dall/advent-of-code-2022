package four_test

import (
	"2022/four"
	"strconv"
	"testing"
)

func TestPartOne(t *testing.T) {
	expected := 2
	test_result := four.PartOne("test_input.txt")
	if test_result != expected {
		t.Fatalf("Part one failed was not %v -- instead it was %v\n", expected, test_result)
	} else {
		println("Sample ok")
		t.Log("SAMPLE OK")
	}
	result := four.PartOne("input.txt")
	println("Result is: " + strconv.Itoa(result))

}

func TestPartTwoSample(t *testing.T) {
	expected := 4
	testResult := four.PartTWO("test_input.txt")
	if testResult != expected {
		t.Fatalf("Part one failed was not %v -- instead it was %v\n", expected, testResult)
	} else {
		println("Sample for TestPartTwoSample ok")
		t.Log("SAMPLE OK")
	}

	result := four.PartTWO("input.txt")
	println("Result of input is: " + strconv.Itoa(result))
}
