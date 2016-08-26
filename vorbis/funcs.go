package vorbis

import "math"

func ilog(x int64) int {
	var r int
	for x > 0 {
		r++
		x >>= 1
	}
	return r
}

func float32_unpack(x uint32) float32 {
	mantissa := int32(x & 0x1fffff)
	sign := x & 0x80000000
	exponent := int64(x&0x7fe00000) >> 21
	if sign != 0 {
		mantissa = -mantissa
	}
	exp := math.Pow(2, float64(exponent-788))
	return float32(mantissa) * float32(exp)
}

func lookup1_values(codebook_entries, codebook_dimensions uint32) uint32 {
	entries := float64(codebook_entries)
	dim := float64(codebook_dimensions)
	r := uint32(math.Floor(math.Exp(math.Log(entries) / dim)))
	if int(math.Floor(math.Pow(float64(r+1), dim))) <= int(entries) {
		r++
	}
	return uint32(r)
}
