package util

import (
	"bufio"
	"log"
	"os"
)

func Read(filepath string) *bufio.Scanner {
	f, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	return bufio.NewScanner(f)
}
