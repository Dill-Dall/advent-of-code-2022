package five_test

import (
	"2022/five"
	"testing"
)

func TestPartOne(t *testing.T) {
	expected := "CMZ"
	test_result := five.PartOne("test_input.txt")
	if test_result != expected {
		t.Fatalf("Part one failed was not %v -- instead it was %v\n", expected, test_result)
	} else {
		println("Sample ok")
		t.Log("SAMPLE OK")
	}
	result := five.PartOne("input.txt")
	println("Result is: " + result)

}

func TestPartTwoSample(t *testing.T) {
	expected := "MCD"
	testResult := five.PartTwo("test_input.txt")
	if testResult != expected {
		t.Fatalf("Part one failed was not %v -- instead it was %v\n", expected, testResult)
	} else {
		println("Sample for TestPartTwoSample ok")
		t.Log("SAMPLE OK")
	}

	result := five.PartTwo("input.txt")
	println("Result of input is: " + result)
}
