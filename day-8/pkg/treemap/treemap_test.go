package treemap

import "testing"

func TestCount(t *testing.T) {
	m := TreeMap{
		Data: [][]int{
			{3, 0, 3, 7, 3}, // W->E: 2 E->W: 2
			{2, 5, 5, 1, 2}, // W->E: 2 E->W: 2
			{6, 5, 3, 3, 2}, // W->E: 1 E->W: 2
			{3, 3, 5, 4, 9}, // W->E: 3 E->W: 1
			{3, 5, 3, 9, 0}, // W->E: 3 E->W: 2
			/*
							V   V  V  V  V
				N->S	2   2  2  1  2
				S->N	2   1  2  1  2
			*/
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

	expected := 21

	output := m.Count()

	if expected != output {
		t.Errorf("expected %d but got %d", expected, output)
	}
}
