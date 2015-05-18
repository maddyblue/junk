package vorbis

type Codebook struct {
	codebook_dimensions       uint32
	codebook_entries          uint32
	codebook_codeword_lengths []uint32
	codebook_multiplicands    []uint32
	value_vector              [][]float32
	t                         *huffmanTree
}
