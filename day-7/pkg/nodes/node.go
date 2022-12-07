package nodes

import "errors"

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

func WithSize(size int) Option {
	return func(n *Node) {
		n.Size = size
	}
}

func (n *Node) IsRoot() bool {
	return n.Name == "/" && n.Parent == nil
}

func (n *Node) IsFile() bool {
	return n.Type == File
}

func (n *Node) IsDir() bool {
	return n.Type == Directory
}

func (n *Node) AddChild(o *Node) error {
	if n.IsFile() {
		return errors.New("cannot add a child to a file")
	}

	o.Parent = n
	n.Children = append(n.Children, o)
	return nil
}

func (n *Node) Eql(o *Node) bool {
	return n.Type == o.Type && n.Name == o.Name && n.Size == o.Size
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
