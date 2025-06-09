package humle

import (
	"io"
	"log"
	"strings"
)

type TagNode struct {
	name          string
	children      []Node
	arguments     Arguments
	isVoidElement bool
}

func (e *TagNode) Type() NodeType {
	return NodeTag
}

// Identifies nodes and assigns them properly for later rendering
func Tag(name string, args ...Argument) *TagNode {
	el := TagNode{
		name:      name,
		children:  []Node{},
		arguments: Arguments{},
	}
	el.arguments = args

	return &el
}

func (el *TagNode) parseChildren(nodes []Node) {
	for _, node := range nodes {
		switch n := node.(type) {
		case *TagNode:
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

func (el *TagNode) Children(children ...Node) *TagNode {
	el.parseChildren(children)
	return el
}

func (el *TagNode) WriteTo(w io.Writer) (int64, error) {
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

func (n *TagNode) String() string {
	var b strings.Builder
	n.WriteTo(&b)
	return b.String()
}

func WithVoidElement(el *TagNode) *TagNode {
	el.isVoidElement = true
	return el
}
