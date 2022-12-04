package assignments

import (
	"strings"
	"testing"
)

func TestAssignmentFromString(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		input := "2-6"

		a, err := NewAssignmentFromString(input)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if a.Lower != 2 {
			t.Errorf("expected 2, got %d", a.Lower)
		}

		if a.Upper != 6 {
			t.Errorf("expected 6, got %d", a.Upper)
		}
	})

	t.Run("too many parts", func(t *testing.T) {
		input := "1-2-3"

		a, err := NewAssignmentFromString(input)

		if a != nil {
			t.Errorf("expected no assignment, got %v", a)
		}

		if err == nil {
			t.Error("expected an error")
		}

		if !strings.Contains(err.Error(), "error parsing:") {
			t.Errorf("expected 'error parsing:', got %v", err)
		}
	})

	t.Run("bad conv", func(t *testing.T) {
		input := "a-6"

		a, err := NewAssignmentFromString(input)

		if a != nil {
			t.Errorf("expected no assignment, got %v", a)
		}

		if err == nil {
			t.Error("expected an error")
		}
	})

	t.Run("nonesense range", func(t *testing.T) {
		input := "2-1"

		a, err := NewAssignmentFromString(input)

		if a != nil {
			t.Errorf("expected no assignment, got %v", a)
		}

		if err == nil {
			t.Error("expected an error")
		}

		if !strings.Contains(err.Error(), "lower is bigger than upper") {
			t.Errorf("expected 'lower is bigger than upper', got %v", err)
		}
	})
}

func TestAssignmentFullyContains(t *testing.T) {
	a := Assignment{Lower: 2, Upper: 6}

	t.Run("Fully Inside", func(t *testing.T) {
		b := Assignment{Lower: 3, Upper: 5}
		if !a.FullyContains(&b) {
			t.Errorf("Expected %v to fully contain %v", a, b)
		}
	})

	t.Run("Same Lower", func(t *testing.T) {
		b := Assignment{Lower: 2, Upper: 5}
		if !a.FullyContains(&b) {
			t.Errorf("Expected %v to fully contain %v", a, b)
		}
	})

	t.Run("Same Upper", func(t *testing.T) {
		b := Assignment{Lower: 3, Upper: 6}
		if !a.FullyContains(&b) {
			t.Errorf("Expected %v to fully contain %v", a, b)
		}
	})

	t.Run("Partial Overlap Lower", func(t *testing.T) {
		b := Assignment{Lower: 1, Upper: 5}
		if a.FullyContains(&b) {
			t.Errorf("Expected %v not to fully contain %v", a, b)
		}
	})

	t.Run("Partial Overlap Upper", func(t *testing.T) {
		b := Assignment{Lower: 3, Upper: 7}
		if a.FullyContains(&b) {
			t.Errorf("Expected %v not to fully contain %v", a, b)
		}
	})

	t.Run("No Overlap", func(t *testing.T) {
		b := Assignment{Lower: 7, Upper: 9}
		if a.FullyContains(&b) {
			t.Errorf("Expected %v not to fully contain %v", a, b)
		}
	})
}

func TestAssignmentsFromLine(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		input := "1-6,2-3"

		first, second, err := AssignmentsFromLine(input)

		if first == nil {
			t.Error("expected assignments not to be nil, but first was nil")
		}

		if second == nil {
			t.Error("expected assignments not to be nil, but second was nil")
		}

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("not enough parts", func(t *testing.T) {
		input := "1-6"

		first, second, err := AssignmentsFromLine(input)

		if first != nil {
			t.Errorf("expected assignments to be nil, but first was %v", first)
		}

		if second != nil {
			t.Errorf("expected assignments to be nil, but second was %v", second)
		}

		if err == nil {
			t.Error("expected an error")
		}

		if !strings.Contains(err.Error(), "error parsing line") {
			t.Errorf("expected 'error parsing line', got %v", err)
		}
	})

	t.Run("bad first", func(t *testing.T) {
		input := "a-b,1-6"

		first, second, err := AssignmentsFromLine(input)

		if first != nil {
			t.Errorf("expected assignments to be nil, but first was %v", first)
		}

		if second != nil {
			t.Errorf("expected assignments to be nil, but second was %v", second)
		}

		if err == nil {
			t.Error("expected an error")
		}
	})

	t.Run("bad second", func(t *testing.T) {
		input := "1-3,a-6"

		first, second, err := AssignmentsFromLine(input)

		if first == nil {
			t.Error("expected first not to be nil, but it was")
		}

		if second != nil {
			t.Errorf("expected assignments to be nil, but second was %v", second)
		}

		if err == nil {
			t.Error("expected an error")
		}
	})
}
