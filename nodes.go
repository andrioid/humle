package humle

import "io"

type NodeType int

const (
	NodeTag = iota
	NodeAttribute
	NodeTextValue
	NodeGroup
	NodeRawHTML
	NodeSlot
	NodeSignal
	//NodeReader // TODO: Allow us to pass io.Reader like files to it
)

type Node interface {
	Type() NodeType
}

// A node capable of writing HTML. Can be used as children
// E.g. Element, Text
type NodeWriter interface {
	Node
	io.WriterTo
}
