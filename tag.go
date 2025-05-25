package humle

import (
	"io"
	"log"
	"strings"
)

type Tag struct {
	name          string
	children      []NodeWriter
	attributes    Attributes
	isVoidElement bool
	// Experimental
	namedSignals map[string][]Node
}

func (e *Tag) Type() NodeType {
	return NodeTag
}

// Identifies nodes and assigns them properly for later rendering
func NewTag(name string, nodes ...Node) *Tag {
	el := parseNodes(name, nodes)
	return el
}

func parseNodes(name string, nodes []Node) *Tag {
	el := Tag{
		name:       name,
		children:   []NodeWriter{},
		attributes: Attributes{},
	}
	for _, node := range nodes {
		switch n := node.(type) {
		case *Attribute:
			// Store attributes for later rendering
			el.attributes.Set(n.name, n)
		case *Tag:
			// Children for recursive rendering
			el.children = append(el.children, n)
		case Group:
			// Flattened into attributes or children
			g := parseNodes(name, n)
			for _, gattr := range g.attributes {
				el.attributes.Set(gattr.name, gattr)
			}
			el.children = append(el.children, g.children...)
		case RawHTML:
			el.children = append(el.children, n)
		case Text:
			// Simple child node
			el.children = append(el.children, n)
		case *SlotNode:
			el.children = append(el.children, n)

		default:
			log.Fatalf("Unknown node type in tag parsing: %T\n", node)
		}
	}
	return &el // We want this to be a fresh allocation
}

func (el *Tag) WriteTo(w io.Writer) (int64, error) {
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

func (el *Tag) String() string {
	var b strings.Builder
	el.WriteTo(&b)
	return b.String()

}

func WithVoidElement(el *Tag) *Tag {
	el.isVoidElement = true
	return el
}
