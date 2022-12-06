package five

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
	fmt.Printf("%v\n\n", msg)
}

func Hello(name string) (string, error) {
	return "Hello " + name, nil
}

func TestOne(t *testing.T) {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var arrayOfCalories = []int{0}
	var index = 0
	var highest = index

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		arrayOfCalories[index] += value
		if err != nil {
			if arrayOfCalories[index] > arrayOfCalories[highest] {
				highest = index
			}
			index += 1
			arrayOfCalories = append(arrayOfCalories, 0)
		}
	}

	fmt.Printf("Highest is index \"%v\" with value of %v\n", highest, arrayOfCalories[highest])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sortedArray := BubbleSort(arrayOfCalories)

	fmt.Printf("The top %v elfs have %v calories total", 3, SumOfHighest(sortedArray, 3))

}

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func SumOfHighest(array []int, nrToSum int) int {
	highestAmongst := 0
	for i := 0; i < nrToSum; i++ {
		highestAmongst += array[len(array)-i-1]
	}
	return highestAmongst
}
