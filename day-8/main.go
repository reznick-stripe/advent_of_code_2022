package main

import (
	"bufio"
	"fmt"
	"log"
	. "main/pkg/parser"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	treeMap, err := Parse(scanner)
	if err != nil {
		log.Fatal(err)
	}

	sum := treeMap.Count()

	fmt.Println(fmt.Sprintf("sum=%d", sum))

	biggestScenicScore := treeMap.GetBiggestScenicScore()

	fmt.Println(fmt.Sprintf("biggest_scenic_score=%d", biggestScenicScore))
}
