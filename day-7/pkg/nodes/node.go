package nodes

import (
	"errors"
	"fmt"
	. "main/pkg/debug"
	"regexp"
	"strconv"
	"strings"
)

type FileType int

const (
	File FileType = iota
	Directory
)

type Node struct {
	Type     FileType
	Name     string
	Size     int
	Children []*Node
	Parent   *Node
}

type Option func(n *Node)

func NewRoot() Node {
	return Node{Name: "/", Type: Directory}
}

func (n *Node) AddChildFromLsString(s string) error {
	r, err := regexp.Compile(`(?P<type_or_size>\w+) (?P<name>[\w.]+)`)

	if err != nil {
		return err
	}

	m := r.FindStringSubmatch(s)

	data := make(map[string]string)

	if len(m) > 0 {
		for i, name := range r.SubexpNames() {
			if i != 0 && name != "" {
				data[name] = m[i]
			}
		}
	}

	if strings.Contains(data["type_or_size"], "dir") {
		d := &Node{Name: data["name"], Type: Directory}

		return n.AddChild(d)
	} else {
		i, err := strconv.Atoi(data["type_or_size"])
		if err != nil {
			return err
		}

		c := &Node{Name: data["name"], Type: File, Size: i}
		return n.AddChild(c)
	}
}

func WithSize(size int) Option {
	return func(n *Node) {
		n.Size = size
	}
}

func (n *Node) IsRoot() bool {
	return n.Parent == nil
}

func (n *Node) IsFile() bool {
	return n.Type == File
}

func (n *Node) IsDir() bool {
	return n.Type == Directory
}

func (n *Node) GetFullPath() string {
	return recursivelyGetPath(n)
}

func recursivelyGetPath(n *Node) string {
	if n == nil {
		return ""
	}

	if n.IsRoot() {
		return "/"
	}

	if n.Parent.IsRoot() {
		return "/" + n.Name
	}

	return recursivelyGetPath(n.Parent) + "/" + n.Name
}

func (n *Node) AddChild(o *Node) error {
	if n.IsFile() {
		return errors.New("cannot add a child to a file")
	}

	o.Parent = n
	n.Children = append(n.Children, o)

	if Debug() {
		LogIt(fmt.Sprintf("add_child=%s child_type=%s", o.GetFullPath(), o.TypeString()))
	}
	return nil
}

func (n *Node) Eql(o *Node) bool {
	return n.Type == o.Type && n.Name == o.Name && n.Size == o.Size
}

func (n *Node) FindChildByName(s string) *Node {
	for _, c := range n.Children {
		if c.Name == s {
			return c
		}
	}

	return nil
}

func (n *Node) FindOrCreateChild(fileType FileType, name string, opts ...Option) (*Node, error) {
	tmp := &Node{
		Type: fileType,
		Name: name,
	}

	for _, opt := range opts {
		opt(tmp)
	}

	c := findChild(tmp, n)

	if c != nil {
		return c, nil
	}

	err := n.AddChild(tmp)

	if err != nil {
		return nil, err
	}

	return tmp, nil
}

func findChild(tmp *Node, n *Node) *Node {
	for _, c := range n.Children {
		if c.Eql(tmp) {
			return c
		}
	}
	return nil
}

func (n *Node) GetSize() int {
	return getSizeRecursively(n)
}

func getSizeRecursively(n *Node) int {
	sum := n.Size

	if n.IsDir() {
		for _, c := range n.Children {
			sum += getSizeRecursively(c)
		}
	}

	return sum
}

func (n *Node) TypeString() string {
	switch n.Type {
	case File:
		return "file"
	case Directory:
		return "directory"
	default:
		return ""
	}
}
