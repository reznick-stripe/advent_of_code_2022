package parser

import (
	"bufio"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	input := `30373
25512
65332
33549
35390
`

	expected := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	scanner := bufio.NewScanner(strings.NewReader(input))

	output, err := Parse(scanner)

	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	for i, row := range *output {
		for j, col := range row {
			e := expected[i][j]
			if e != col {
				t.Errorf("expected %d at (%d,%d) but got %d", e, i, j, col)
			}
		}
	}
}
