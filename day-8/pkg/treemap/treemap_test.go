package treemap

import "testing"

func TestCountForARow(t *testing.T) {
	t.Run("east", func(t *testing.T) {
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
		}

		row := 1
		expected := 2
		actual := m.CountForARow(row, "east")
		if expected != actual {
			t.Errorf("expected %d for row %d but got %d instead", expected, row, actual)
		}
	})

	t.Run("west", func(t *testing.T) {
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
		}

		row := 4
		expected := 2
		actual := m.CountForARow(row, "west")
		if expected != actual {
			t.Errorf("expected %d for row %d but got %d instead", expected, row, actual)
		}
	})
}

func TestCountForAColumn(t *testing.T) {
	t.Run("north", func(t *testing.T) {
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
		}

		col := 2
		expected := 2
		actual := m.CountForAColumn(col, "north")
		if expected != actual {
			t.Errorf("expected %d for column %d but got %d instead", expected, col, actual)
		}
	})

	t.Run("south", func(t *testing.T) {
		m := TreeMap{
			Data: [][]int{
				{3, 0, 3, 7, 3},
				{2, 5, 5, 1, 2},
				{6, 5, 3, 3, 2},
				{3, 3, 5, 4, 9},
				{3, 5, 3, 6, 0},
			},
			RowCount: 5,
			ColCount: 5,
		}

		col := 3
		expected := 1
		actual := m.CountForAColumn(col, "south")
		if expected != actual {
			t.Errorf("expected %d for column %d but got %d instead", expected, col, actual)
		}
	})
}
