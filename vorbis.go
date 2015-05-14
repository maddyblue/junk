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

	codebook_codeword_lengths []uint32
	codebook_multiplicands    []uint32
}

func NewVorbis(r io.Reader) (*Vorbis, error) {
	v := &Vorbis{
		o: ogg.NewDecoder(r),
	}
	steps := []func() error{
		v.decodeIdentification,
		v.decodeComment,
		v.decodeSetup,
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

func (v *Vorbis) ReadBool() bool {
	return v.ReadBits(1) == 1
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

func (v *Vorbis) ReadFloat32() float32 {
	return float32_unpack(v.ReadBits(32))
}

func (v *Vorbis) expect(bs ...byte) error {
	for _, b := range bs {
		r := v.ReadByte()
		if r != b {
			return fmt.Errorf("vorbis: expected %02x, got %02x", b, r)
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
	v.Version = uint32(v.ReadBits(32))
	v.Channels = v.ReadByte()
	v.SampleRate = uint32(v.ReadBits(32))
	v.BitrateMax = uint32(v.ReadBits(32))
	v.BitrateNorm = uint32(v.ReadBits(32))
	v.BitrateMin = uint32(v.ReadBits(32))
	v.Blocksize0 = 1 << v.ReadBits(4)
	v.Blocksize1 = 1 << v.ReadBits(4)
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
		l := int(v.ReadBits(32))
		bytes := make([]byte, l)
		for i := 0; i < l; i++ {
			bytes[i] = byte(v.ReadByte())
		}
		return string(bytes)
	}
	v.Vendor = read()
	v.Comments = make(map[string][]string)
	comments := int(v.ReadBits(32))
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

func (v *Vorbis) decodeSetup() error {
	if err := v.decode(typeSetup); err != nil {
		return err
	}
	if err := v.decodeCodebooks(); err != nil {
		return err
	}
	if v.ReadByte() != 1 {
		return ErrFraming
	}
	return nil
}

func (v *Vorbis) decodeCodebooks() error {
	vorbis_codebook_count := int(v.ReadByte()) + 1
	if err := v.expect(0x42, 0x43, 0x56); err != nil {
		return err
	}
	codebook_dimensions := v.Read(16)
	codebook_entries := v.Read(24)

	// codeword lengths
	ordered := v.ReadBool()
	v.codebook_codeword_lengths = make([]uint32, codebook_entries)
	if !ordered {
		sparse := v.ReadBool()
		for i := uint32(0); i < codebook_entries; i++ {
			if sparse {
				flag := v.ReadBool()
				if flag {
					length := v.Read(5) + 1
					v.codebook_codeword_lengths[i] = length
				} else {
					// leave nil
				}
			} else {
				length := v.Read(5) + 1
				v.codebook_codeword_lengths[i] = length
			}
		}
	} else if ordered {
		current_entry := uint32(0)
		current_length := v.Read(5) + 1
		for current_entry < codebook_entries {
			number := v.Read(uint(ilog(int64(codebook_entries) - int64(current_entry))))
			for i := uint32(0); i < number; i++ {
				v.codebook_codeword_lengths[i+current_entry] = current_length
			}
			current_entry += number
			current_length++
			if current_entry > codebook_entries {
				return fmt.Errorf("vorbis: current_entry > codebook_entries")
			}
		}
	}

	// vector lookup table
	codebook_lookup_type := v.Read(4)
	switch codebook_lookup_type {
	case 0:
		// no lookup
	case 1, 2:
		codebook_minimum_value := v.ReadFloat32()
		codebook_delta_value := v.ReadFloat32()
		codebook_value_bits := v.Read(4) + 1
		codebook_sequence_p := v.ReadBool()
		var codebook_lookup_values uint32
		if codebook_lookup_type == 1 {
			codebook_lookup_values = lookup1_values(codebook_entries, codebook_dimensions)
		} else {
			codebook_lookup_values = codebook_entries * codebook_dimensions
		}
		v.codebook_multiplicands = make([]uint32, codebook_lookup_values)
		for i := range v.codebook_multiplicands {
			v.codebook_multiplicands[i] = v.Read(uint(codebook_value_bits))
		}
	default:
		return fmt.Errorf("vorbis: unknown codebook_lookup_type: %v", codebook_lookup_type)
	}

	return nil
}
