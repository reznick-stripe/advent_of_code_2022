package treemap

type TreeMap [][]int

func (t TreeMap) CountForARow(row int, vantage string) int {
	visibleCount := 0
	biggestYet := 0
	r := t[row]
	if vantage == "east" {
		// facing east from the west
		for i, n := range r {
			if i == 0 || n > biggestYet {
				visibleCount++
				biggestYet = n
				continue
			}
		}
	} else if vantage == "west" {
		// facing west from the east
		endIndex := len(r) - 1
		for i := endIndex; i >= 0; i-- {
			n := r[i]
			if i == endIndex || n > biggestYet {
				visibleCount++
				biggestYet = n
				continue
			}
		}
	} else {
		panic("invalid direction")
	}

	return visibleCount
}

func (t TreeMap) CountForAColumn(col int, vantage string) int {
	visibleCount := 0
	biggestYet := 0

	if vantage == "south" {
		// facing south from the north
		for i, row := range t {
			n := row[col]
			if i == 0 || n > biggestYet {
				visibleCount++
				biggestYet = n
				continue
			}
		}
	} else if vantage == "north" {
		// facing north from the south
		endIndex := len(t) - 1
		for i := endIndex; i >= 0; i-- {
			n := t[i][col]
			if i == endIndex || n > biggestYet {
				visibleCount++
				biggestYet = n
				continue
			}
		}
	} else {
		panic("invalid direction")
	}

	return visibleCount
}
