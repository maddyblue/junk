package vorbis

type Mapping struct {
	submaps        uint32
	coupling_steps uint32
	magnitude      []uint32
	angle          []uint32
	mux            []uint32
	submap_floor   []uint32
	submap_residue []uint32
}
