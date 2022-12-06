package cursor

import "testing"

func TestScan(t *testing.T) {
	input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	expected := 7
	ok, output := Scan(input)

	if !ok {
		t.Error("wat")
	}

	if output != expected {
		t.Errorf("expected 7 but got %d", output)
	}
}
