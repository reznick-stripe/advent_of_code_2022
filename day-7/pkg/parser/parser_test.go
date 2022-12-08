package parser

import (
	"bufio"
	"fmt"
	. "main/pkg/nodes"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	input := `$ cd /
$ ls
dir dfmhjhd
307728 ghpqs
dir hztjntff
dir rvstq
dir sjt
120579 whhj.pqt
dir wrmm
$ cd dfmhjhd
$ ls
301486 ngtqtf
13488 wfgqtw.sqr
$ cd ..
$ cd hztjntff
$ ls
dir cwsf
288227 ftq.cjn
176977 hwtj
234858 nzdgz.mpw
157857 rhs.mbd
dir sthqhrc
`
	scanner := bufio.NewScanner(strings.NewReader(input))

	err, tree := Parse(scanner)

	if err != nil {
		t.Errorf("expected no error but got %s", err)
	}

	if !tree.Pwd.IsRoot() {
		t.Error("tree must return to root")
	}
	expectedNames := []string{
		"dfmhjhd",
		"ghpqs",
		"hztjntff",
		"rvstq",
		"sjt",
		"whhj.pqt",
		"wrmm",
	}

	for i, s := range expectedNames {
		e := tree.Pwd.Children[i].Name
		if e != s {
			t.Errorf("expected %s to be %s", e, s)
		}
	}

	tree.Visit(tree.Pwd.FindChildByName("dfmhjhd"))

	expectedNames = []string{
		"ngtqtf",
		"wfgqtw.sqr",
	}

	for i, s := range expectedNames {
		e := tree.Pwd.Children[i].Name
		if e != s {
			t.Errorf("expected %s to be %s", e, s)
		}
	}

	tree.Visit(tree.Pwd.Parent)
	tree.Visit(tree.Pwd.FindChildByName("hztjntff"))
	expectedNames = []string{
		"cwsf",
		"ftq.cjn",
		"hwtj",
		"nzdgz.mpw",
		"rhs.mbd",
		"sthqhrc",
	}

	for i, s := range expectedNames {
		e := tree.Pwd.Children[i].Name
		if e != s {
			t.Errorf("expected %s to be %s", e, s)
		}
	}
}

func TestEndToEnd(t *testing.T) {
	input := `$ cd /
$ ls
dir hztjntff
$ cd hztjntff
$ ls
dir cwsf
288227 ftq.cjn
176977 hwtj
234858 nzdgz.mpw
157857 rhs.mbd
dir sthqhrc
$ cd cwsf
$ ls
dir msr
272534 mvfclz.srm
226182 nldtb
252294 qnwzs
$ cd msr
$ ls
272198 hwtj
$ cd ..
$ cd ..
$ ls
45847 vzntwhzl.htf
`

	//path=/ =>                 1926974
	//path=/hztjntff =>         1926974
	//path=/hztjntff/cwsf =>    1023208
	//path=/hztjntff/cwsf/msr => 272198
	//path=/hztjntff/sthqhrc =>   45847

	scanner := bufio.NewScanner(strings.NewReader(input))

	err, tree := Parse(scanner)

	if err != nil {
		t.Errorf("expected no error but got %s", err)
	}

	criteria := func(n *Node) bool {
		s := n.GetSize()

		return s >= 150_000 && n.IsDir()
	}

	results := tree.WalkWithCriteria(criteria)
	if len(results) != 4 {
		for _, r := range results {
			t.Log(fmt.Sprintf("%v", r))
		}
		t.Error("not 4")
	}
}
