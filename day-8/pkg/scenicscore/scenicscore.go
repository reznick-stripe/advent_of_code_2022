package scenicscore

import (
	. "main/pkg/treemap"
)

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

func (d Direction) String() string {
	switch d {
	case North:
		return "north"
	case South:
		return "south"
	case East:
		return "east"
	case West:
		return "west"
	default:
		return ""
	}
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
