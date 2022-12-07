package tree

import (
	commands "main/pkg/commands"
	nodes "main/pkg/nodes"
	"strings"
	"testing"
)

func TestTreeExec(t *testing.T) {
	t.Run("cd happy path", func(t *testing.T) {
		r := NewTree()
		d := nodes.Node{Type: nodes.Directory, Name: "d", Children: []*nodes.Node{}}

		r.Root.AddChild(&d)

		cmdString := "$ cd d"

		err := r.Exec(cmdString)

		if err != nil {
			t.Errorf("expected no error but got %s", err)
		}

		if !r.Pwd.Eql(&d) {
			t.Error("expected to update pwd")
		}
	})

	t.Run("cd ..", func(t *testing.T) {
		r := NewTree()
		d := nodes.Node{Type: nodes.Directory, Name: "d", Children: []*nodes.Node{}}

		r.Root.AddChild(&d)

		r.Pwd = &d

		cmdString := "$ cd .."

		err := r.Exec(cmdString)

		if err != nil {
			t.Errorf("expected no error but got %s", err)
		}

		if !r.Pwd.Eql(r.Root) {
			t.Error("expected to update pwd")
		}
	})

	t.Run("cd /", func(t *testing.T) {
		r := NewTree()
		d := nodes.Node{Type: nodes.Directory, Name: "d", Children: []*nodes.Node{}}

		r.Root.AddChild(&d)

		r.Pwd = &d

		cmdString := "$ cd /"

		err := r.Exec(cmdString)

		if err != nil {
			t.Errorf("expected no error but got %s", err)
		}

		if !r.Pwd.Eql(r.Root) {
			t.Error("expected to update pwd")
		}
	})

	t.Run("cd a non existent target", func(t *testing.T) {
		r := NewTree()
		d := nodes.Node{Type: nodes.Directory, Name: "d", Children: []*nodes.Node{}}

		r.Root.AddChild(&d)

		cmdString := "$ cd b"

		err := r.Exec(cmdString)

		if err == nil {
			t.Error("expected an error but got none")
		}

		if !strings.Contains(err.Error(), "no such file or directory") {
			t.Errorf("expected 'no such file or directory', but got %s", err)
		}
	})

	t.Run("cd file", func(t *testing.T) {
		r := NewTree()
		d := nodes.Node{Type: nodes.File, Name: "d", Children: []*nodes.Node{}}

		r.Root.AddChild(&d)

		cmdString := "$ cd d"

		err := r.Exec(cmdString)

		if err == nil {
			t.Error("expected an error but got none")
		}

		if !strings.Contains(err.Error(), "not a directory") {
			t.Errorf("expected 'not a directory', but got %s", err)
		}
	})

	t.Run("ls happy path", func(t *testing.T) {
		r := NewTree()
		d := nodes.Node{Type: nodes.File, Name: "d", Children: []*nodes.Node{}}

		r.Root.AddChild(&d)

		cmdString := "$ ls d"
		data := []string{
			"dir dfmhjhd",
			"307728 ghpqs",
			"dir hztjntff",
			"dir rvstq",
			"dir sjt",
			"120579 whhj.pqt",
			"dir wrmm",
		}

		r.Exec(cmdString, commands.WithData(data))

		nodes := []*nodes.Node{
			{Type: nodes.Directory, Name: "dfmhjhd"},
			{Type: nodes.File, Size: 307728, Name: "ghpqs"},
			{Type: nodes.Directory, Name: "hztjntff"},
			{Type: nodes.Directory, Name: "rvstq"},
			{Type: nodes.Directory, Name: "sjt"},
			{Type: nodes.File, Size: 120579, Name: "whhj.pqt"},
			{Type: nodes.Directory, Name: "wrmm"},
		}

		for _, c := range nodes {
			f := r.Pwd.FindChildByName(c.Name)

			if f == nil {
				t.Errorf("expected %v but got nil", c)
			}

			if !c.Eql(f) {
				t.Errorf("expected %v to equal %v", c, f)
			}
		}
	})
}

func TestWalkWithCriteria(t *testing.T) {
	txt := nodes.Node{Type: nodes.File, Name: "txt", Size: 50}
	img := nodes.Node{Type: nodes.File, Name: "img", Size: 20}
	dat := nodes.Node{Type: nodes.File, Name: "dat", Size: 10}
	exe := nodes.Node{Type: nodes.File, Name: "exe", Size: 70}
	gif := nodes.Node{Type: nodes.File, Name: "gif", Size: 15}
	d := nodes.Node{Type: nodes.Directory, Name: "d", Children: []*nodes.Node{}}
	d.AddChild(&gif) // size = 15
	c := nodes.Node{Type: nodes.Directory, Name: "c", Children: []*nodes.Node{}}
	c.AddChild(&exe)
	c.AddChild(&d) // size = 85
	b := nodes.Node{Type: nodes.Directory, Name: "b", Children: []*nodes.Node{}}
	b.AddChild(&dat)
	b.AddChild(&img) // size = 30
	a := nodes.Node{Type: nodes.Directory, Name: "a", Children: []*nodes.Node{}}
	a.AddChild(&b)
	a.AddChild(&c)
	a.AddChild(&txt) // size = 165

	r := NewTree()

	r.Root.AddChild(&a)

	criteria := func(n *nodes.Node) bool {
		return n.GetSize() <= 30 && n.IsDir()
	}

	expected := []*nodes.Node{
		&b,
		&d,
	}

	output := r.WalkWithCriteria(criteria)

	for i, c := range output {
		if !c.Eql(expected[i]) {
			t.Errorf("expected %v, actual %v", expected[i], c)
		}
	}

}
