package five

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PartOne(input string) string {
	return execute(input, "ONE")
}

func PartTwo(input string) string {
	return execute(input, "TWO")
}

func execute(input string, part string) string {

	f, err := os.Open(input)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Do setup
	stacks, lineNr := setupStacks(lines)

	//Move crates
	switch part {
	case "ONE":
		moveCratesOne(lineNr, lines, stacks)
	case "TWO":
		moveCratesTwo(lineNr, lines, stacks)
	}

	//Output top crates
	topCrates := getTopCrates(stacks)

	outputStacks(stacks)

	return topCrates
}

func setupStacks(lines []string) ([]*list.List, int) {
	length := len(lines[0])
	fmt.Println(length)

	howManyStacks := (len(lines[0]) - len(lines[0])/4) / 3
	var stacks = make([]*list.List, howManyStacks)

	lineNr := 0
	for i := 0; i < howManyStacks; i++ {
		stacks[i] = list.New()
	}

	for i := 0; i < len(lines); i++ {
		if !strings.Contains(lines[i], "[") {
			lineNr = i + 2
			break
		}

		stacknr := 0
		for y := 1; y < len(lines[i])-1; y += 4 {

			value := lines[i][y]
			crateChar := regexp.MustCompile(`[A-Z]`)
			matches := crateChar.MatchString(string(value))
			if matches {
				stacks[stacknr].PushBack(value)
			}
			stacknr++
		}
	}
	return stacks, lineNr
}

func moveCratesTwo(lineNr int, lines []string, stacks []*list.List) {
	regexp := regexp.MustCompile(`\d+`)
	for i := lineNr; i < len(lines); i++ {

		matches := regexp.FindAllString(lines[i], -1)

		if len(matches) != 3 {
			panic("Not correct parse for matches")
		}

		amoutOfPops, _ := strconv.Atoi(matches[0])
		moveFrom, _ := strconv.Atoi(matches[1])
		moveTo, _ := strconv.Atoi(matches[2])

		tempList := list.New()
		for y := 0; y < amoutOfPops; y++ {
			crate := stacks[moveFrom-1].Front()
			tempList.PushFront(crate.Value.(uint8))
			stacks[moveFrom-1].Remove(crate)
		}

		for e := tempList.Front(); e != nil; e = e.Next() {
			stacks[moveTo-1].PushFront(e.Value.(uint8))
		}
	}
}

func moveCratesOne(lineNr int, lines []string, stacks []*list.List) {
	regexp := regexp.MustCompile(`\d+`)
	for i := lineNr; i < len(lines); i++ {

		matches := regexp.FindAllString(lines[i], -1)

		if len(matches) != 3 {
			panic("Not correct parse for matches")
		}

		amoutOfPops, _ := strconv.Atoi(matches[0])
		moveFrom, _ := strconv.Atoi(matches[1])
		moveTo, _ := strconv.Atoi(matches[2])

		for y := 0; y < amoutOfPops; y++ {
			crate := stacks[moveFrom-1].Front()
			stacks[moveFrom-1].Remove(crate)
			stacks[moveTo-1].PushFront(crate.Value.(uint8))
		}
	}
}

func getTopCrates(stacks []*list.List) string {
	topCrates := ""
	for i := 0; i < len(stacks); i++ {
		topCrates += string(stacks[i].Front().Value.(uint8))
	}
	return topCrates
}

func outputStacks(stacks []*list.List) {
	fmt.Println("STACK 1:")
	for e := stacks[0].Front(); e != nil; e = e.Next() {
		fmt.Println(string(e.Value.(uint8)))
	}
	fmt.Println("STACK 2:")
	for e := stacks[1].Front(); e != nil; e = e.Next() {
		fmt.Println(string(e.Value.(uint8)))
	}
	fmt.Println("STACK 3:")
	for e := stacks[2].Front(); e != nil; e = e.Next() {
		fmt.Println(string(e.Value.(uint8)))
	}
}

func initialStack(text string, stacks []*list.List) {
	replacer := strings.NewReplacer("[", "", "]", "")
	strippedBrackets := replacer.Replace(text)

	crates := strings.Fields(strippedBrackets)

	if len(crates) > len(stacks) {
		stacks = append(stacks, list.New())
		for i, value := range crates {
			stacks[i].PushFront(value)
		}
	}
}

func SolveOne(text string) byte {
	return 0
}

func SolveTwo(text string) byte {
	return 0
}
