package boards

import (
	"bufio"
	. "main/pkg/crates"
	. "main/pkg/instructions"
	"strings"
	"testing"
)

func TestBoardMove(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		b := Board{Crate{'A', 'B', 'C'}, Crate{'D', 'E', 'F'}}

		inst := Instruction{From: 0, To: 1, Count: 1}
		err := b.Move(&inst)

		if err != nil {
			t.Errorf("expected no error but got %s", err)
		}

		expectedZero := Crate{'A', 'B'}
		for i, r := range expectedZero {
			if b[0][i] != r {
				t.Errorf("expected the new crate to have %c in index %d, but it had %c", r, i, b[0][i])
			}
		}

		expectedOne := Crate{'D', 'E', 'F', 'C'}
		for i, r := range expectedOne {
			if b[1][i] != r {
				t.Errorf("expected the new crate to have %c in index %d, but it had %c", r, i, b[1][i])
			}
		}
	})

	t.Run("end 2 end", func(t *testing.T) {
		b := Board{Crate{'A', 'B', 'C'}, Crate{'D', 'E', 'F'}, Crate{'G', 'H', 'I'}}
		input := ` move 1 from 0 to 1
 move 1 from 2 to 0
 move 2 from 1 to 2`
		// start: [ABC] [DEF] [GHI]
		// 1.     [AB] [DEFC] [GHI]
		// 2.     [ABI] [DEFC] [GH]
		// 2.     [ABI] [DE] [GHFC]

		scanner := bufio.NewScanner(strings.NewReader(input))

		for scanner.Scan() {
			inst, err := NewInstructionFromInput(scanner.Text())
			if err != nil {
				t.Errorf("expected no error but got %s", err)
			}
			err = b.Move(inst)
		}

		output := b.Top()

		if output != "IEC" {
			t.Errorf("expected 'IEC' but got %s", output)
		}
	})

	t.Run("empty source", func(t *testing.T) {
		b := Board{Crate{}, Crate{'D', 'E', 'F'}}

		inst := Instruction{From: 0, To: 1, Count: 1}
		err := b.Move(&inst)

		if err == nil {
			t.Error("expected an error but got none")
		}

		if !strings.Contains(err.Error(), "Cannot pop an empty crate") {
			t.Errorf("expected 'Cannot pop an empty crate', got %v", err)
		}
	})

	t.Run("nonesense instructions", func(t *testing.T) {
		b := Board{Crate{}, Crate{'D', 'E', 'F'}}

		inst := Instruction{From: 0, To: 1, Count: 1}
		err := b.Move(&inst)

		if err == nil {
			t.Error("expected an error but got none")
		}

		if !strings.Contains(err.Error(), "Cannot pop an empty crate") {
			t.Errorf("expected 'Cannot pop an empty crate', got %v", err)
		}
	})
}

func TestBoardLast(t *testing.T) {
	b := Board{Crate{'A', 'B', 'C'}, Crate{}, Crate{'D', 'E', 'F'}}

	s := b.Top()

	if s != "CF" {
		t.Errorf("expected 'CF', but got %s", s)
	}
}
