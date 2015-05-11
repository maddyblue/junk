package vorbis

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/mccoyst/ogg"
)

type Vorbis struct {
	o   *ogg.Decoder
	br  *bitReader
	Err error

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
		if v.Err != nil {
			return nil, v.Err
		}
	}
	for i := 0; true; i++ {
		if err := v.decode(0); err != nil {
			break
		}
	}
	return v, nil
}

func (v *Vorbis) ReadBits(bits uint) uint32 {
	if v.br == nil {
		p, err := v.o.Decode()
		if err != nil {
			v.Err = err
			return 0
		}
		br := newBitReader(bytes.NewReader(p.Packet))
		v.br = &br
	}
	b := v.br.ReadBits(bits)
	err := v.br.Err()
	if err != nil {
		v.br = nil
		return v.ReadBits(bits)
	}
	return b
}

func (v *Vorbis) Read(bits int) uint32 {
	var u uint32
	for offset := uint(0); bits > 0; bits, offset = bits-8, offset+8 {
		rem := uint(bits)
		if rem > 8 {
			rem = 8
		}
		v := v.ReadBits(rem)
		u |= v << offset
	}
	return u
}

func (v *Vorbis) ReadByte() byte {
	return byte(v.ReadBits(8))
}

func (v *Vorbis) decode(typ uint8) error {
	if t := v.ReadByte(); t != typ {
		return fmt.Errorf("unexpected header %02x, expected %02x", t, typ)
	}
	for _, c := range "vorbis" {
		if b := rune(v.ReadByte()); b != c {
			return fmt.Errorf("unexpected character %c, expected %c", b, c)
		}
	}
	return nil
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
	if err := v.decode(typeIdentification); err != nil {
		return err
	}
	v.Version = uint32(v.Read(32))
	v.Channels = v.ReadByte()
	v.SampleRate = uint32(v.Read(32))
	v.BitrateMax = uint32(v.Read(32))
	v.BitrateNorm = uint32(v.Read(32))
	v.BitrateMin = uint32(v.Read(32))
	v.Blocksize0 = 1 << v.Read(4)
	v.Blocksize1 = 1 << v.Read(4)
	if v.Blocksize0 > v.Blocksize1 || v.Blocksize0 == 0 || v.Blocksize1 == 0 {
		return fmt.Errorf("vorbis: bad blocksize")
	}
	if v.ReadByte() != 1 {
		return ErrFraming
	}
	return nil
}

func (v *Vorbis) decodeComment() error {
	if err := v.decode(typeComment); err != nil {
		return err
	}
	read := func() string {
		l := int(v.Read(32))
		bytes := make([]byte, l)
		for i := 0; i < l; i++ {
			bytes[i] = byte(v.ReadByte())
		}
		return string(bytes)
	}
	v.Vendor = read()
	v.Comments = make(map[string][]string)
	comments := int(v.Read(32))
	for i := 0; i < comments; i++ {
		c := read()
		sp := strings.SplitN(c, "=", 2)
		if len(sp) != 2 {
			continue
		}
		v.Comments[sp[0]] = append(v.Comments[sp[0]], sp[1])
	}
	if v.ReadByte() != 1 {
		return ErrFraming
	}
	return nil
}
