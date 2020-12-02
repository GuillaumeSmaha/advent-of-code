package p13

import (
	"testing"
)

func TestSeverity(t *testing.T) {
	tests := []struct {
		Input  string
		Result int
	}{
		{
			Input:  "@t.txt",
			Result: 24,
		},
		{
			Input:  "@a.txt",
			Result: 632,
		},
	}

	for i, test := range tests {
		n := Severity(test.Input)
		if n != test.Result {
			t.Fatalf("a.%d: %s '%d' (should be %d)", i, test.Input, n, test.Result)
		}
	}
}

func TestDelay(t *testing.T) {
	tests := []struct {
		Input  string
		Result int
	}{
		{
			Input:  "@t.txt",
			Result: 10,
		},
		{
			Input:  "@a.txt",
			Result: 3849742,
		},
	}

	for i, test := range tests {
		n := Delay2(test.Input)
		if n != test.Result {
			t.Fatalf("b.%d: %s '%d' (should be %d)", i, test.Input, n, test.Result)
		}
	}
}
