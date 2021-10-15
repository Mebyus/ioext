package ioext

import "io"

// StrDeck implements io.Reader. Zero value for StrDeck
// will return EOF on the first call of Read.
type StrDeck struct {
	ss []string // underlying slice
	i  int64    // reading index of current string
	j  int      // index of current string
}

// NewStrDeck creates an instance of Reader which reads
// sequentially from underlying slice of strings.
//
// Takes ownership of the slice.
func NewStrDeck(ss []string) *StrDeck {
	return &StrDeck{
		ss: ss,
	}
}

func (r *StrDeck) Read(b []byte) (n int, err error) {
	if r.j >= len(r.ss) {
		return 0, io.EOF
	}
	n = copy(b, r.ss[r.j][r.i:])
	r.i += int64(n)
	if r.i >= int64(len(r.ss[r.j])) {
		r.j++
		r.i = 0
	}
	return
}

// Extend appends additional strings to underlying slice. Can be
// used even after reaching EOF via Read.
func (r *StrDeck) Extend(str ...string) {
	r.ss = append(r.ss, str...)
}
