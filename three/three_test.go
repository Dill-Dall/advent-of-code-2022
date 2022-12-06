package three_test

import (
	"2022/three"
	"strconv"
	"testing"
)

func TestPartOne(t *testing.T) {
	expected := 157
	test_result := three.PartOne("test_input.txt")
	if test_result != 157 {
		t.Fatalf("Part one failed was not %v -- instead it was %v\n", expected, test_result)
	} else {
		println("Sample ok")
		t.Log("SAMPLE OK")
	}
	result := three.PartOne("input.txt")
	println("Result is: " + strconv.Itoa(result))

}

func PartOneTest(t *testing.T) {

}

func TestPartTwoSample(t *testing.T) {
	expected := 70
	testResult := three.PartTWO("test_input.txt")
	if testResult != expected {
		t.Fatalf("Part one failed was not %v -- instead it was %v\n", expected, testResult)
	} else {
		t.Log("SAMPLE OK")
	}

	result := three.PartTWO("input.txt")
	println("Result of input is: " + strconv.Itoa(result))
}
