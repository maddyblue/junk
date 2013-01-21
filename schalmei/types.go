package schalmei

import (
	"appengine"
)

type Rank struct {
	Name string
}

type Note struct {
	Freq float64
	Blob appengine.BlobKey
}
