package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"main/pkg/assignments"
	"main/pkg/debug"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
		text := scanner.Text()
		first, second, err := assignments.AssignmentsFromLine(text)
		if err != nil {
			if debug.Debug() {
				debug.LogIt(fmt.Sprintf("input_line=%d error='%v' text='%s'", lines, err, text))
				debug.LogIO.Flush()
			}
			panic(err)
		}

		if first.FullyContains(second) || second.FullyContains(first) {
			count++
		}
	}

	fmt.Printf("\nfully_contains_count=%d\n", count)
}
