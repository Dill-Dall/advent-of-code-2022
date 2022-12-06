package two

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Round struct {
	OPPONENT rune
	YOU      rune
}

/*
	OPPONENT      |YOU
	A|ROCK		X|ROCK = 1
	B|PAPER		Y|PAPER = 2
	C|SCISSORS 	Z|SCISSORS = 3

	WIN = 6 + selection(1|2|3)
	DRAW = 3 + selection(1|2|3)
	LOSE = 0
*/

func PartOne(input string) int {

	f, err := os.Open(input)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var points = 0
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())

		round := Round{

			OPPONENT: []rune(values[0])[0],
			YOU:      []rune(values[1])[0],
		}

		resultPoint, typePoint := determineRoundPoints(round)
		roundPoints := typePoint + resultPoint
		fmt.Printf("Round %c VS %c gives %v points to YOU\n", round.OPPONENT, round.YOU, roundPoints)
		points += roundPoints
	}

	fmt.Printf("Points = %v\n", points)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return points
}

func PartTWO(input string) int {

	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var points = 0
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())

		round := Round{ // b == Student{"Bob", 0}

			OPPONENT: []rune(values[0])[0],
			YOU:      []rune(values[1])[0],
		}
		//rounds = append(rounds, round)

		resultPoint, typePoint := determineRoundPointsV2(round)
		roundPoints := typePoint + resultPoint
		fmt.Printf("Round %c VS %c gives %v points to YOU\n", round.OPPONENT, round.YOU, roundPoints)
		points += roundPoints
	}

	fmt.Printf("Points = %v\n", points)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return points
}

func determineRoundPoints(round Round) (int, int) {
	var resultPoint = 0
	var typePoint = 0

	switch round.YOU {
	case 'X':
		typePoint = 1
		switch round.OPPONENT {
		case 'A':
			resultPoint = 3
		case 'B':
			resultPoint = 0
		case 'C':
			resultPoint = 6
		default:
			log.Fatal("Invalid OPPONENT")
		}
	//ROCK
	case 'Y':
		typePoint = 2
		switch round.OPPONENT {
		case 'A':
			resultPoint = 6
		case 'B':
			resultPoint = 3
		case 'C':
			resultPoint = 0
		default:
			log.Fatal("Invalid OPPONENT")
		}
	case 'Z':
		typePoint = 3
		switch round.OPPONENT {
		case 'A':
			resultPoint = 0
		case 'B':
			resultPoint = 6
		case 'C':
			resultPoint = 3
		default:
			log.Fatal("Invalid OPPONENT")
		}
	default:
		log.Fatal("Invalid YOU")
	}
	return resultPoint, typePoint
}

func determineRoundPointsV2(round Round) (int, int) {
	var resultPoint = 0
	var typePoint = 0

	switch round.YOU {
	//loose
	case 'X':
		resultPoint = 0
		switch round.OPPONENT {
		case 'A': //  -> C
			typePoint = 3
		case 'B': // ->  A
			typePoint = 1
		case 'C': // -> B
			typePoint = 2
		default:
			log.Fatal("Invalid OPPONENT")
		}
	//draw
	case 'Y':
		resultPoint = 3
		switch round.OPPONENT {
		case 'A': // -> A
			typePoint = 1
		case 'B': // -> B
			typePoint = 2
		case 'C': // -> C
			typePoint = 3
		default:
			log.Fatal("Invalid OPPONENT")
		}
	//win
	case 'Z':
		resultPoint = 6
		switch round.OPPONENT {
		case 'A': // -> B
			typePoint = 2
		case 'B': // -> C
			typePoint = 3
		case 'C': // -> A
			typePoint = 1
		default:
			log.Fatal("Invalid OPPONENT")
		}
	default:
		log.Fatal("Invalid YOU")
	}
	return resultPoint, typePoint
}
