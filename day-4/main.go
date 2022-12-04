package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func debug() bool {
	return os.Getenv("DEBUG") == "true"
}

var logIO = bufio.NewWriter(os.Stdout)

func logIt(s string) {
	fmt.Fprint(logIO, fmt.Sprintf("%s ", s))
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	}
}
