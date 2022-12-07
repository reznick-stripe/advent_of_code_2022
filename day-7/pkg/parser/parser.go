package parser

import (
	"bufio"
	"fmt"
	command "main/pkg/commands"
	. "main/pkg/debug"
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
		if Debug() {
			LogIt(fmt.Sprintf("line=%d", i))
		}
		//build the string
		if strings.ContainsRune(l, '$') {
			cmdString = l
			if Debug() {
				LogIt(fmt.Sprintf("line_type=cmd cmd='%s' pwd=%s", cmdString, tree.Pwd.GetFullPath()))
			}
		} else {
			dataBuff = append(dataBuff, l)
			if Debug() {
				LogIt(fmt.Sprintf("line_type=data cmd='%s' data_buff_len=%d entry='%s' pwd=%s", cmdString, len(dataBuff), l, tree.Pwd.GetFullPath()))
			}
		}

		// don't lookahead
		if i+1 == len(lines) {
			if Debug() {
				LogIt("next_eof=true")
			}
			err := tree.Exec(cmdString, command.WithData(dataBuff))
			if err != nil {
				return err, nil
			}
			if Debug() {
				LogIt("\n")
			}
			dataBuff = []string{}
			continue
		}

		if Debug() {
			LogIt("next_eof=false")
		}

		//lookahead
		if strings.ContainsRune(lines[i+1], '$') {
			err := tree.Exec(cmdString, command.WithData(dataBuff))
			if err != nil {
				return err, nil
			}
			dataBuff = []string{}
		}
		if Debug() {
			LogIt("\n")
		}
	}

	tree.ReturnToRoot()
	return nil, &tree
}
