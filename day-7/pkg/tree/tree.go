package tree

import (
	"errors"
	"fmt"
	command "main/pkg/commands"
	. "main/pkg/debug"
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

func (t *Tree) ReturnToRoot() {
	t.Visit(t.Root)
}

func (t *Tree) Visit(n *Node) {
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

	if Debug() {
		LogIt(fmt.Sprintf("pwd_full_path=%s pwd_name=%s cmd=%s", t.Pwd.GetFullPath(), t.Pwd.Name, cmdString))
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
		t.Visit(t.Root)
		return nil
	}

	if cmd.Target == ".." {
		if t.Pwd.IsRoot() {
			t.Visit(t.Root)
			return nil
		} else {
			t.Visit(t.Pwd.Parent)
			return nil
		}
	}

	target := t.Pwd.FindChildByName(cmd.Target)
	if target == nil {
		if Debug() {
			LogIt(fmt.Sprintf("error=true error=404 target=%s", target.GetFullPath()))
		}
		return errors.New(fmt.Sprintf("pwd=%s cd: no such file or directory: %s", t.Pwd.GetFullPath(), cmd.Target))
	}

	if target.IsFile() {
		if Debug() {
			LogIt(fmt.Sprintf("error=true error=not_a_dir target_path=%s target_type=%s", target.GetFullPath(), target.TypeString()))
		}
		return errors.New(fmt.Sprintf("pwd=%s cd: not a directory: %s", t.Pwd.GetFullPath(), cmd.Target))
	}

	t.Visit(target)
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
	}

	for _, c := range n.Children {
		collector = visitChildren(collector, c, criteria)
	}

	return collector
}
