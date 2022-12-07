package command

import (
	"strings"
	"testing"
)

func TestCommandFromPrompt(t *testing.T) {
	t.Run("cd", func(t *testing.T) {
		input := "$ cd foo"

		expected := Command{Type: CD, Target: "foo"}
		actual, err := CommandFromPrompt(input)

		if err != nil {
			t.Errorf("expected no error but got %s", err)
		}

		if expected.Type != actual.Type {
			t.Errorf("expected %s but got %s", expected.Type, actual.Type)
		}

		if expected.Target != actual.Target {
			t.Errorf("expected %s but got %s", expected.Target, actual.Target)
		}
	})

	t.Run("ls", func(t *testing.T) {
		input := "$ ls foo"

		expected := Command{Type: LS, Target: "foo"}
		actual, err := CommandFromPrompt(input)

		if err != nil {
			t.Errorf("expected no error but got %s", err)
		}

		if expected.Type != actual.Type {
			t.Errorf("expected %s but got %s", expected.Type, actual.Type)
		}

		if expected.Target != actual.Target {
			t.Errorf("expected %s but got %s", expected.Target, actual.Target)
		}
	})

	t.Run("no target", func(t *testing.T) {
		input := "$ ls"

		_, err := CommandFromPrompt(input)

		if err == nil {
			t.Error("expected an error but got none")
		}

		if !strings.Contains(err.Error(), "no target") {
			t.Errorf("expected 'no target', got %v", err)
		}
	})

	t.Run("no command", func(t *testing.T) {
		input := "$ cat foo"

		_, err := CommandFromPrompt(input)

		if err == nil {
			t.Error("expected an error but got none")
		}

		if !strings.Contains(err.Error(), "no parsable command from $ cat foo") {
			t.Errorf("expected 'no target', got %v", err)
		}
	})
}
