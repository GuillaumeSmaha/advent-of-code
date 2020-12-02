package p14

import (
	"testing"
)

func TestUsed(t *testing.T) {
	tests := []struct {
		Hash   string
		Result int
	}{
		{
			Hash:   "flqrgnkx",
			Result: 8108,
		},
		{
			Hash:   "vbqugkhl",
			Result: 8148,
		},
	}

	for i, test := range tests {
		n := Used(test.Hash)
		if n != test.Result {
			t.Fatalf("a.%d: %s '%d' (should be %d)", i, test.Hash, n, test.Result)
		}
	}
}

func TestRegions(t *testing.T) {
	tests := []struct {
		Hash   string
		Result int
	}{
		{
			Hash:   "flqrgnkx",
			Result: 1242,
		},
		{
			Hash:   "vbqugkhl",
			Result: 1180,
		},
	}

	for i, test := range tests {
		n := Regions(test.Hash)
		if n != test.Result {
			t.Fatalf("b.%d: %s '%d' (should be %d)", i, test.Hash, n, test.Result)
		}
	}
}
