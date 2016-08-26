package vorbis

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/mccoyst/ogg"
)

type Vorbis struct {
	o   *ogg.Decoder
	br  *bitReader
	err error

	Version     uint32
	Channels    uint8
	SampleRate  uint32
	BitrateMax  uint32
	BitrateNorm uint32
	BitrateMin  uint32
	blocksize0  int
	blocksize1  int

	Vendor   string
	Comments map[string][]string

	codebooks []*codebook
	floors    []floor
	residues  []residue
	mappings  []mapping
	modes     []mode
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
		if v.err != nil {
			return nil, v.err
		}
	}
	return v, v.err
}

func (v *Vorbis) readBits(bits uint) uint32 {
	if v.br == nil {
		p, err := v.o.Decode()
		if err != nil {
			v.err = err
			return 0
		}
		br := newBitReader(bytes.NewReader(p.Packet))
		v.br = &br
	}
	b := v.br.ReadBits(bits)
	err := v.br.Err()
	if err != nil {
		v.br = nil
		b = v.readBits(bits)
	}
	return b
}

func (v *Vorbis) readBool() bool {
	return v.readBits(1) == 1
}

func (v *Vorbis) readByte() byte {
	return byte(v.readBits(8))
}

func (v *Vorbis) decode(typ uint8) error {
	if t := v.readByte(); t != typ {
		return fmt.Errorf("unexpected header %02x, expected %02x", t, typ)
	}
	for _, c := range "vorbis" {
		if b := rune(v.readByte()); b != c {
			return fmt.Errorf("unexpected character %c, expected %c", b, c)
		}
	}
	return nil
}

func (v *Vorbis) readFloat32() float32 {
	return float32_unpack(v.readBits(32))
}

