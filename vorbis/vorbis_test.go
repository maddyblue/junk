package vorbis

import (
	"os"
	"testing"
)

func TestVorbis(t *testing.T) {
	f, err := os.Open("Hydrate-Kenny_Beltrey.ogg")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	v, err := NewVorbis(f)
	if err != nil {
		t.Fatal(err)
	}
	if v.Version != 0 || v.Channels != 2 || v.SampleRate != 44100 || v.blocksize0 != 256 || v.blocksize1 != 2048 {
		t.Fatalf("bad identification")
	}
	if v.Vendor != "Xiph.Org libVorbis I 20020713" || len(v.Comments) != 6 || v.Comments["TITLE"][0] != "Hydrate - Kenny Beltrey" {
		t.Fatalf("bad comments")
	}
}
