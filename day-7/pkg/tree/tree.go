package tree

import (
	"errors"
	"fmt"
	command "main/pkg/commands"
	. "main/pkg/nodes"
)

type Tree struct {
	Root *Node
	Pwd  *Node
}

func NewTree() Tree {
	root := NewRoot()

	return Tree{Root: &root, Pwd: &root}
}

func (t *Tree) visit(n *Node) {
	t.Pwd = n
}

func (t *Tree) Exec(cmdString string, opts ...command.Option) error {
	cmd, err := command.CommandFromPrompt(cmdString)

	if err != nil {
		return err
	}

	for _, o := range opts {
		o(cmd)
	}

	if err != nil {
		return err
	}

	switch cmd.Type {
	case command.CD:
		return t.handleCd(cmd)
	case command.LS:
		return t.handleLs(cmd)
	default:
		return errors.New("absurd")
	}
}

func (t *Tree) handleCd(cmd *command.Command) error {
	if cmd.Target == "/" {
		t.visit(t.Root)
		return nil
	}

	if cmd.Target == ".." {
		if t.Pwd.IsRoot() {
			t.visit(t.Root)
			return nil
		} else {
			t.visit(t.Pwd.Parent)
			return nil
		}
	}

	target := t.Pwd.FindChildByName(cmd.Target)
	if target == nil {
		return errors.New(fmt.Sprintf("pwd=%s cd: no such file or directory: %s", t.Pwd.GetFullPath(), cmd.Target))
	}

	if target.IsFile() {
		return errors.New(fmt.Sprintf("pwd=%s cd: not a directory: %s", t.Pwd.GetFullPath(), cmd.Target))
	}

	t.visit(target)
	return nil
}

func (t *Tree) handleLs(cmd *command.Command) error {
	for _, s := range cmd.Data {
		err := t.Pwd.AddChildFromLsString(s)

		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Tree) WalkWithCriteria(criteria func(*Node) bool) []*Node {
	var collector []*Node

	return visitChildren(collector, t.Root, criteria)
}

func visitChildren(collector []*Node, n *Node, criteria func(*Node) bool) []*Node {
	if criteria(n) {
		collector = append(collector, n)
		return collector
	}
	for _, c := range n.Children {
		collector = visitChildren(collector, c, criteria)
	}

	return collector
}