func (v *Vorbis) expect(bs ...byte) error {
	for _, b := range bs {
		r := v.readByte()
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
	ErrSetup   = errors.New("vorbis: invalid setup header")
	ErrFraming = fmt.Errorf("vorbis: expected framing bit")
)

func (v *Vorbis) decodeIdentification() error {
	if err := v.decode(typeIdentification); err != nil {
		return err
	}
	v.Version = uint32(v.readBits(32))
	v.Channels = v.readByte()
	v.SampleRate = uint32(v.readBits(32))
	v.BitrateMax = uint32(v.readBits(32))
	v.BitrateNorm = uint32(v.readBits(32))
	v.BitrateMin = uint32(v.readBits(32))
	v.blocksize0 = 1 << v.readBits(4)
	v.blocksize1 = 1 << v.readBits(4)
	if v.blocksize0 > v.blocksize1 || v.blocksize0 == 0 || v.blocksize1 == 0 {
		return fmt.Errorf("vorbis: bad blocksize")
	}
	if v.readByte() != 1 {
		return ErrFraming
	}
	return nil
}

func (v *Vorbis) decodeComment() error {
	if err := v.decode(typeComment); err != nil {
		return err
	}
	read := func() string {
		l := int(v.readBits(32))
		bytes := make([]byte, l)
		for i := 0; i < l; i++ {
			bytes[i] = byte(v.readByte())
		}
		return string(bytes)
	}
	v.Vendor = read()
	v.Comments = make(map[string][]string)
	comments := int(v.readBits(32))
	for i := 0; i < comments; i++ {
		c := read()
		sp := strings.SplitN(c, "=", 2)
		if len(sp) != 2 {
			continue
		}
		v.Comments[sp[0]] = append(v.Comments[sp[0]], sp[1])
	}
	if v.readByte() != 1 {
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
	vorbis_time_count := v.readBits(6) + 1
	for i := uint32(0); i < vorbis_time_count; i++ {
		if v.readBits(16) != 0 {
			return errors.New("vorbis: expected 0 time count value")
		}
	}

	// floors
	vorbis_floor_count := v.readBits(6) + 1
	vorbis_floor_types := make([]uint32, vorbis_floor_count)
	v.floors = make([]floor, vorbis_floor_count)
	for i := uint32(0); i < vorbis_floor_count; i++ {
		f := v.readBits(16)
		vorbis_floor_types[i] = f
		switch f {
		case 0:
			f0 := floor0{
				order:            v.readBits(8),
				rate:             v.readBits(16),
				bark_map_size:    v.readBits(16),
				amplitude_bits:   v.readBits(6),
				amplitude_offset: v.readBits(8),
				number_of_books:  v.readBits(4) + 1,
			}
			f0.book_list = make([]uint32, f0.number_of_books)
			for i := range f0.book_list {
				f0.book_list[i] = v.readBits(8)
			}
			v.floors[i] = f0
		case 1:
			f1 := floor1{
				partitions: v.readBits(5),
			}
			maximum_class := uint32(0)
			f1.partition_class_list = make([]uint32, f1.partitions)
			for i := uint32(0); i < f1.partitions; i++ {
				c := v.readBits(4)
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
				f1.class_dimensions[i] = v.readBits(3) + 1
				f1.class_subclasses[i] = v.readBits(2)
				if f1.class_subclasses[i] != 0 {
					f1.class_masterbooks[i] = v.readBits(8)
				}
				cs2 := 1 << f1.class_subclasses[i]
				f1.subclass_books[i] = make([]uint32, cs2)
				for j := 0; j < cs2; j++ {
					f1.subclass_books[i][j] = v.readBits(8) - 1
				}
			}
			f1.multiplier = v.readBits(2) + 1
			rangebits := v.readBits(4)
			f1.X_list = make([]uint32, 2)
			f1.X_list[1] = 1 << rangebits
			for i := uint32(0); i < f1.partitions; i++ {
				current_class_number := f1.partition_class_list[i]
				for j := uint32(0); j < f1.class_dimensions[current_class_number]; j++ {
					f1.X_list = append(f1.X_list, v.readBits(uint(rangebits)))
				}
			}
			v.floors[i] = f1
		default:
			return fmt.Errorf("vorbis: unknown floor type %v", f)
		}
	}

	// residues
	vorbis_residue_count := v.readBits(6) + 1
	v.residues = make([]residue, vorbis_residue_count)
	for ri := range v.residues {
		t := v.readBits(16)
		switch t {
		case 0, 1, 2:
			r := residue{
				typ:             t,
				begin:           v.readBits(24),
				end:             v.readBits(24),
				partition_size:  v.readBits(24) + 1,
				classifications: v.readBits(6) + 1,
				classbook:       v.readBits(8),
			}
			r.cascade = make([]uint32, r.classifications)
			r.books = make([][8]int64, r.classifications)
			for i := uint32(0); i < r.classifications; i++ {
				high_bits := uint32(0)
				low_bits := v.readBits(3)
				bitflag := v.readBool()
				if bitflag {
					high_bits = v.readBits(5)
				}
				r.cascade[i] = high_bits*8 + low_bits
			}
			for i := uint32(0); i < r.classifications; i++ {
				for j := uint(0); j < 8; j++ {
					if r.cascade[i]&(1<<j) != 0 {
						r.books[i][j] = int64(v.readBits(8))
					} else {
						r.books[i][j] = -1
					}
				}
			}
			v.residues[ri] = r
		default:
			return fmt.Errorf("vorbis: unknown residue type %v", t)
		}
	}

	// mappings
	vorbis_mapping_count := v.readBits(6) + 1
	v.mappings = make([]mapping, vorbis_mapping_count)
	for mi := uint32(0); mi < vorbis_mapping_count; mi++ {
		t := v.readBits(16)
		switch t {
		case 0:
			m := mapping{}
			if v.readBool() {
				m.submaps = v.readBits(4) + 1
			} else {
				m.submaps = 1
			}
			if v.readBool() {
				m.coupling_steps = v.readBits(8) + 1
				ic := uint(ilog(int64(v.Channels) - 1))
				m.magnitude = make([]uint32, ic)
				m.angle = make([]uint32, ic)
				for j := uint32(0); j < m.coupling_steps; j++ {
					m.magnitude[j] = v.readBits(ic)
					m.angle[j] = v.readBits(ic)
				}
			}
			if v.readBits(2) != 0 {
				return errors.New("vorbis: expected 0")
			}
			if m.submaps > 1 {
				m.mux = make([]uint32, v.Channels)
				for j := uint8(0); j < v.Channels; j++ {
					m.mux[j] = v.readBits(4)
				}
			}
			m.submap_floor = make([]uint32, m.submaps)
			m.submap_residue = make([]uint32, m.submaps)
			for j := uint32(0); j < m.submaps; j++ {
				v.readBits(8)
				m.submap_floor[j] = v.readBits(8)
				m.submap_residue[j] = v.readBits(8)
			}
			v.mappings[mi] = m
		default:
			return fmt.Errorf("vorbis: unknown mapping type %v", t)
		}
	}

	// modes
	vorbis_mode_count := v.readBits(6) + 1
	v.modes = make([]mode, vorbis_mode_count)
	for mi := uint32(0); mi < vorbis_mode_count; mi++ {
		m := mode{
			blockflag:     v.readBool(),
			windowtype:    v.readBits(16),
			transformtype: v.readBits(16),
			mapping:       v.readBits(8),
		}
		if m.windowtype != 0 || m.transformtype != 0 {
			return ErrSetup
		}
		v.modes[mi] = m
	}

	if v.readByte() != 1 {
		return ErrFraming
	}
	return nil
}

func (v *Vorbis) decodeCodebooks() error {
	vorbis_codebook_count := int(v.readByte()) + 1
	v.codebooks = make([]*codebook, vorbis_codebook_count)
	for i := 0; i < vorbis_codebook_count; i++ {
		c, err := v.decodeCodebook()
		if err != nil {
			return err
		}
		v.codebooks[i] = c
	}
	return nil
}

func (v *Vorbis) decodeCodebook() (*codebook, error) {
	var c codebook
	if err := v.expect(0x42, 0x43, 0x56); err != nil {
		return nil, err
	}
	c.codebook_dimensions = v.readBits(16)
	c.codebook_entries = v.readBits(24)

	// codeword lengths
	ordered := v.readBool()
	c.codebook_codeword_lengths = make([]uint32, c.codebook_entries)
	if !ordered {
		sparse := v.readBool()
		for i := uint32(0); i < c.codebook_entries; i++ {
			if sparse {
				flag := v.readBool()
				if flag {
					length := v.readBits(5) + 1
					c.codebook_codeword_lengths[i] = length
				}
			} else {
				length := v.readBits(5) + 1
				c.codebook_codeword_lengths[i] = length
			}
		}
	} else if ordered {
		current_entry := uint32(0)
		current_length := v.readBits(5) + 1
		for current_entry < c.codebook_entries {
			number := v.readBits(uint(ilog(int64(c.codebook_entries) - int64(current_entry))))
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
	codebook_lookup_type := v.readBits(4)
	switch codebook_lookup_type {
	case 0:
		// no lookup
	case 1, 2:
		codebook_minimum_value := v.readFloat32()
		codebook_delta_value := v.readFloat32()
		codebook_value_bits := v.readBits(4) + 1
		codebook_sequence_p := v.readBool()
		var codebook_lookup_values uint32
		if codebook_lookup_type == 1 {
			codebook_lookup_values = lookup1_values(c.codebook_entries, c.codebook_dimensions)
		} else {
			codebook_lookup_values = c.codebook_entries * c.codebook_dimensions
		}
		c.codebook_multiplicands = make([]uint32, codebook_lookup_values)
		for i := range c.codebook_multiplicands {
			c.codebook_multiplicands[i] = v.readBits(uint(codebook_value_bits))
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
