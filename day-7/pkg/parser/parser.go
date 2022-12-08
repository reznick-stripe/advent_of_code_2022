package parser

import (
	"bufio"
	command "main/pkg/commands"
	. "main/pkg/tree"
	"strings"
)

func Parse(scanner *bufio.Scanner) (error, *Tree) {
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	tree := NewTree()

	var dataBuff []string
	cmdString := ""

	for i, l := range lines {
		//build the string
		if strings.ContainsRune(l, '$') {
			cmdString = l
		} else {
			dataBuff = append(dataBuff, l)
		}

		// don't lookahead
		if i+1 == len(lines) {
			err := tree.Exec(cmdString, command.WithData(dataBuff))
			if err != nil {
				return err, nil
			}
			dataBuff = []string{}
			continue
		}

		//lookahead
		if strings.ContainsRune(lines[i+1], '$') {
			err := tree.Exec(cmdString, command.WithData(dataBuff))
			if err != nil {
				return err, nil
			}
			dataBuff = []string{}
		}
	}

	tree.ReturnToRoot()
	return nil, &tree
}
