package ioext

import "io"

// StrDeck implements io.ReadWriter. Zero value for StrDeck
// will return EOF on the first call of Read.
type StrDeck struct {
	ss []string // underlying slice
	i  int      // reading index of current string
	j  int      // index of current string
}

// NewStrDeck creates an instance of StrDeck which reads bytes
// sequentially from underlying slice of strings or writes bytes by
// appending to this slice.
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
	r.i += n
	if r.i >= len(r.ss[r.j]) {
		for r.j++; r.j < len(r.ss) && len(r.ss[r.j]) == 0; r.j++ {
			// skip empty strings to avoid calls
			// that will read zero bytes
		}
		r.i = 0
	}
	return
}

func (r *StrDeck) Write(b []byte) (n int, err error) {
	if len(b) == 0 {
		return
	}
	cp := make([]byte, len(b))
	n = copy(cp, b)
	r.ss = append(r.ss, string(cp))
	return
}

// Extend appends additional strings to underlying slice. Can be
// used even after reaching EOF via Read.
func (r *StrDeck) Extend(str ...string) {
	r.ss = append(r.ss, str...)
}
