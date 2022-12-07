package main

import (
	"bufio"
	"fmt"
	"log"
	. "main/pkg/debug"
	. "main/pkg/nodes"
	. "main/pkg/parser"
	"os"
)

const SIZE_MAX = 100_000

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	err, tree := Parse(scanner)

	if err != nil {
		if Debug() {
			LogIt("\n")
			LogIO.Flush()
		}
		log.Fatal(err)
	}

	criteria := func(n *Node) bool {
		s := n.GetSize()
		if Debug() {
			if s <= SIZE_MAX && n.IsDir() {
				LogIt(fmt.Sprintf("size=%d", s))
			}
		}
		return s <= SIZE_MAX && n.IsDir()
	}

	results := tree.WalkWithCriteria(criteria)

	sum := 0

	for _, r := range results {
		sum += r.GetSize()
	}
	if Debug() {
		LogIt("\n")
		LogIO.Flush()
	}

	fmt.Println(fmt.Sprintf("total_size=%d", sum))
}
