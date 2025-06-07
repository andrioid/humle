package humle

import (
	"html/template"
	"io"
)

type Text string

func (n Text) Type() NodeType {
	return NodeTextValue
}

func (n Text) ArgumentType() ArgumentType {
	return ArgumentText
}

func (n Text) WriteTo(w io.Writer) (int64, error) {
	template.HTMLEscape(w, []byte(n))
	return 0, nil
}

// TODO: I'd really like for Text to be work as an attribute too
