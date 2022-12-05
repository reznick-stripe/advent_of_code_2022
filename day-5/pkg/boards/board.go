package boards

import (
	. "main/pkg/crates"
	. "main/pkg/instructions"
)

type Board map[int]Crate

func (b Board) Move(inst *Instruction) error {
	for i := 0; i < inst.Count; i++ {
		r, c, err := b[inst.From].Pop()
		if err != nil {
			return err
		}

		b[inst.From] = c
		b[inst.To] = b[inst.To].Push(r)
	}

	return nil
}

func (b Board) Top() string {
	str := ""

	for _, v := range b {
		if len(v) == 0 {
			continue
		}

		str += string(v.Last())
	}

	return str
}
