package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Choice struct {
	Beats *Choice
	Ties  *Choice
	Loses *Choice
	Score int
	Name  string
}

var logIO = bufio.NewWriter(os.Stdout)

func logIt(s string) {
	fmt.Fprint(logIO, fmt.Sprintf("%s ", s))
}

func debug() bool {
	return os.Getenv("DEBUG") == "true"
}

func (c *Choice) WinsAgainst(o *Choice) bool {
	return c.Beats == o
}

func (c *Choice) TiesAgainst(o *Choice) bool {
	return c.Ties == o
}

func (c *Choice) LosesAgainst(o *Choice) bool {
	return c.Loses == o
}

var youWin = func(opponentChoice *Choice) *Choice {
	if debug() {
		logIt("desired_outcome=win")
	}
	return opponentChoice.Loses
}

var youLose = func(opponentChoice *Choice) *Choice {
	if debug() {
		logIt("desired_outcome=lose")
	}
	return opponentChoice.Beats
}

var youTie = func(opponentChoice *Choice) *Choice {
	if debug() {
		logIt("desired_outcome=tie")
	}
	return opponentChoice.Ties
}

var rock = Choice{Name: "rock", Score: 1}
var paper = Choice{Name: "paper", Score: 2}
var scissors = Choice{Name: "scissors", Score: 3}

var opponentChoiceMap = map[string]*Choice{
	"A": &rock,
	"B": &paper,
	"C": &scissors,
}

var yourChoiceMap = map[string](func(*Choice) *Choice){
	"X": youLose,
	"Y": youTie,
	"Z": youWin,
}

func ScoreRound(opponentChoice *Choice, yourChoice *Choice) int {
	if debug() {
		logIt(fmt.Sprintf("opponent_choice=%s your_choice=%s", opponentChoice.Name, yourChoice.Name))
		logIt(fmt.Sprintf("your_choice_score=%d", yourChoice.Score))
	}

	if yourChoice.WinsAgainst(opponentChoice) {
		if debug() {
			logIt("outcome=win outcome_score=6")
		}
		return yourChoice.Score + 6
	}

	if yourChoice.TiesAgainst(opponentChoice) {
		if debug() {
			logIt("outcome=tie outcome_score=3")
		}
		return yourChoice.Score + 3
	}

	if debug() {
		logIt("outcome=loss outcome_score=0")
	}
	return yourChoice.Score
}

func unpack(s string) (string, string) {
	x := strings.Split(s, " ")
	return x[0], x[1]
}

func main() {
	rock.Beats = &scissors
	rock.Ties = &rock
	rock.Loses = &paper
	scissors.Beats = &paper
	scissors.Ties = &scissors
	scissors.Loses = &rock
	paper.Beats = &rock
	paper.Ties = &paper
	paper.Loses = &scissors

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0

	for scanner.Scan() {
		str := scanner.Text()
		opponentInput, outcomeInput := unpack(str)
		if debug() {
			logIt(fmt.Sprintf("opponent_input=%s outcome_input=%s", opponentInput, outcomeInput))
		}
		opponentChoice := opponentChoiceMap[opponentInput]
		yourChoice := yourChoiceMap[outcomeInput](opponentChoice)
		score += ScoreRound(opponentChoice, yourChoice)
		if debug() {
			logIt(fmt.Sprintf("total_score=%d", score))
			logIt("\n")
			logIO.Flush()
		}
	}

	fmt.Printf("highest score: %d\n", score)
}
