package three

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//type M map[string]int

func PartOne(input string) int {

	f, err := os.Open(input)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	setList := []map[byte]int{}

	for scanner.Scan() {
		wholeBag := scanner.Text()
		compartment1 := wholeBag[0 : len(wholeBag)/2]
		compartment2 := wholeBag[len(wholeBag)/2:]
		println(compartment1)
		println(compartment2)

		srcMap := make(map[byte]int)
		destMap := make(map[byte]int)

		for _, value := range compartment1 {
			if _, ok := srcMap[byte(value)]; ok {
				srcMap[byte(value)]++
			} else {
				srcMap[byte(value)] = 1
			}
		}

		for _, value := range compartment2 {
			if _, ok := srcMap[byte(value)]; ok {
				if _, ok := destMap[byte(value)]; ok {
					destMap[byte(value)] = destMap[byte(value)] + 1
				}
				destMap[byte(value)] = 1
			}
		}

		setList = append(setList, destMap)
	}

	sumOfPriorities := 0
	for _, set := range setList {
		for key, element := range set {
			if element >= 1 {
				sumOfPriorities += getPriority(key) * element
			}
		}
	}

	return sumOfPriorities

}

func getPriority(val byte) int {
	if val < 91 {
		return int(val) - 38
	} else {
		return int(val) - 96
	}
}

func PartTWO(input string) int {

	f, err := os.Open(input)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	setList := []map[byte]int{}

	var firstElf string
	var secondElf string
	var thirdElf string

	var elfUnit = 1

	for scanner.Scan() {

		stuff := scanner.Text()

		switch elfUnit {
		case 1:
			firstElf = stuff
			elfUnit++

		case 2:
			secondElf = stuff
			elfUnit++
		case 3:
			thirdElf = stuff
			firstMap := setMap(firstElf)

			firstMerge := mergeMaps(secondElf, firstMap)
			secondMerge := mergeMaps(thirdElf, firstMerge)
			fmt.Println(secondMerge)

			setList = append(setList, secondMerge)
			elfUnit = 1
		}

	}

	sumOfPriorities := 0
	for _, set := range setList {
		for key, element := range set {
			if element >= 1 {
				sumOfPriorities += getPriority(key) * element
			}
		}
	}

	return sumOfPriorities

}

func mergeMaps(secondElf string, srcMap map[byte]int) map[byte]int {
	destMap := make(map[byte]int)
	for _, value := range secondElf {
		if _, ok := srcMap[byte(value)]; ok {
			if _, ok := destMap[byte(value)]; ok {
				destMap[byte(value)] = destMap[byte(value)] + 1
			}
			destMap[byte(value)] = 1
		}
	}
	return destMap
}

func setMap(firstElf string) map[byte]int {
	srcMap := make(map[byte]int)

	for _, value := range firstElf {
		if _, ok := srcMap[byte(value)]; ok {
			srcMap[byte(value)]++
		} else {
			srcMap[byte(value)] = 1
		}
	}
	return srcMap
}

/**
for i := 0; i < len(compartment1); i++ {
	if _, ok := compartment1Set[]; ok {
		fmt.Println("element exists")
		setList[index][compartment1[i]] = setList[index][compartment1[i]] + 1
	} else {
		fmt.Println("element not found")
		setList[index][compartment1[i]] = 1
	}
}
*/
