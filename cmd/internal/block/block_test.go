package block

import "testing"

func TestBlockRange(t *testing.T) {
	blockRange := NewBlockRange(0, 10000)

	if len := blockRange.Len(); len != 10000 {
		t.Error("Expected block range to be 10000 but got", len)
	}

	if min := blockRange.Min(); min != 0 {
		t.Error("Expected block minimun to be 0 but got", min)
	}

	if max := blockRange.Max(); max != 10000 {
		t.Error("Expected block maximum to be 10000 but got", max)
	}
}

func TestBlockRangeMaxIsZero(t *testing.T) {
	blockRange := NewBlockRange(0, 0)

	if len := blockRange.Len(); len != 0 {
		t.Error("Expected block range to be 0 but got", len)
	}
}
