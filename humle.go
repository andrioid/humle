package humle

import "io"

func WriteHTML(w io.Writer, n NodeWriter) (int64, error) {
	return n.WriteTo(w)
}
