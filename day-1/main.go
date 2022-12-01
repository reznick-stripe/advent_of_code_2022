package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	Number   int
	Calories int
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	counts := make(map[int]int)

	elf_counter := 1

	biggest_yet := 0

	var elves []Elf

	for scanner.Scan() {
		str := scanner.Text()
		if len(str) == 0 {
			if counts[elf_counter] >= biggest_yet {
				biggest_yet = counts[elf_counter]
			}
			e := Elf{Number: elf_counter, Calories: counts[elf_counter]}
			elves = append(elves, e)
			elf_counter += 1
		} else {
			i, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			counts[elf_counter] += i
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].Calories > elves[j].Calories
	})

	top_three_count := 0

	for _, elf := range elves[0:3] {
		top_three_count += elf.Calories
	}

	fmt.Printf("Most Calories %d\n", biggest_yet)
	fmt.Printf("Top Three Total %d\n", top_three_count)
}
