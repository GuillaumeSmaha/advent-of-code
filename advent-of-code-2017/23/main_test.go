package p23

import (
	"testing"
)

func TestMuls(t *testing.T) {
	tests := []struct {
		Input  string
		Result int
	}{
		{
			Input:  "@a.txt",
			Result: 3025,
		},
	}

	for i, test := range tests {
		n := Muls(test.Input)
		if n != test.Result {
			t.Fatalf("a.%d: %s '%d' (should be %d)", i, test.Input, n, test.Result)
		}
	}
}

func TestDebug(t *testing.T) {
	tests := []struct {
		Result int
	}{
		{
			Result: 915,
		},
	}

	for i, test := range tests {
		n := Debug()
		if n != test.Result {
			t.Fatalf("b.%d: '%d' (should be %d)", i, n, test.Result)
		}
	}
}
