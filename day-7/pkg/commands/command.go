package command

import (
	"errors"
	"fmt"
	"regexp"
)

type CommandType int

const (
	CD CommandType = iota
	LS
)

func (c CommandType) String() string {
	switch c {
	case CD:
		return "cd"
	case LS:
		return "ls"
	default:
		return "unknown"
	}
}

type Command struct {
	Type   CommandType
	Target string
	Data   []string
}

type Option func(c *Command)

func WithData(data []string) Option {
	return func(c *Command) {
		c.Data = data
	}
}

func CommandFromPrompt(s string) (*Command, error) {
	r, err := regexp.Compile(`^\$ (?P<cmd>\w+)[ ]?(?P<target>[a-zA-Z0-9-_./]+)?`)

	if err != nil {
		return nil, err
	}

	m := r.FindStringSubmatch(s)

	data := make(map[string]string)

	if len(m) > 0 {
		for i, name := range r.SubexpNames() {
			if i != 0 && name != "" {
				data[name] = m[i]
			}
		}
	}

	switch data["cmd"] {
	case "cd":
		if data["target"] == "" {
			return nil, errors.New(fmt.Sprintf("no target from %s", s))
		}

		return &Command{Type: CD, Target: data["target"]}, nil
	case "ls":
		return &Command{Type: LS, Target: ""}, nil
	default:
		return nil, errors.New(fmt.Sprintf("no parsable command from %s", s))
	}
}
