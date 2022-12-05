package crates

import "errors"

type Crate []rune

func (c Crate) Pop() (rune, Crate, error) {
	if len(c) == 0 {
		return '0', c, errors.New("Cannot pop an empty crate")
	}
	return c[len(c)-1], c[:len(c)-1], nil
}

func (c Crate) Push(r rune) Crate {
	return append(c, r)
}

func (c Crate) Last() rune {
	return c[len(c)-1]
}
