package instructions

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Instruction struct {
	Count int
	From  int
	To    int
}

func (i Instruction) String() string {
	return fmt.Sprintf("move %d from %d to %d", i.Count, i.From, i.To)
}

func NewInstructionFromInput(s string) (*Instruction, error) {
	r, err := regexp.Compile(`^ move (?P<count>\d+) from (?P<from>\d+) to (?P<to>\d+)`)

	if err != nil {
		return nil, err
	}

	m := r.FindStringSubmatch(s)

	data := make(map[string]int)

	if len(m) != 4 {
		return nil, errors.New(fmt.Sprintf("bad parse: %s", s))
	}

	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			n, err := strconv.Atoi(m[i])

			if err != nil {
				return nil, err
			}
			data[name] = n
		}
	}

	return &Instruction{Count: data["count"], From: data["from"], To: data["to"]}, nil
}
