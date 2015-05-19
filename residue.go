package vorbis

type Residue struct {
	typ             uint32
	begin           uint32
	end             uint32
	partition_size  uint32
	classifications uint32
	classbook       uint32
	cascade         []uint32
	books           [][8]int64
}
