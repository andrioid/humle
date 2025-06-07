package humle

import (
	"io"
	"log"
)

type Tag struct {
	name          string
	children      []Node
	arguments     Arguments
	isVoidElement bool
}

func (e *Tag) Type() NodeType {
	return NodeTag
}

// Identifies nodes and assigns them properly for later rendering
func NewTag(name string, args ...Argument) *Tag {
	el := Tag{
		name:      name,
		children:  []Node{},
		arguments: Arguments{},
	}
	el.arguments = args

	return &el
}

func (el *Tag) parseChildren(nodes []Node) {
	for _, node := range nodes {
		switch n := node.(type) {
		case *Tag:
			// Children for recursive rendering
			el.children = append(el.children, n)
		case Group:
			// Flattened into attributes or children
			el.parseChildren(n)
		case Node:
			el.children = append(el.children, n)
		default:
			log.Fatalf("Unknown node type in tag parsing: %T\n", node)
		}
	}
}

func (el *Tag) Children(children ...Node) *Tag {
	el.parseChildren(children)
	return el
}

func (el *Tag) WriteTo(w io.Writer) (int64, error) {
	attrs := el.arguments.GetAttributes().String()
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
	// Node-like attributes
	for _, child := range el.arguments {
		if n, ok := child.(Node); ok {
			n.WriteTo(w)
		}
	}

	if el.isVoidElement {
		return 0, nil
	}

	w.Write([]byte("</" + el.name + ">"))

	return 0, nil
}

func WithVoidElement(el *Tag) *Tag {
	el.isVoidElement = true
	return el
}
