package humle

import (
	"io"
	"strings"
)

type Element struct {
	name          string
	children      []NodeWriter
	attributes    *Attributes
	isVoidElement bool
}

func (e *Element) Type() NodeType {
	return NodeElement
}

// Identifies nodes and assigns them properly for later rendering
func NewElement(name string, nodes ...Node) *Element {
	el := parseNodes(name, nodes)
	return &el
}

func parseNodes(name string, nodes []Node) Element {
	el := Element{
		name:       name,
		children:   []NodeWriter{},
		attributes: &Attributes{},
	}
	for _, node := range nodes {
		switch n := node.(type) {
		case *Attribute:
			// Store attributes for later rendering
			el.attributes.Set(n.name, *n)
		case *Element:
			// Children for recursive rendering
			el.children = append(el.children, n)
		case *Group:
			// Flattened into attributes or children
			g := parseNodes(name, *n)
			for _, gattr := range *g.attributes {
				el.attributes.Set(gattr.name, gattr)
			}
			for _, gel := range g.children {
				el.children = append(el.children, gel)
			}
		case Text:
			// Simple child node
			el.children = append(el.children, n)
		}
	}
	return el
}

func (el *Element) WriteTo(w io.Writer) (int64, error) {
	attrs := el.attributes.String()
	w.Write([]byte("<" + el.name))
	if attrs != "" {
		w.Write([]byte(" " + attrs))
	}
	w.Write([]byte(">"))

	// Inner
	for _, child := range el.children {
		if len, err := child.WriteTo(w); err != nil {
			return len, err
		}
	}

	if el.isVoidElement {
		return 0, nil
	}

	w.Write([]byte("</" + el.name + ">"))

	return 0, nil
}

func (el *Element) String() string {
	var b strings.Builder
	el.WriteTo(&b)
	return b.String()

}

func WithVoidElement(el *Element) *Element {
	el.isVoidElement = true
	return el
}
