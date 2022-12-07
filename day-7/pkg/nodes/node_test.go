package nodes

import "testing"

func TestGetSize(t *testing.T) {
	t.Run("leaf node", func(t *testing.T) {
		n := Node{Type: File, Name: "a", Size: 1}

		expected := 1
		actual := n.GetSize()

		if expected != actual {
			t.Errorf("expected %d but got %d", expected, actual)
		}
	})

	t.Run("children", func(t *testing.T) {
		txt := Node{Type: File, Name: "txt", Size: 50}
		img := Node{Type: File, Name: "img", Size: 20}
		dat := Node{Type: File, Name: "dat", Size: 10}
		exe := Node{Type: File, Name: "exe", Size: 70}
		gif := Node{Type: File, Name: "gif", Size: 15}
		d := Node{Type: Directory, Name: "d", Children: []*Node{&gif}}
		c := Node{Type: Directory, Name: "c", Children: []*Node{&d, &exe}}
		b := Node{Type: Directory, Name: "b", Children: []*Node{&dat, &img}}
		a := Node{Type: Directory, Name: "a", Children: []*Node{&b, &c, &txt}}

		expected := 165
		actual := a.GetSize()

		if expected != actual {
			t.Errorf("expected %d but got %d", expected, actual)
		}
	})
}
