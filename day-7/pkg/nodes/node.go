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
