package cursor

import (
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	t.Run("length longer than input", func(t *testing.T) {
		input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
		expected := 0
		output, err := Scan(input, 100)

		if err == nil {
			t.Error("expected an error, but got none")
		}

		if !strings.Contains(err.Error(), "length is longer than input") {
			t.Errorf("expected 'length is longer than input', got %v", err)
		}

		if output != expected {
			t.Errorf("expected 0 but got %d", output)
		}
	})

	t.Run("not found", func(t *testing.T) {
		input := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		expected := 0
		output, err := Scan(input, 4)

		if err == nil {
			t.Error("expected an error, but got none")
		}

		if !strings.Contains(err.Error(), "not found") {
			t.Errorf("expected 'not found', got %v", err)
		}

		if output != expected {
			t.Errorf("expected 0 but got %d", output)
		}
	})

	t.Run("start-of-packet", func(t *testing.T) {
		input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
		expected := 7
		output, err := Scan(input, 4)

		if err != nil {
			t.Errorf("expected no error, got: %s", err)
		}

		if output != expected {
			t.Errorf("expected 7 but got %d", output)
		}
	})

	t.Run("start-of-message", func(t *testing.T) {
		input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
		expected := 19
		output, err := Scan(input, 14)

		if err != nil {
			t.Errorf("expected no error, got: %s", err)
		}

		if output != expected {
			t.Errorf("expected 7 but got %d", output)
		}
	})
}
