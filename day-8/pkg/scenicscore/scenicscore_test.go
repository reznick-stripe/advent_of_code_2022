package scenicscore

import (
	. "main/pkg/treemap"
	"testing"
)

func TestGetScoreForDirection(t *testing.T) {
	m := TreeMap{
		Data: [][]int{
			{3, 0, 3, 7, 3},
			{2, 5, 5, 1, 2},
			{6, 5, 3, 3, 2},
			{3, 3, 5, 4, 9},
			{3, 5, 3, 9, 0},
		},
		RowCount: 5,
		ColCount: 5,
		VisibleMap: [][]rune{
			{'n', 'n', 'n', 'n', 'n'},
			{'n', 'n', 'n', 'n', 'n'},
			{'n', 'n', 'n', 'n', 'n'},
			{'n', 'n', 'n', 'n', 'n'},
			{'n', 'n', 'n', 'n', 'n'},
		},
	}

	row := 3
	col := 2

	testCases := []struct {
		Expected int
		Actual   int
		Dir      Direction
	}{
		{2, GetScoreForDirection(&m, row, col, North), North},
		{1, GetScoreForDirection(&m, row, col, South), South},
		{2, GetScoreForDirection(&m, row, col, East), East},
		{2, GetScoreForDirection(&m, row, col, West), West},
	}

	for _, c := range testCases {
		if c.Expected != c.Actual {
			t.Errorf("expected %d but got %d for (%d,%d) and %s", c.Expected, c.Actual, row, col, c.Dir)
		}
	}
}

func TestGetScoreForLocation(t *testing.T) {
	m := TreeMap{
		Data: [][]int{
			{3, 0, 3, 7, 3},
			{2, 5, 5, 1, 2},
			{6, 5, 3, 3, 2},
			{3, 3, 5, 4, 9},
			{3, 5, 3, 9, 0},
		},
		RowCount: 5,
		ColCount: 5,
		VisibleMap: [][]rune{
			{'n', 'n', 'n', 'n', 'n'},
			{'n', 'n', 'n', 'n', 'n'},
			{'n', 'n', 'n', 'n', 'n'},
			{'n', 'n', 'n', 'n', 'n'},
			{'n', 'n', 'n', 'n', 'n'},
		},
	}

	row := 3
	col := 2

	expected := 8
	actual := GetScoreForLocation(&m, row, col)

	if actual != expected {
		t.Errorf("expected %d but got %d for (%d,%d)", expected, actual, row, col)
	}
}
