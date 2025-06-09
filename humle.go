package humle

import (
	"io"
	"strings"
)

func WriteHTML(w io.Writer, n Node) (int64, error) {
	return n.WriteTo(w)
}

func WriteString(n Node) (string, error) {
	if n == nil {
		return "", nil
	}
	var b strings.Builder
	_, err := n.WriteTo(&b)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
