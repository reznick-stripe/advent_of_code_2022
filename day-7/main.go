package main

import (
	"bufio"
	"fmt"
	"log"
	. "main/pkg/debug"
	. "main/pkg/nodes"
	. "main/pkg/parser"
	"os"
	"sort"
)

const SIZE_MAX = 40_000_000

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

	remainingSize := tree.Root.GetSize() - SIZE_MAX

	if Debug() {
		LogIt(fmt.Sprintf("to_remove=%d", remainingSize))
		LogIt("\n")
	}

	criteria := func(n *Node) bool {
		return n.GetSize() >= remainingSize && n.IsDir()
	}

	results := tree.WalkWithCriteria(criteria)

	sort.Slice(results, func(i, j int) bool {
		return results[i].GetSize() < results[j].GetSize()
	})

	if Debug() {
		LogIt(fmt.Sprintf("lowest=%d highest=%d", results[0].GetSize(), results[len(results)-1].GetSize()))
		LogIt("\n")
		LogIO.Flush()
	}

	fmt.Println(fmt.Sprintf("total_size=%d", results[0].GetSize()))
}
