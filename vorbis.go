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

	Codebooks []*Codebook
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
		b = v.ReadBits(bits)
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

	// time domain transforms
	vorbis_time_count := v.ReadBits(6) + 1
	for i := uint32(0); i < vorbis_time_count; i++ {
		if v.ReadBits(16) != 0 {
			return errors.New("vorbis: expected 0 time count value")
		}
	}

	// floors
	vorbis_floor_count := v.ReadBits(6) + 1
	vorbis_floor_types := make([]uint32, vorbis_floor_count)
	vorbis_floor_configurations := make([]Floor, vorbis_floor_count)
	for i := uint32(0); i < vorbis_floor_count; i++ {
		f := v.ReadBits(16)
		vorbis_floor_types[i] = f
		switch f {
		case 0:
			f0 := Floor0{
				order:            v.ReadBits(8),
				rate:             v.ReadBits(16),
				bark_map_size:    v.ReadBits(16),
				amplitude_bits:   v.ReadBits(6),
				amplitude_offset: v.ReadBits(8),
				number_of_books:  v.ReadBits(4) + 1,
			}
			f0.book_list = make([]uint32, f0.number_of_books)
			for i := range f0.book_list {
				f0.book_list[i] = v.ReadBits(8)
			}
			vorbis_floor_configurations[i] = f0
		case 1:
			f1 := Floor1{
				partitions: v.ReadBits(5),
			}
			maximum_class := uint32(0)
			f1.partition_class_list = make([]uint32, f1.partitions)
			for i := uint32(0); i < f1.partitions; i++ {
				c := v.ReadBits(4)
				f1.partition_class_list[i] = c
				if c > maximum_class {
					maximum_class = c
				}
			}
			f1.class_dimensions = make([]uint32, maximum_class+1)
			f1.class_subclasses = make([]uint32, maximum_class+1)
			f1.class_masterbooks = make([]uint32, maximum_class+1)
			f1.subclass_books = make([][]uint32, maximum_class+1)
			for i := uint32(0); i <= maximum_class; i++ {
				f1.class_dimensions[i] = v.ReadBits(3) + 1
				f1.class_subclasses[i] = v.ReadBits(2)
				if f1.class_subclasses[i] != 0 {
					f1.class_masterbooks[i] = v.ReadBits(8)
				}
				cs2 := 1 << f1.class_subclasses[i]
				f1.subclass_books[i] = make([]uint32, cs2)
				for j := 0; j < cs2; j++ {
					f1.subclass_books[i][j] = v.ReadBits(8) - 1
				}
			}
			f1.multiplier = v.ReadBits(2) + 1
			rangebits := v.ReadBits(4)
			f1.X_list = make([]uint32, 2)
			f1.X_list[1] = 1 << rangebits
			for i := uint32(0); i < f1.partitions; i++ {
				current_class_number := f1.partition_class_list[i]
				for j := uint32(0); j < f1.class_dimensions[current_class_number]; j++ {
					f1.X_list = append(f1.X_list, v.ReadBits(uint(rangebits)))
				}
			}
			vorbis_floor_configurations[i] = f1
		default:
			return fmt.Errorf("vorbis: unknown floor type %v", f)
		}
	}

	if v.ReadByte() != 1 {
		return ErrFraming
	}
	return nil
}

func (v *Vorbis) decodeCodebooks() error {
	vorbis_codebook_count := int(v.ReadByte()) + 1
	v.Codebooks = make([]*Codebook, vorbis_codebook_count)
	for i := 0; i < vorbis_codebook_count; i++ {
		c, err := v.decodeCodebook()
		if err != nil {
			return err
		}
		v.Codebooks[i] = c
	}
	return nil
}

func (v *Vorbis) decodeCodebook() (*Codebook, error) {
	var c Codebook
	if err := v.expect(0x42, 0x43, 0x56); err != nil {
		return nil, err
	}
	c.codebook_dimensions = v.ReadBits(16)
	c.codebook_entries = v.ReadBits(24)

	// codeword lengths
	ordered := v.ReadBool()
	c.codebook_codeword_lengths = make([]uint32, c.codebook_entries)
	if !ordered {
		sparse := v.ReadBool()
		for i := uint32(0); i < c.codebook_entries; i++ {
			if sparse {
				flag := v.ReadBool()
				if flag {
					length := v.ReadBits(5) + 1
					c.codebook_codeword_lengths[i] = length
				}
			} else {
				length := v.ReadBits(5) + 1
				c.codebook_codeword_lengths[i] = length
			}
		}
	} else if ordered {
		current_entry := uint32(0)
		current_length := v.ReadBits(5) + 1
		for current_entry < c.codebook_entries {
			number := v.ReadBits(uint(ilog(int64(c.codebook_entries) - int64(current_entry))))
			for i := uint32(0); i < number; i++ {
				c.codebook_codeword_lengths[i+current_entry] = current_length
			}
			current_entry += number
			current_length++
			if current_entry > c.codebook_entries {
				return nil, fmt.Errorf("vorbis: current_entry > c.codebook_entries")
			}
		}
	}

	var lens []uint32
	for _, v := range c.codebook_codeword_lengths {
		if v > 0 {
			lens = append(lens, v)
		}
	}

	t, err := newHuffmanTree(lens)
	if err != nil {
		return nil, err
	}
	c.t = &t

	// vector lookup table
	codebook_lookup_type := v.ReadBits(4)
	switch codebook_lookup_type {
	case 0:
		// no lookup
	case 1, 2:
		codebook_minimum_value := v.ReadFloat32()
		codebook_delta_value := v.ReadFloat32()
		codebook_value_bits := v.ReadBits(4) + 1
		codebook_sequence_p := v.ReadBool()
		var codebook_lookup_values uint32
		if codebook_lookup_type == 1 {
			codebook_lookup_values = lookup1_values(c.codebook_entries, c.codebook_dimensions)
		} else {
			codebook_lookup_values = c.codebook_entries * c.codebook_dimensions
		}
		c.codebook_multiplicands = make([]uint32, codebook_lookup_values)
		for i := range c.codebook_multiplicands {
			c.codebook_multiplicands[i] = v.ReadBits(uint(codebook_value_bits))
		}
		c.value_vector = make([][]float32, c.codebook_entries)
		for lookup_offset := uint32(0); lookup_offset < c.codebook_entries; lookup_offset++ {
			c.value_vector[lookup_offset] = make([]float32, c.codebook_dimensions)
			switch codebook_lookup_type {
			case 1:
				var last float32
				index_divisor := uint32(1)
				for i := uint32(0); i < c.codebook_dimensions; i++ {
					multiplicand_offset := (lookup_offset / index_divisor) % codebook_lookup_values
					c.value_vector[lookup_offset][i] = float32(c.codebook_multiplicands[multiplicand_offset])*codebook_delta_value + codebook_minimum_value + last
					if codebook_sequence_p {
						last = c.value_vector[lookup_offset][i]
					}
					index_divisor *= codebook_lookup_values
				}
			case 2:
				var last float32
				multiplicand_offset := lookup_offset * c.codebook_dimensions
				for i := uint32(0); i < c.codebook_dimensions; i++ {
					c.value_vector[lookup_offset][i] = float32(c.codebook_multiplicands[multiplicand_offset])*codebook_delta_value + codebook_minimum_value + last
					if codebook_sequence_p {
						last = c.value_vector[lookup_offset][i]
					}
					multiplicand_offset++
				}
			}
		}
	default:
		return nil, fmt.Errorf("vorbis: unknown codebook_lookup_type: %v", codebook_lookup_type)
	}

	return &c, nil
}
