package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var priorityMap = make(map[rune]int)

// abcdef
// 012345
// len = 6
// 6/2 = 3

func getRucksacks(s string) ([]rune, []rune) {
	r := []rune(s)
	return r[0 : len(r)/2], r[len(r)/2:]
}

func intersection(left, right []rune) rune {
	h := make(map[rune]bool)
	for _, e := range left {
		h[e] = true
	}

	for _, e := range right {
		if h[e] {
			return e
		}
	}

	panic("no common rune")
}

func main() {
	for i, r := range alphabet {
		priorityMap[r] = i + 1
	}

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		left, right := getRucksacks(scanner.Text())
		sum += priorityMap[intersection(left, right)]
	}
	fmt.Printf("\npriority_sum=%d\n", sum)
}
