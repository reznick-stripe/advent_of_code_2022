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
	return opponentChoice.Loses
}

var youLose = func(opponentChoice *Choice) *Choice {
	return opponentChoice.Beats
}

var youTie = func(opponentChoice *Choice) *Choice {
	return opponentChoice.Ties
}

var rock = Choice{Score: 1}
var paper = Choice{Score: 2}
var scissors = Choice{Score: 3}

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

	if yourChoice.WinsAgainst(opponentChoice) {
		return yourChoice.Score + 6
	}

	if yourChoice.TiesAgainst(opponentChoice) {
		return yourChoice.Score + 3
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
		opponentChoice := opponentChoiceMap[opponentInput]
		yourChoice := yourChoiceMap[outcomeInput](opponentChoice)
		score += ScoreRound(opponentChoice, yourChoice)
	}

	fmt.Printf("highest score: %d\n", score)
}
