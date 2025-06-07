package humle

import "io"

func WriteHTML(w io.Writer, n Node) (int64, error) {
	return n.WriteTo(w)
}
