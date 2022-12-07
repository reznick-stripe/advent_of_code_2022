package nodes

import (
	"strings"
	"testing"
)

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
		d := Node{Type: Directory, Name: "d", Children: []*Node{}}
		d.AddChild(&gif)
		c := Node{Type: Directory, Name: "c", Children: []*Node{}}
		c.AddChild(&exe)
		c.AddChild(&d)
		b := Node{Type: Directory, Name: "b", Children: []*Node{}}
		b.AddChild(&dat)
		b.AddChild(&img)
		a := Node{Type: Directory, Name: "a", Children: []*Node{}}
		a.AddChild(&b)
		a.AddChild(&c)
		a.AddChild(&txt)

		expected := 165
		actual := a.GetSize()

		if expected != actual {
			t.Errorf("expected %d but got %d", expected, actual)
		}
	})
}

func TestAddChild(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		d := Node{Type: Directory, Name: "d", Children: []*Node{}}

		img := Node{Type: File, Name: "img", Size: 20}

		err := d.AddChild(&img)

		if err != nil {
			t.Errorf("expected no error but got %s", err)
		}

		if len(d.Children) != 1 {
			t.Error("expected to add a child")
		}

		if img.Parent != &d {
			t.Error("expected to add a child")
		}
	})

	t.Run("error case", func(t *testing.T) {
		gif := Node{Type: File, Name: "gif", Size: 15}

		img := Node{Type: File, Name: "img", Size: 20}

		err := gif.AddChild(&img)

		if err == nil {
			t.Error("expected an error but got none")
		}

		if !strings.Contains(err.Error(), "cannot add a child to a file") {
			t.Errorf("expected 'cannot add a child to a file', got %v", err)
		}
	})
}
