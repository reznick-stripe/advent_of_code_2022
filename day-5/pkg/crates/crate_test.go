package crates

import (
	"strings"
	"testing"
)

func TestPop(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		c := Crate{'A', 'B', 'C'}

		r, n, err := c.Pop(1)

		if err != nil {
			t.Errorf("expected no error but got %s", err)
		}

		if r[0] != 'C' {
			t.Errorf("expected 'C' but got %c", r)
		}

		if len(n) != 2 {
			t.Errorf("expected a length of 2 but got %d", len(n))
		}
	})

	t.Run("pop multiple", func(t *testing.T) {
		c := Crate{'A', 'B', 'C'}

		r, n, err := c.Pop(2)

		if err != nil {
			t.Errorf("expected no error but got %s", err)
		}

		if r[0] != 'B' {
			t.Errorf("expected 'B' but got %c", r)
		}

		if r[1] != 'C' {
			t.Errorf("expected 'C' but got %c", r)
		}

		if len(n) != 1 {
			t.Errorf("expected a length of 1 but got %d", len(n))
		}
	})

	t.Run("empty crate", func(t *testing.T) {
		c := Crate{}

		_, _, err := c.Pop(1)

		if err == nil {
			t.Error("expected an error but got none")
		}

		if !strings.Contains(err.Error(), "Cannot pop an empty crate") {
			t.Errorf("expected 'Cannot pop an empty crate', got %v", err)
		}
	})
}

func TestPush(t *testing.T) {
	c := Crate{'A', 'B', 'C'}

	n := c.Push('D')

	expected := Crate{'A', 'B', 'C', 'D'}
	for i, r := range expected {
		if n[i] != r {
			t.Errorf("expected the new crate to have %c in index %d, but it had %c", r, i, n[i])
		}
	}
}
