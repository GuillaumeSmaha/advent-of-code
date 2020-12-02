package p21

import (
	"testing"
)

func TestPixels(t *testing.T) {
	tests := []struct {
		Input      string
		Iterations int
		Result     int
	}{
		{
			Input:      "@t.txt",
			Iterations: 2,
			Result:     12,
		},
		{
			Input:      "@a.txt",
			Iterations: 5,
			Result:     205,
		},
		{
			Input:      "@a.txt",
			Iterations: 18,
			Result:     3389823,
		},
	}

	for i, test := range tests {
		n := Pixels(test.Input, test.Iterations)
		if n != test.Result {
			t.Fatalf("a.%d: %s @ %d '%d' (should be %d)", i, test.Iterations, test.Input, n, test.Result)
		}
	}
}
