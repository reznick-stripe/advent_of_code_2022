package parser

import (
	"bufio"
	. "main/pkg/treemap"
)

func Parse(scanner *bufio.Scanner) (*TreeMap, error) {
	var grid TreeMap

	for scanner.Scan() {
		var row []int
		line := scanner.Text()
		for _, c := range line {
			row = append(row, int(c-'0'))
		}
		grid = append(grid, row)
	}
	return &grid, nil
}
