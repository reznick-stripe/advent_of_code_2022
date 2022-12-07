package nodes

import (
	"bufio"
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

func TestFindOrCreateChild(t *testing.T) {
	t.Run("find file", func(t *testing.T) {
		gif := Node{Type: File, Name: "gif", Size: 15}
		d := Node{Type: Directory, Name: "d", Children: []*Node{&gif}}

		output, err := d.FindOrCreateChild(File, "gif", WithSize(15))

		if err != nil {
			t.Errorf("expected no error but got %s", err)
		}

		if output != &gif {
			t.Errorf("expected %v but got %v", &gif, output)
		}
	})

	t.Run("create file", func(t *testing.T) {
		gif := Node{Type: File, Name: "gif", Size: 15}
		img := Node{Type: File, Name: "img", Size: 15}
		d := Node{Type: Directory, Name: "d", Children: []*Node{&gif}}

		output, err := d.FindOrCreateChild(File, "img", WithSize(15))

		if err != nil {
			t.Errorf("expected no error but got %s", err)
		}

		if !output.Eql(&img) {
			t.Errorf("expected %v but got %v", &gif, output)
		}

		if len(d.Children) != 2 {
			t.Errorf("expected children lenght of 2 but got %d", len(d.Children))
		}
	})
}

func TestGetFullPath(t *testing.T) {
	t.Run("/a/b/c/gif", func(t *testing.T) {
		gif := Node{Type: File, Name: "gif", Size: 15}
		c := Node{Type: Directory, Name: "c", Children: []*Node{}}
		c.AddChild(&gif)
		b := Node{Type: Directory, Name: "b", Children: []*Node{}}
		b.AddChild(&c)
		a := Node{Type: Directory, Name: "a", Children: []*Node{}}
		a.AddChild(&b)
		n := NewRoot()
		n.AddChild(&a)

		expected := "/a/b/c/gif"

		actual := gif.GetFullPath()

		if expected != actual {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	})

	t.Run("/", func(t *testing.T) {
		n := NewRoot()
		expected := "/"

		actual := n.GetFullPath()

		if expected != actual {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	})

	t.Run("/a/b/c", func(t *testing.T) {
		c := Node{Type: Directory, Name: "c", Children: []*Node{}}
		b := Node{Type: Directory, Name: "b", Children: []*Node{}}
		b.AddChild(&c)
		a := Node{Type: Directory, Name: "a", Children: []*Node{}}
		a.AddChild(&b)
		n := NewRoot()
		n.AddChild(&a)

		expected := "/a/b/c"

		actual := c.GetFullPath()

		if expected != actual {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	})
}

func TestAddChlidFromLsString(t *testing.T) {
	n := NewRoot()
	input := `dir dfmhjhd
307728 ghpqs
dir hztjntff
dir rvstq
dir sjt
120579 whhj.pqt
dir wrmm
`

	nodes := []*Node{
		{Type: Directory, Name: "dfmhjhd"},
		{Type: File, Size: 307728, Name: "ghpqs"},
		{Type: Directory, Name: "hztjntff"},
		{Type: Directory, Name: "rvstq"},
		{Type: Directory, Name: "sjt"},
		{Type: File, Size: 120579, Name: "whhj.pqt"},
		{Type: Directory, Name: "wrmm"},
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		n.AddChildFromLsString(scanner.Text())
	}

	for _, c := range nodes {
		f := n.FindChildByName(c.Name)
		if !c.Eql(f) {
			t.Errorf("expected %v to equal %v", c, f)
		}
	}
}
