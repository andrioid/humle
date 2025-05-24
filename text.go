package humle

import "io"

type Text string

func (n Text) Type() NodeType {
	return NodeText
}

func (n Text) WriteTo(w io.Writer) (int64, error) {
	len, err := w.Write([]byte(n))
	return int64(len), err
}
