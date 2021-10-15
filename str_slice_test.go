package ioext

import (
	"bytes"
	"io"
	"testing"
)

func TestStrDeck_Copy(t *testing.T) {
	tests := []struct {
		name     string
		initial  []string
		wantCopy string
	}{
		{
			name:     "nil slice",
			initial:  nil,
			wantCopy: "",
		},
		{
			name:     "empty slice",
			initial:  []string{},
			wantCopy: "",
		},
		{
			name:     "single empty string",
			initial:  []string{""},
			wantCopy: "",
		},
		{
			name:     "multiple empty strings",
			initial:  []string{"", "", "", ""},
			wantCopy: "",
		},
		{
			name:     "single string",
			initial:  []string{"abc_ 123"},
			wantCopy: "abc_ 123",
		},
		{
			name:     "two strings",
			initial:  []string{"abc_ 123", "  kk()"},
			wantCopy: "abc_ 123  kk()",
		},
		{
			name:     "leading empty string",
			initial:  []string{"", "abc_ 123", "  kk()"},
			wantCopy: "abc_ 123  kk()",
		},
		{
			name:     "unicode characters",
			initial:  []string{"abc_ 123", "[","ыййц", "", "]"},
			wantCopy: "abc_ 123[ыййц]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewStrDeck(tt.initial)
			buf := &bytes.Buffer{}
			wantN := int64(0)
			for _, str := range tt.initial {
				wantN += int64(len(str))
			}
			gotN, err := io.Copy(buf, r)
			if err != nil {
				t.Errorf("Copy to buffer error = %v", err)
				return
			}
			if gotN != wantN {
				t.Errorf("Number of copied bytes = %v, want %v", gotN, wantN)
			}
			wantBytes := []byte(tt.wantCopy)
			cmp := bytes.Compare(buf.Bytes(), wantBytes)
			if cmp != 0 {
				t.Errorf("Read bytes = %v, want = %v", buf.Bytes(), wantBytes)
			}
		})
	}
}
