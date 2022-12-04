package day4

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Assignment struct {
	Lower int
	Upper int
}

func (a *Assignment) FullyContains(other *Assignment) bool {
	return a.Lower <= other.Lower && a.Upper >= other.Upper
}

func NewAssignmentFromString(s string) (*Assignment, error) {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return nil, errors.New(fmt.Sprintf("error parsing: %s", s))
	}

	lower, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	upper, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	if lower == upper {
		return nil, errors.New("lower and upper match")
	}

	if lower > upper {
		return nil, errors.New("lower is bigger than upper")
	}

	return &Assignment{Lower: lower, Upper: upper}, nil
}

func AssignmentsFromLine(s string) (*Assignment, *Assignment, error) {
	parts := strings.Split(s, ",")

	if len(parts) != 2 {
		return nil, nil, errors.New(fmt.Sprintf("error parsing line: %s", s))
	}

	first, err := NewAssignmentFromString(parts[0])
	if err != nil {
		return first, nil, err
	}

	second, err := NewAssignmentFromString(parts[1])
	if err != nil {
		return first, second, err
	}

	return first, second, nil
}

func debug() bool {
	return os.Getenv("DEBUG") == "true"
}

var logIO = bufio.NewWriter(os.Stdout)

func logIt(s string) {
	fmt.Fprint(logIO, fmt.Sprintf("%s ", s))
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	}
}
