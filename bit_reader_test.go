package vorbis

import (
	"bytes"
	"testing"
)

func TestBitReader(t *testing.T) {
	b := newBitReader(bytes.NewReader([]byte{252, 72, 206, 6}))
	tests := []struct {
		bits uint
		val  uint32
	}{
		{4, 12},
		{3, 7},
		{7, 17},
		{13, 6969},
	}
	for _, bt := range tests {
		v := b.ReadBits(bt.bits)
		if b.Err() != nil {
			t.Fatal(b.Err())
		}
		if v != bt.val {
			t.Fatalf("expected %v, got %v", bt.val, v)
		}
	}
}
