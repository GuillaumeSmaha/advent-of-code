package p4

import (
	"testing"
)

func TestValidCount(t *testing.T) {
	tests := []struct {
		Input  string
		Result int
	}{
		{
			Input:  "@ta.txt",
			Result: 2,
		},
		{
			Input:  "@a.txt",
			Result: 451,
		},
	}

	for i, test := range tests {
		n := ValidCount(test.Input)
		if n != test.Result {
			t.Fatalf("a.%d: %d (should be %d)", i, n, test.Result)
		}
	}
}

func TestNewValidCount(t *testing.T) {
	tests := []struct {
		Input  string
		Result int
	}{
		{
			Input:  "@tb.txt",
			Result: 3,
		},
		{
			Input:  "@b.txt",
			Result: 223,
		},
	}

	for i, test := range tests {
		n := NewValidCount(test.Input)
		if n != test.Result {
			t.Fatalf("b.%d: %d (should be %d)", i, n, test.Result)
		}
	}
}
