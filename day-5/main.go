package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	. "main/pkg/boards"
	. "main/pkg/crates"
	. "main/pkg/debug"
	. "main/pkg/instructions"
)

// [F]         [L]     [M]
// [T]     [H] [V] [G] [V]
// [N]     [T] [D] [R] [N]     [D]
// [Z]     [B] [C] [P] [B] [R] [Z]
// [M]     [J] [N] [M] [F] [M] [V] [H]
// [G] [J] [L] [J] [S] [C] [G] [M] [F]
// [H] [W] [V] [P] [W] [H] [H] [N] [N]
// [J] [V] [G] [B] [F] [G] [D] [H] [G]
//  1   2   3   4   5   6   7   8   9

func main() {
	board := Board{Crate{}, Crate{'J', 'H', 'G', 'M', 'Z', 'N', 'T', 'F'}, Crate{'V', 'W', 'J'}, Crate{'G', 'V', 'L', 'J', 'B', 'T', 'H'}, Crate{'B', 'P', 'J', 'N', 'C', 'D', 'V', 'L'}, Crate{'F', 'W', 'S', 'M', 'P', 'R', 'G'}, Crate{'G', 'H', 'C', 'F', 'B', 'N', 'V', 'M'}, Crate{'D', 'H', 'G', 'M', 'R'}, Crate{'H', 'N', 'M', 'V', 'Z', 'D'}, Crate{'G', 'N', 'F', 'H'}}

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := 0
	for scanner.Scan() {
		lines++
		text := scanner.Text()

		instruction, err := NewInstructionFromInput(text)
		if err != nil {
			if Debug() {
				LogIt(fmt.Sprintf("input_line=%d error='%v' input='%s'", lines, err, text))
				LogIO.Flush()
			}
			log.Fatal(err)
		}

		err = board.Move(instruction)
		if err != nil {
			if Debug() {
				LogIt(fmt.Sprintf("input_line=%d error='%v' instruction='%s'", lines, err, instruction))
				LogIO.Flush()
			}
			log.Fatal(err)
		}

		if Debug() {
			LogIt(fmt.Sprintf("input_line=%d instruction=%s top=%s\n", lines, instruction, board.Top()))
			LogIO.Flush()
		}
	}

	fmt.Printf("\noutput=%s\n", board.Top())
}
