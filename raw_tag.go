package humle

import (
	"io"
)

// Satisfies Node interface
// RawHTML is a node that represents raw HTML content.
// It does not escape the content, so it should be used with caution.
type RawHTML string

func (n RawHTML) Type() NodeType {
	return NodeRawHTML
}

func (n RawHTML) WriteTo(w io.Writer) (int64, error) {
	len, err := w.Write([]byte(n))
	return int64(len), err
}
