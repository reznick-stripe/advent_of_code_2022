package treemap

import (
	"fmt"
)

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

func (d Direction) String() string {
	return [...]string{"north", "south", "east", "west"}[d]
}

func GetScoreForDirection(tree *TreeMap, row int, col int, direction Direction) int {
	count := 0
	startingHeight := tree.Data[row][col]

	switch direction {
	case North:
		for i := row - 1; i >= 0; i-- {
			if tree.Data[i][col] < startingHeight {
				count++
			} else {
				count++
				break
			}
		}
	case South:
		for i := row + 1; i < tree.RowCount; i++ {
			if tree.Data[i][col] < startingHeight {
				count++
			} else {
				count++
				break
			}
		}
	case East:
		for i := col + 1; i < tree.ColCount; i++ {
			if tree.Data[row][i] < startingHeight {
				count++
			} else {
				count++
				break
			}
		}
	case West:
		for i := col - 1; i >= 0; i-- {
			if tree.Data[row][i] < startingHeight {
				count++
			} else {
				count++
				break
			}
		}
	}

	return count
}

func GetScoreForLocation(tree *TreeMap, row int, col int) int {
	compass := []Direction{North, South, East, West}

	product := 1

	for _, d := range compass {
		product *= GetScoreForDirection(tree, row, col, d)
	}

	return product
}

type TreeMap struct {
	Data       [][]int
	RowCount   int
	ColCount   int
	VisibleMap [][]rune
}

func (t *TreeMap) ScanARow(row int, vantage string) error {
	biggestYet := 0
	r := t.Data[row]
	if vantage == "east" {
		// facing east from the west
		for i, n := range r {
			if n == -1 {
				continue
			}
			if i == 0 || n > biggestYet {
				t.VisibleMap[row][i] = 'y'
				biggestYet = n
				continue
			}
		}
	} else if vantage == "west" {
		// facing west from the east
		endIndex := t.ColCount - 1
		for i := endIndex; i >= 0; i-- {
			n := r[i]
			if n == -1 {
				continue
			}
			if i == endIndex || n > biggestYet {
				t.VisibleMap[row][i] = 'y'
				biggestYet = n
				continue
			}
		}
	} else {
		return fmt.Errorf("bad vantage: %s", vantage)
	}

	return nil
}

func (t *TreeMap) ScanAColumn(col int, vantage string) error {
	biggestYet := 0

	if vantage == "south" {
		// facing south from the north
		for i, row := range t.Data {
			n := row[col]
			if n == -1 {
				continue
			}
			if i == 0 || n > biggestYet {
				t.VisibleMap[i][col] = 'y'
				biggestYet = n
				continue
			}
		}
	} else if vantage == "north" {
		// facing north from the south
		endIndex := t.RowCount - 1
		for i := endIndex; i >= 0; i-- {
			n := t.Data[i][col]
			if n == -1 {
				continue
			}
			if i == endIndex || n > biggestYet {
				t.VisibleMap[i][col] = 'y'
				biggestYet = n
				continue
			}
		}
	} else {
		return fmt.Errorf("bad vantage: %s", vantage)
	}

	return nil
}

func (t *TreeMap) scan() {
	for y := range t.Data {
		for _, direction := range []string{"east", "west"} {
			t.ScanARow(y, direction)
		}
	}

	for x := 0; x < t.ColCount-1; x++ {
		for _, direction := range []string{"north", "south"} {
			t.ScanAColumn(x, direction)
		}
	}
}

func (t TreeMap) Count() int {
	t.scan()
	sum := 0

	for y, row := range t.Data {
		for x := range row {
			if t.VisibleMap[y][x] == 'y' {
				sum++
			}
		}
	}

	return sum
}
