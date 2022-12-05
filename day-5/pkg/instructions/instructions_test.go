package instructions

import (
	"strings"
	"testing"
)

func TestNewInstructionFromInput(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		input := "move 1 from 2 to 3"

		i, err := NewInstructionFromInput(input)

		if err != nil {
			t.Errorf("Expected no error but got %s", err)
		}

		if i.Count != 1 {
			t.Errorf("Expected Count to be 1 but was %d", i.Count)
		}

		if i.From != 2 {
			t.Errorf("Expected From to be 2 but was %d", i.From)
		}

		if i.To != 3 {
			t.Errorf("Expected To to be 3 but was %d", i.To)
		}
	})

	t.Run("bad string", func(t *testing.T) {
		input := "move "

		_, err := NewInstructionFromInput(input)

		if err == nil {
			t.Error("Expected an error but got nones")
		}

		if !strings.Contains(err.Error(), "bad parse") {
			t.Errorf("expected 'bad parse', got %v", err)
		}
	})

	t.Run("bad numbers", func(t *testing.T) {
		input := "move 1 from b to 0"

		_, err := NewInstructionFromInput(input)

		if err == nil {
			t.Error("Expected an error but got nones")
		}

		if !strings.Contains(err.Error(), "bad parse") {
			t.Errorf("expected 'bad parse', got %v", err)
		}
	})
}
