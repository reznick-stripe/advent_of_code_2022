package parser

import (
	"bufio"
	. "main/pkg/treemap"
)

func Parse(scanner *bufio.Scanner) (*TreeMap, error) {
	var grid TreeMap

	for scanner.Scan() {
		grid.RowCount++
		var row []int
		line := scanner.Text()
		for _, c := range line {
			row = append(row, int(c-'0'))
		}

		if grid.RowCount == 1 {
			grid.ColCount = len(row)
		}
		grid.Data = append(grid.Data, row)
	}
	return &grid, nil
}
