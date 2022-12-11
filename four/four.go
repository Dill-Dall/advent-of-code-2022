package four

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

//type M map[string]int

func PartOne(input string) int {

	f, err := os.Open(input)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	nrOfRangeContains := 0

	for scanner.Scan() {
		text := scanner.Text()

		if doesOneRangeContainAnother(text) {
			nrOfRangeContains++
		}

	}

	return nrOfRangeContains
}

func PartTWO(input string) int {

	f, err := os.Open(input)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	nrOfRangeContains := 0

	for scanner.Scan() {
		text := scanner.Text()

		if anyOverLap(text) {
			nrOfRangeContains++
		}

	}

	return nrOfRangeContains
}

type ResponsibilityRange struct {
	elf1 [2]int
	elf2 [2]int
}

func NewResponsibilityRange(range1 string, range2 string) *ResponsibilityRange {

	array1 := strings.Split(range1, "-")
	array2 := strings.Split(range2, "-")

	workPair := new(ResponsibilityRange)
	workPair.elf1[0], _ = strconv.Atoi(array1[0])
	workPair.elf1[1], _ = strconv.Atoi(array1[1])
	workPair.elf2[0], _ = strconv.Atoi(array2[0])
	workPair.elf2[1], _ = strconv.Atoi(array2[1])
	return workPair
}

func doesOneRangeContainAnother(text string) bool {
	ranges := strings.Split(text, ",")
	workPair := NewResponsibilityRange(ranges[0], ranges[1])

	//elf1 encompasses elf 2
	if ((workPair.elf1[0] <= workPair.elf2[0]) && (workPair.elf1[1] >= workPair.elf2[1])) ||
		//elf2 encompasses elf 1
		((workPair.elf2[0] <= workPair.elf1[0]) && (workPair.elf2[1] >= workPair.elf1[1])) {
		return true
	}
	return false
}

func anyOverLap(text string) bool {
	ranges := strings.Split(text, ",")
	workPair := NewResponsibilityRange(ranges[0], ranges[1])

	// x1 <= C <= x2 && y1 <= C <= y2   - https://stackoverflow.com/a/3269471

	if workPair.elf1[0] <= workPair.elf2[1] && workPair.elf2[0] <= workPair.elf1[1] {
		return true
	}
	return false
}
