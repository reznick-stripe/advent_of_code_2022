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
}

func (c *Choice) WinsAgainst(o *Choice) bool {
	return c.Beats == o
}

func (c *Choice) TiesAgainst(o *Choice) bool {
	return c == o
}

var rock = Choice{}
var paper = Choice{}
var scissors = Choice{}

var opponentChoiceMap = map[string]*Choice{
	"A": &rock,
	"B": &paper,
	"C": &scissors,
}

var yourChoiceMap = map[string]*Choice{
	"X": &rock,
	"Y": &paper,
	"Z": &scissors,
}

func ScoreRound(opponentChoice *Choice, yourChoice *Choice, scoreMap map[Choice]int) int {
	choiceScore := scoreMap[*yourChoice]

	outcomeScore := 0

	if yourChoice.WinsAgainst(opponentChoice) {
		outcomeScore += 6
	}

	if yourChoice.TiesAgainst(opponentChoice) {
		outcomeScore += 3
	}

	return choiceScore + outcomeScore
}

func unpack(s string) (string, string) {
	x := strings.Split(s, " ")
	return x[0], x[1]
}

func main() {
	rock.Beats = &scissors
	scissors.Beats = &paper
	paper.Beats = &rock

	scoreMap := make(map[Choice]int)

	scoreMap[rock] = 1
	scoreMap[paper] = 2
	scoreMap[scissors] = 3

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0

	for scanner.Scan() {
		str := scanner.Text()
		opponentInput, yourInput := unpack(str)
		opponentChoice := opponentChoiceMap[opponentInput]
		yourChoice := yourChoiceMap[yourInput]
		score += ScoreRound(opponentChoice, yourChoice, scoreMap)
	}

	fmt.Printf("highest score: %d\n", score)
}
