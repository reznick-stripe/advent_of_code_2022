package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Rucksack []rune

type Group struct {
	Rucksacks [3]Rucksack
}

func debug() bool {
	return os.Getenv("DEBUG") == "true"
}

var logIO = bufio.NewWriter(os.Stdout)

func logIt(s string) {
	fmt.Fprint(logIO, fmt.Sprintf("%s ", s))
}

var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var priorityMap = make(map[rune]int)

func intersection(rucksacks [3]Rucksack) rune {
	h := make(map[rune]uint8)
	for _, e := range rucksacks[0] {
		h[e] |= 0b100
	}

	for _, e := range rucksacks[1] {
		h[e] |= 0b010
	}

	for _, e := range rucksacks[2] {
		h[e] |= 0b001
	}

	for r, b := range h {
		if b == 0b111 {
			if debug() {
				logIt(fmt.Sprintf("rune=%c binary=%03b", r, b))
			}
			return r
		}
	}

	return rune('0')
}

func main() {
	for i, r := range alphabet {
		priorityMap[r] = i + 1
	}

	var groups []Group

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	index := 0
	sum := 0
	for scanner.Scan() {
		if debug() {
			logIt(fmt.Sprintf("line=%d", index))
		}
		m := index % 3
		if m == 0 {
			groups = append(groups, Group{Rucksacks: [3]Rucksack{[]rune(scanner.Text())}})
			if debug() {
				logIt(fmt.Sprintf("group=%d", len(groups)-1))
				logIt("\n")
				logIO.Flush()
			}
		} else if m == 1 {
			if debug() {
				logIt(fmt.Sprintf("group=%d", len(groups)-1))
				logIt("\n")
				logIO.Flush()
			}
			groups[len(groups)-1].Rucksacks[m] = []rune(scanner.Text())
		} else if m == 2 {
			groups[len(groups)-1].Rucksacks[m] = []rune(scanner.Text())
			sum += priorityMap[intersection(groups[len(groups)-1].Rucksacks)]
			if debug() {
				logIt(fmt.Sprintf("group=%d", len(groups)-1))
				logIt(fmt.Sprintf("priority_sum=%d", sum))
				logIt("\n")
				logIO.Flush()
			}
		} else {
			panic("math!?")
		}
		index++
	}

	fmt.Printf("\npriority_sum=%d\n", sum)
}
