package vorbis

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/mccoyst/ogg"
)

type Vorbis struct {
	o *ogg.Decoder

	Version     uint32
	Channels    uint8
	SampleRate  uint32
	BitrateMax  uint32
	BitrateNorm uint32
	BitrateMin  uint32
	Blocksize0  int
	Blocksize1  int

	Vendor   string
	Comments map[string][]string
}

func NewVorbis(r io.Reader) (*Vorbis, error) {
	v := &Vorbis{
		o: ogg.NewDecoder(r),
	}
	steps := []func() error{
		v.decodeIdentification,
		v.decodeComment,
	}
	for _, step := range steps {
		if err := step(); err != nil {
			return nil, err
		}
	}
	for i := 0; true; i++ {
		_, err := v.decode(0)
		if err != nil {
			break
		}
	}
	return v, nil
}

type oggPage struct {
	ogg.Page
	*bitReader
}

func (o *oggPage) Read(bits int) uint32 {
	var u uint32
	for offset := uint(0); bits > 0; bits, offset = bits-8, offset+8 {
		rem := uint(bits)
		if rem > 8 {
			rem = 8
		}
		v := o.ReadBits(rem)
		u |= v << offset
	}
	return u
}

func (v *Vorbis) decode(typ uint8) (*oggPage, error) {
	p, err := v.o.Decode()
	if err != nil {
		return nil, err
	}
	br := newBitReader(bytes.NewReader(p.Packet))
	o := &oggPage{
		Page:      p,
		bitReader: &br,
	}
	if t := uint8(o.ReadBits(8)); t != typ {
		return nil, fmt.Errorf("unexpected header %02x, expected %02x", t, typ)
	}
	for _, c := range "vorbis" {
		if b := rune(o.ReadBits(8)); b != c {
			return nil, fmt.Errorf("unexpected character %c, expected %c", b, c)
		}
	}
	return o, o.Err()
}

const (
	typeIdentification = 1
	typeComment        = 3
	typeSetup          = 5

	pageContinued = 1
	pageBegin     = 2
	pageEnd       = 4
)

var (
	ErrFraming = fmt.Errorf("vorbis: expected framing bit")
)

func (v *Vorbis) decodeIdentification() error {
	p, err := v.decode(typeIdentification)
	if err != nil {
		return err
	}
	v.Version = uint32(p.Read(32))
	v.Channels = uint8(p.Read(8))
	v.SampleRate = uint32(p.Read(32))
	v.BitrateMax = uint32(p.Read(32))
	v.BitrateNorm = uint32(p.Read(32))
	v.BitrateMin = uint32(p.Read(32))
	v.Blocksize0 = 1 << p.Read(4)
	v.Blocksize1 = 1 << p.Read(4)
	if v.Blocksize0 > v.Blocksize1 || v.Blocksize0 == 0 || v.Blocksize1 == 0 {
		return fmt.Errorf("vorbis: bad blocksize")
	}
	if p.Read(1) != 1 {
		return ErrFraming
	}
	return p.Err()
}

func (v *Vorbis) decodeComment() error {
	p, err := v.decode(typeComment)
	if err != nil {
		return err
	}
	read := func() string {
		l := int(p.Read(32))
		bytes := make([]byte, l)
		for i := 0; i < l; i++ {
			bytes[i] = byte(p.Read(8))
		}
		return string(bytes)
	}
	v.Vendor = read()
	v.Comments = make(map[string][]string)
	comments := int(p.Read(32))
	for i := 0; i < comments; i++ {
		c := read()
		sp := strings.SplitN(c, "=", 2)
		if len(sp) != 2 {
			continue
		}
		v.Comments[sp[0]] = append(v.Comments[sp[0]], sp[1])
	}
	if p.Read(1) != 1 {
		return ErrFraming
	}
	return p.Err()
}
