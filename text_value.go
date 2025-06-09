package humle

import (
	"html/template"
	"io"
	"strings"
)

type Text string

func (n Text) Type() NodeType {
	return NodeTextValue
}

func (n Text) WriteTo(w io.Writer) (int64, error) {
	template.HTMLEscape(w, []byte(n))
	return 0, nil
}

func (n Text) String() string {
	var b strings.Builder
	n.WriteTo(&b)
	return b.String()
}

// TODO: I'd really like for Text to be work as an attribute too
func (n Text) ArgumentType() ArgumentType {
	return ArgumentText
}
