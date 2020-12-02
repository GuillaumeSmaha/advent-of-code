package p11

import (
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		Path   string
		Result int
	}{
		{
			Path:   "ne,ne,ne",
			Result: 3,
		},
		{
			Path:   "ne,ne,sw,sw",
			Result: 0,
		},
		{
			Path:   "ne,ne,s,s",
			Result: 2,
		},
		{
			Path:   "se,sw,se,sw,sw",
			Result: 3,
		},
		{
			Path:   "@a.txt",
			Result: 644,
		},
	}

	for i, test := range tests {
		n := Distance(test.Path)
		if n != test.Result {
			t.Fatalf("%d: %s '%d' (should be %d)", i, test.Path, n, test.Result)
		}
	}
}
