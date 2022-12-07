package nodes

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
}

func (n *Node) IsFile() bool {
	return n.Type == File
}

func (n *Node) IsDir() bool {
	return n.Type == Directory
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
