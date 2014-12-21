// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package appstats

import (
	"fmt"
)

type byteSize float64

const (
	_B          = iota
	kB byteSize = 1 << (10 * iota)
	mB
	gB
	tB
	pB
	eB
	zB
	yB
)

func (b byteSize) String() string {
	switch {
	case b >= yB:
		return fmt.Sprintf("%.2fYB", b/yB)
	case b >= zB:
		return fmt.Sprintf("%.2fZB", b/zB)
	case b >= eB:
		return fmt.Sprintf("%.2fEB", b/eB)
	case b >= pB:
		return fmt.Sprintf("%.2fPB", b/pB)
	case b >= tB:
		return fmt.Sprintf("%.2fTB", b/tB)
	case b >= gB:
		return fmt.Sprintf("%.2fGB", b/gB)
	case b >= mB:
		return fmt.Sprintf("%.2fMB", b/mB)
	case b >= kB:
		return fmt.Sprintf("%.2fKB", b/kB)
	}
	return fmt.Sprintf("%.2fB", b)
}
