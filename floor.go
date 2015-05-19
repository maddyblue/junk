package vorbis

type Floor interface {
}

type Floor0 struct {
	order            uint32
	rate             uint32
	bark_map_size    uint32
	amplitude_bits   uint32
	amplitude_offset uint32
	number_of_books  uint32
	book_list        []uint32
}

type Floor1 struct {
	partitions           uint32
	partition_class_list []uint32
	class_dimensions     []uint32
	class_subclasses     []uint32
	class_masterbooks    []uint32
	subclass_books       [][]uint32
	multiplier           uint32
	X_list               []uint32
	values               uint32
}
