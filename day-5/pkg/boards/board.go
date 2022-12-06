package boards

import (
	. "main/pkg/crates"
	. "main/pkg/instructions"
)

type Board []Crate

func (b Board) Move(inst *Instruction) error {
	r, c, err := b[inst.From].Pop(inst.Count)
	if err != nil {
		return err
	}

	b[inst.From] = c

	for i := 0; i < len(r); i++ {
		b[inst.To] = b[inst.To].Push(r[i])
	}

	return nil
}

func (b Board) Top() string {
	str := ""

	for i := 0; i < len(b); i++ {
		v := b[i]
		if len(v) == 0 {
			continue
		}

		str += string(v.Last())
	}

	return str
}
